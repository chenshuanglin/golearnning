package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		UserId int    `json:"user_id" bson:"user_id"`
		Name   string `json:"user_name" bson:"user_name_bson"`
	}

	u := &user{
		UserId: 1,
		Name:   "chenshuanglin",
	}

	b, _ := json.Marshal(&u)
	fmt.Println(string(b))

	t := reflect.TypeOf(u)
	f := t.Elem().Field(1)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("bson"))
}
