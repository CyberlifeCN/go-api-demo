package models

import (
	// "strconv"
	// "time"
	// "database/sql"
  "log"
	// "fmt"
  // "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  // Id bson.ObjectId `bson:"_id,omitempty"`
  Id        string    `bson:"_id,omitempty"`
  Name      string    `json:"name"`
  Phone     string    `json:"phone"`
}

type PersonQueryOneResp struct {
	Code   		int 											`json:"err_code"`
	Msg    		string 										`json:"err_msg"`
  Rs   	 		Person 										`json:"rs"`
}

type PersonActionOneResp struct {
	Code   		int 											`json:"err_code"`
	Msg    		string 										`json:"err_msg"`
  Rs   	 		string 										`json:"rs"`
}


func GetPerson(uid string) *Person {
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("people")

  //查询数据
  var result = &Person{}
  query := c.Find(bson.M{"_id": uid})
  var err = query.One(result)
  if err != nil {
    log.Print(err)
    return nil
  }
  // If you must detect "not found" case:
  if result == nil {
      // No result
      return nil
  }

	return result
}


func AddPerson(p Person) {
	//插入数据
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("people")

  var err = c.Insert(p)
  if err != nil {
    panic(err)
  }
}


func UpdatePerson(p Person) {
	//修改数据
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("people")

  var err = c.Update(bson.M{"_id": p.Id}, bson.M{"$set": bson.M{"name": p.Name, "phone":p.Phone}})
  if err != nil {
    panic(err)
  }
}


func DeletePerson(uid string) {
	//删除数据
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("people")

  var _, err = c.RemoveAll(bson.M{"_id": uid})
  if err != nil {
    panic(err)
  }
}
