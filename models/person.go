package models

import (
	// "strconv"
	// "time"
	// "database/sql"
  "log"
	// "fmt"
  "gopkg.in/mgo.v2"
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


const (
  URL = "mongodb://legend_dev:uWx-nJs-8J3-vA9@123.56.165.59:3717/legend_dev"
)

func GetPerson(uid string) *Person {
  session, err := mgo.Dial(URL)
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("legend_dev").C("people")

  //查询数据
  var result = &Person{}
  query := c.Find(bson.M{"_id": uid})
  err = query.One(result)
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
  session, err := mgo.Dial(URL)
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("legend_dev").C("people")

  err = c.Insert(p)
  if err != nil {
    panic(err)
  }
}


func UpdatePerson(p Person) {
	//修改数据
  session, err := mgo.Dial(URL)
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("legend_dev").C("people")

  err = c.Update(bson.M{"_id": p.Id}, bson.M{"$set": bson.M{"name": p.Name, "phone":p.Phone}})
  if err != nil {
    panic(err)
  }
}


func DeletePerson(uid string) {
	//删除数据
  session, err := mgo.Dial(URL)
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("legend_dev").C("people")

  _, err = c.RemoveAll(bson.M{"_id": uid})
  if err != nil {
    panic(err)
  }
}
