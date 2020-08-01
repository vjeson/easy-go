package middleware

import (
	"demo/service"
	"demo/util"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)



func Jwt() gin.HandlerFunc  {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			res := &util.Response{
				Code: 1100,
				Msg: "请求未携带token，无权限访问",
			}
			res.Json(c)
			c.Abort()
			return
		}
		logrus.Debug("get Token:", token)

		j := NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			res := &util.Response{
				Code: 1100,
				Msg: err.Error(),
			}
			res.Json(c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		//claims = c.MustGet("claims").(*CustomClaims)

		c.Next()
	}
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Id    string `json:"userId"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

var (
	TokenExpired     error  = errors.New("token is expired")
	TokenNotValidYet error  = errors.New("token not active yet")
	TokenMalformed   error  = errors.New("that's not even a token")
	TokenInvalid     error  = errors.New("couldn't handle this token")
	SignKey          string = "bg-lnmp"
)

func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

func (j *JWT) GenerateToken(user *service.Users) (map[string]interface{}, error) {

	expires := 7200
	claims := CustomClaims{
		Id:  			strconv.Itoa(user.Id),
		Name: 			user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + int64(expires)), // 过期时间 一小时
			Issuer:    SignKey,                   //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)

	if err != nil {
		return nil, err
	}

	logrus.Debug("generateToken >> " ,token)

	return map[string]interface{}{"toke": token, "expires": expires}, nil
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

