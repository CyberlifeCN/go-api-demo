package controllers

import (
	"go-api-demo/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
	"github.com/satori/go.uuid"
	"strings"
	// "fmt"
	// _"github.com/astaxie/beego/orm"
	// _"strconv"
	// "strconv"
	// "github.com/astaxie/beego/orm"
)


// Operations about Test
type TestController struct {
	beego.Controller
}


// @Title Get
// @Description get test by uid
// @Param	uid		path 	string	true		"The key for test"
// @Success 200 {object} models.QueryTestOneResponse
// @Failure 403 :uid is empty
// @router /:uid [get]
func (t *TestController) Get() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
  beego.Trace(uid)

	test := models.GetTest(uid)
	beego.Trace(test)

	var rs = &models.QueryTestOneResponse{
		Code: 200,
		Msg: "Success",
		Rs: test,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}


// @Title GetAll
// @Description get all Test by pagination
// header-->请求参数的获取：@RequestHeader
// query-->请求参数的获取：@RequestParam
// path（用于restful接口）-->请求参数的获取：@PathVariable
// body（不常用）
// form（不常用）
// @Param	page 		query		int		1		false		"The page number of dataset in mysql:test"
// @Param	limit 	query		int		20	false		"One page size"
// @Success 200		{object} models.QueryTestAllResponse
// @router / [get]
func (t *TestController) GetAll() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	var page int = 0
	t.Ctx.Input.Bind(&page, "page")  //page==1
	var limit int = 20
	t.Ctx.Input.Bind(&limit, "limit")  //limit==20
	if (page <= 0) {
		page = 1
	}
	if (limit <= 0) {
		limit = 20
	}
	// Debug, Info, Warn, Error, Fatal
	beego.Trace("page:", page, " limit:", limit)

	var idx int = (page-1) * limit
	tests, size := models.GetAllTest(idx, limit)
	total_num := models.GetAllTestCount()
	beego.Trace("total_num:", total_num, "tests:", tests)

	var total_page int = total_num/limit
	if (total_num % limit > 0) {
		total_page++
	}
	beego.Trace("total_page:", total_page)

	var rs = &models.QueryTestAllResponse{
		Code: 200,
		Msg: "Success",
		Rs: models.TestPaginationResultSet{
			Page: page,
			Size: size,
			TotalPage: total_page,
			Datas: tests,
		},
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}


// @Title CreateTest
// @Description create test
// @Param	body		body 	models.Test	true		"body for test content"
// @Success 200 {object} models.ActionTestOneResponse
// @Failure 403 body is empty
// @router / [post]
func (t *TestController) Post() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)
	beego.Info(t.Ctx.Request.Body)

	var test models.Test
	json.Unmarshal(t.Ctx.Input.RequestBody, &test)
	beego.Trace(test)

	// Creating UUID Version 4
	uid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	test.Id = uid
  timestamp := time.Now().UnixNano() / 1000000 // 毫秒
	test.Ctime = timestamp
	beego.Trace(test)

	models.AddTest(test)

	var rs = &models.ActionTestOneResponse{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}


// @Title Update
// @Description update the test
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Test	true		"body for test content"
// @Success 200 {object} models.Test
// @Failure 403 :uid is empty, or body is empty
// @router /:uid [put]
func (t *TestController) Put() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
	beego.Trace(uid)
	var test models.Test
	json.Unmarshal(t.Ctx.Input.RequestBody, &test)
	beego.Trace(test)

	test.Id = uid
	timestamp := time.Now().UnixNano() / 1000000 // 毫秒
	test.Mtime = timestamp
	models.UpdateTest(uid, &test)

	var rs = &models.ActionTestOneResponse{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}


// @Title Delete
// @Description delete the test
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {object} models.ActionTestOneResponse
// @Failure 403 uid is empty
// @router /:uid [delete]
func (t *TestController) Delete() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
	beego.Trace(uid)

	models.DeleteTest(uid)

	var rs = &models.ActionTestOneResponse{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}
