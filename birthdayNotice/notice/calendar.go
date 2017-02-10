package notice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const CalendarUrl = "http://v.juhe.cn/calendar/day"

type Calendar struct {
	Code   int `json:"error_code"`
	Reason string
	Result Message
}

type Message struct {
	Data CData
}

type CData struct {
	Holiday     string
	Avoid       string
	AnimalsYear string
	Desc        string
	Weekday     string
	Suit        string
	Lunar       string
	YearMonth   string `json:"year-month"`
	Date        string
}

func NewCalendar(date string) (*Calendar, error) {
	var c Calendar
	bytes, err := getCal(date)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &c)
	return &c, nil
}

func getCal(date string) ([]byte, error) {
	vals := make(url.Values)
	appKey := GetAppid()
	vals.Add("key", appKey)
	vals.Add("date", date)
	return Get(CalendarUrl, vals)
}

func Get(orgurl string, paras url.Values) ([]byte, error) {
	u, err := url.Parse(orgurl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
