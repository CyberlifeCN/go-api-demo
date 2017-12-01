package controllers

import (
	"go-api-demo/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
	"github.com/satori/go.uuid"
	"strings"
)


// Operations about Memcache
type MemcacheController struct {
	beego.Controller
}


// @Title Get
// @Description get memcache by uid
// @Param	uid		path 	string	true		"The key for memcache"
// @Success 200 {object} models.TestQueryOneResp
// @Failure 403 :uid is empty
// @router /:uid [get]
func (t *MemcacheController) Get() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
  beego.Trace(uid)

	var test = models.GetMemcache(uid)
	if (test != nil) {
		beego.Trace(test)

		var rs = &models.TestQueryOneResp{
			Code: 200,
			Msg: "Success",
			Rs: *test,
		}

		t.Data["json"] = *rs
		t.ServeJSON()
	} else {
		var rs = &models.TestQueryOneResp{
			Code: 404,
			Msg: "Not Found",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
	}
}


// @Title CreateMemcache
// @Description create memcache
// @Param	Authorization		header 	string	"Bearer access_token"   true		"Bearer access_token"
// @Param	body		body 	models.Test	true		"body for test content"
// @Success 200 {object} models.TestActionOneResp
// @Failure 403 body is empty
// @router / [post]
func (t *MemcacheController) Post() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)
	beego.Info(t.Ctx.Request.Body)

	var test models.Test
	json.Unmarshal(t.Ctx.Input.RequestBody, &test)
	beego.Trace(test)
	if (test.Name == "") {
		var rs = &models.TestActionOneResp{
			Code: 403,
			Msg: "Bad Request",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
		return
	}

	//Authorization=="Bearer access_token"
	auth := t.Ctx.Input.Header("Authorization")
	beego.Trace(auth)
	if (auth == "") {
		var rs = &models.TestActionOneResp{
			Code: 401,
			Msg: "Unauthorized",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
		return
	}

	access_token := strings.Replace(auth, "Bearer ", "", -1)
	beego.Trace(access_token)
	if (access_token == "") {
		var rs = &models.TestActionOneResp{
			Code: 401,
			Msg: "Unauthorized",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
		return
	}

	// TODO check access_token

	// Creating UUID Version 4
	uid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	test.Id = uid
  timestamp := time.Now().UnixNano() / 1000000 // 毫秒
	test.Ctime = timestamp
	test.Mtime = timestamp
	beego.Trace(test)

	models.AddMemcache(test)

	var rs = &models.TestActionOneResp{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}
