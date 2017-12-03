package controllers

import (
	"go-api-demo/models"
	// "encoding/json"
	"github.com/astaxie/beego"
	// "time"
	// "github.com/satori/go.uuid"
	// "strings"
  "log"
  "fmt"
  "net/rpc"
)


// Operations about RPC
type RpcController struct {
	beego.Controller
}

var GlobalRpcClient *rpc.Client
var err error

func init() {
  service := "127.0.0.1:1234"
  client, err := rpc.Dial("tcp", service)
  if err != nil {
      log.Fatal("dialing:", err)
  }
  GlobalRpcClient = client
}

// @Title Get
// @Description get test by uid though RPC
// @Param	uid		path 	string	true		"The key for test"
// @Success 200 {object} models.TestQueryOneResp
// @Failure 403 :uid is empty
// @router /:uid [get]
func (t *RpcController) Get() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	uid := t.GetString(":uid")
  beego.Trace(uid)

  args := models.Args{uid}
  var test = &models.Test{}
  err = GlobalRpcClient.Call("Mysql.QueryTest", args, &test)
  if err != nil {
      log.Fatal("mysql error :", err)
  }
  fmt.Printf("Mysql: %d*%d=%d\n", args.Id, test)
  // defer GlobalRpcClient.Close()

	if (test != nil && test.Id != "") {
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
