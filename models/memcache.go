package models

import (
  "encoding/json"
  "fmt"

  "github.com/bradfitz/gomemcache/memcache"
)


func GetMemcache(uid string) *Test {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  // Get a single value
  val, err := mc.Get(uid)
  if err != nil {
    panic(err)
  }

  fmt.Println("Item:", val)
  fmt.Println("Key:", val.Key)
  fmt.Println("Value:", val.Value)

  var t = &Test{}
  json.Unmarshal(val.Value, &t)
  fmt.Println("Test:", t)

	return t
}


func AddMemcache(t Test) {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  jsonBytes, err := json.Marshal(t)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes:", jsonBytes)

  // Set some values
  mc.Set(&memcache.Item{Key: t.Id, Value: []byte(jsonBytes)})
}
