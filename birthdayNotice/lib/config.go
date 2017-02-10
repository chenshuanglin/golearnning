package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type config struct {
	Version string
	Appid   string
	Users   []User
}

type User struct {
	Name   string
	Date   string `json:"month-day"`
	Type   string
	Before int32
}

const configPath = "../config.json"

var c config

func init() {
	//捕获抛出的异常
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败，原因是:%v", err))
	}
	log.Println("读取config.json配置文件成功")
	err = json.Unmarshal(bytes, &c)
}

func GetAppid() string {
	return c.Appid
}

func GetVersion() string {
	return c.Version
}

func GetUsers() []User {
	return c.Users
}
