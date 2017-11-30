package models

import (
	// "strconv"
	// "time"
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

var dbPool *sql.DB


func init() {
	dbPool, _ = sql.Open("mysql", "legend_dev:need4sPeed@tcp(rm-2zeyubz4yre340644o.mysql.rds.aliyuncs.com:3306)/legend_dev?charset=utf8mb4")
	dbPool.SetMaxOpenConns(20)
	dbPool.SetMaxIdleConns(10)
	dbPool.Ping()
	// defer dbPool.Close()
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
