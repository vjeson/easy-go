package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type tableCol struct {
	columnName    string
	dataType      string
	columnComment string
}

func GenTableToStruct()  {
	Db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/ad")
	if err != nil {
		panic(err.Error())
	}


	sqlStr := "SELECT column_name, data_type,column_comment  FROM information_schema.columns WHERE table_schema = ? AND table_name = ?"
	rows, err := Db.Query(sqlStr, "ad", "phone_user")

	if err != nil {
		panic(err.Error())
	}

	var tabs []*tableCol
	for rows.Next() {
		ts := &tableCol{}
		rows.Scan(&ts.columnName, &ts.dataType, &ts.columnComment)
		tabs = append(tabs, ts)
	}
}
