package controllers

import (
	"go-api-demo/models"
	"encoding/json"
	"github.com/astaxie/beego"
	// "time"
	"github.com/satori/go.uuid"
	"strings"
	// "fmt"
	// _"github.com/astaxie/beego/orm"
	// _"strconv"
	// "strconv"
	// "github.com/astaxie/beego/orm"
)


// Operations about Person
type PersonController struct {
	beego.Controller
}


// @Title Get
// @Description get person by uid
// @Param	uid		path 	string	true		"The key for person"
// @Success 200 {object} models.PersonQueryOneResp
// @Failure 403 :uid is empty
// @router /:uid [get]
func (t *PersonController) Get() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
  beego.Trace(uid)

	var person = models.GetPerson(uid)
	if (person != nil) {
		beego.Trace(person)

		var rs = &models.PersonQueryOneResp{
			Code: 200,
			Msg: "Success",
			Rs: *person,
		}

		t.Data["json"] = *rs
		t.ServeJSON()
	} else {
		var rs = &models.PersonQueryOneResp{
			Code: 404,
			Msg: "Not Found",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
	}
}


// @Title CreatePerson
// @Description create person
// @Param	Authorization		header 	string	"Bearer access_token"		true		"Bearer access_token"
// @Param	body		body 	models.Person	true		"body for person content"
// @Success 200 {object} models.PersonActionOneResp
// @Failure 403 body is empty
// @router / [post]
func (t *PersonController) Post() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)
	beego.Info(t.Ctx.Request.Body)

	var person models.Person
	json.Unmarshal(t.Ctx.Input.RequestBody, &person)
	beego.Trace(person)
	if (person.Name == "") {
		var rs = &models.PersonActionOneResp{
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
		var rs = &models.PersonActionOneResp{
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
		var rs = &models.PersonActionOneResp{
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
	person.Id = uid
	beego.Trace(person)

	models.AddPerson(person)

	var rs = &models.PersonActionOneResp{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}


// @Title Update
// @Description update the person
// @Param	Authorization		header 	string	"Bearer access_token"		true		"Bearer access_token"
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Person	true		"body for person content"
// @Success 200 {object} models.PersonActionOneResp
// @Failure 403 :uid is empty, or body is empty
// @router /:uid [put]
func (t *PersonController) Put() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
	beego.Trace(uid)
	var person models.Person
	json.Unmarshal(t.Ctx.Input.RequestBody, &person)
	beego.Trace(person)

	if (person.Name == "") {
		var rs = &models.PersonActionOneResp{
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
		var rs = &models.PersonActionOneResp{
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
		var rs = &models.PersonActionOneResp{
			Code: 401,
			Msg: "Unauthorized",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
		return
	}

	// TODO check access_token

	person.Id = uid
	models.UpdatePerson(person)

	var rs = &models.PersonActionOneResp{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}


// @Title Delete
// @Description delete the person
// @Param	Authorization		header 	string	"Bearer access_token"		true		"Bearer access_token"
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {object} models.PersonActionOneResp
// @Failure 403 uid is empty
// @router /:uid [delete]
func (t *PersonController) Delete() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
	beego.Trace(uid)

	//Authorization=="Bearer access_token"
	auth := t.Ctx.Input.Header("Authorization")
	beego.Trace(auth)
	if (auth == "") {
		var rs = &models.PersonActionOneResp{
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
		var rs = &models.PersonActionOneResp{
			Code: 401,
			Msg: "Unauthorized",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
		return
	}

	// TODO check access_token

	models.DeletePerson(uid)

	var rs = &models.PersonActionOneResp{
		Code: 200,
		Msg: "Success",
		Rs: uid,
	}

	t.Data["json"] = *rs
	t.ServeJSON()
}
