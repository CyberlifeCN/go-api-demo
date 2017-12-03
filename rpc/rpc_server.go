package main

import (
    // "errors"
    "fmt"
    "net"
    "net/rpc"
    "os"
    // "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "go-api-demo/models"
)

//global
var GlobalMysqlConnPool *sql.DB
type Mysql int


func init() {
	// init mysql connection pool
	GlobalMysqlConnPool, _ = sql.Open("mysql", "legend_dev:need4sPeed@tcp(rm-2zeyubz4yre340644o.mysql.rds.aliyuncs.com:3306)/legend_dev?charset=utf8mb4")
	GlobalMysqlConnPool.SetMaxOpenConns(20)
	GlobalMysqlConnPool.SetMaxIdleConns(10)
	GlobalMysqlConnPool.Ping()
	// defer GlobalMysqlConnPool.Close()
}


func (t *Mysql) QueryTest(args *models.Args, reply *models.Test) error {
  var uid = args.Id
  fmt.Printf("Mysql.QueryTest: %d", uid)

  //查询数据
	var _id string
	var name string
	var ctime int64
	var mtime int64
	err := GlobalMysqlConnPool.QueryRow("SELECT _id,name,ctime,mtime FROM test WHERE _id=?", uid).Scan(&_id, &name, &ctime, &mtime)
	if (err != nil) {
    fmt.Printf("Mysql.QueryTest: %d=nil\n", uid)
    reply = nil
		return nil
	} else {
    var test = &models.Test{}
  	test.Id = _id
  	test.Name = name
  	test.Ctime = ctime
  	test.Mtime = mtime

    *reply = *test
    fmt.Printf("Mysql.QueryTest: %d=%d\n", uid, *reply)
    return nil
	}
}


func main() {
    mysql := new(Mysql)
    rpc.Register(mysql)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
    if err != nil {
        fmt.Println("Fatal error:", err)
        os.Exit(1)
    }

    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        fmt.Println("Fatal error:", err)
        os.Exit(1)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        rpc.ServeConn(conn)
    }

    // for {
    //     time.Sleep(1 * time.Second)
    // }
}
