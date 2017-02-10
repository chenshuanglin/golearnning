package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	urlCalendar = "http://v.juhe.cn/calendar/day"
	appid       = "169eedea0b95976f7da065ca002a9805"
)

type Calender struct {
	Code   int `json:"error_code"`
	Reason string
	Result ResultData
}

type ResultData struct {
	Data Message
}

type Message struct {
	Holiday   string
	Avoid     string
	Animals   string `json:"animalsYear"`
	Desc      string
	Weekday   string
	Suit      string
	LunarYear string
	Lunar     string
	yearMonth string `json:"year-month"`
	Date      string
}

func main() {
	c, _ := GetDateMessage("2017-1-4")
	fmt.Println(c.Result.Data.Lunar)
}

//根据当前的时间点，获取当天的信息
func GetDateMessage(date string) (Calender, error) {
	params := url.Values{}
	params.Set("date", date)
	params.Set("key", appid)

	var c Calender
	res, err := Get(urlCalendar, params)
	if err != nil {
		fmt.Printf("获取万年历信息失败:%v\n", err)
		return c, err
	}
	err = json.Unmarshal(res, &c)
	if err != nil {
		fmt.Println("解析的json字符串有误：%s,错误:%v\n", string(res), err)
		return c, err
	}
	return c, nil
}

//Get带参数使用该方法
func Get(apiurl string, params url.Values) (rs []byte, err error) {
	u, err := url.Parse(apiurl)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v\n", err)
		return nil, err
	}
	u.RawQuery = params.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("http获取%s:\r\n%v\n", u.String(), err)
		return nil, err
	}
	defer resp.Body.Close()
	rs, _ = ioutil.ReadAll(resp.Body)
	return rs, nil
}
