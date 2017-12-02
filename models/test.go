package models

import (
	// "strconv"
	// "time"
	// "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Test struct {
	Id      	string										`json:"_id"`
	Name    	string										`json:"name"`
  Ctime   	int64											`json:"ctime"`
	Mtime   	int64											`json:"mtime"`
}

type TestPaginationResultSet struct {
	Page 			int 											`json:"page"`
	Size 			int 											`json:"size"`
	TotalPage int 											`json:"total_page"`
	Datas 		[]Test 										`json:"datas"`
}

type TestQueryAllResp struct {
	Code   		int 											`json:"err_code"`
	Msg    		string 										`json:"err_msg"`
  Rs   	 		TestPaginationResultSet 	`json:"rs"`
}

type TestQueryOneResp struct {
	Code   		int 											`json:"err_code"`
	Msg    		string 										`json:"err_msg"`
  Rs   	 		Test 											`json:"rs"`
}


func GetAllTest(idx int, limit int) ([]Test,int) {
	//查询数据
	rows, err := GlobalMysqlConnPool.Query("SELECT _id,name,ctime,mtime FROM test limit ?,?", idx, limit)
	checkErr(err)

	var size int = 0
	var result = make([]Test, 0)
	for rows.Next() {
		var _id string
		var name string
		var ctime int64
		var mtime int64
		err = rows.Scan(&_id, &name, &ctime, &mtime)
		checkErr(err)

		var test = &Test{}
		test.Id = _id
		test.Name = name
		test.Ctime = ctime
		test.Mtime = mtime
		result = append(result, *test)

		size++
 	}

	return result, size
}


func GetAllTestCount() int {
	//查询数据
	var num int
	err := GlobalMysqlConnPool.QueryRow("SELECT count(_id) as num FROM test").Scan(&num)
	checkErr(err)

	return num
}


func GetTest(uid string) *Test {
	//查询数据
	var _id string
	var name string
	var ctime int64
	var mtime int64
	err := GlobalMysqlConnPool.QueryRow("SELECT _id,name,ctime,mtime FROM test WHERE _id=?", uid).Scan(&_id, &name, &ctime, &mtime)
	if (err != nil) {
		return nil
	} else {
	  panic(err)
	}

	var test = &Test{}
	test.Id = _id
	test.Name = name
	test.Ctime = ctime
	test.Mtime = mtime
	fmt.Println(test)

	return test
}


func AddTest(tt Test) {
	//插入数据
	stmt, err := GlobalMysqlConnPool.Prepare("INSERT test SET _id=?,name=?,ctime=?,mtime=?")
	checkErr(err)

	res, err := stmt.Exec(tt.Id, tt.Name, tt.Ctime, tt.Mtime)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}


func UpdateTest(uid string, tt *Test) {
	//更新数据
	stmt, err := GlobalMysqlConnPool.Prepare("UPDATE test set name=?,mtime=? WHERE _id=?")
	checkErr(err)

	res, err := stmt.Exec(tt.Name, tt.Mtime, uid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}


func DeleteTest(uid string) {
	//删除数据
	stmt, err := GlobalMysqlConnPool.Prepare("DELETE FROM test WHERE _id=?")
	checkErr(err)

	res, err := stmt.Exec(uid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
