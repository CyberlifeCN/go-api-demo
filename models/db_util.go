package models

import (
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
  // "gopkg.in/mgo.v2/bson"
)

//global
var GlobalMysqlConnPool *sql.DB
var GlobalMgoSession *mgo.Session

const (
  URL = "mongodb://legend_dev:uWx-nJs-8J3-vA9@123.56.165.59:3717/legend_dev"
)


type ActionOneResp struct {
	Code   		int 											`json:"err_code"`
	Msg    		string 										`json:"err_msg"`
  Rs   	 		string 										`json:"rs"`
}


func init() {
	// init mysql connection pool
	GlobalMysqlConnPool, _ = sql.Open("mysql", "legend_dev:need4sPeed@tcp(rm-2zeyubz4yre340644o.mysql.rds.aliyuncs.com:3306)/legend_dev?charset=utf8mb4")
	GlobalMysqlConnPool.SetMaxOpenConns(20)
	GlobalMysqlConnPool.SetMaxIdleConns(10)
	GlobalMysqlConnPool.Ping()
	// defer GlobalMysqlConnPool.Close()

	// init mongodb connection pool
	globalMgoSession, err := mgo.DialWithTimeout(URL, 10 * time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	// Optional. Switch the session to a monotonic behavior.
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(30)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
