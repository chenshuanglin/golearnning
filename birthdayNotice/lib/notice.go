/*
 * author:chenshuanglin
 * descript: 主要是用来判断是否按计划快到生日了
 */
package lib

import (
	"bytes"
	"log"
	"strings"
	"time"
)

//农历月份跟新历月份做个对应
var monthLunar map[string]string

//新历的天数和农历的天数做个对应
var dayLunar map[string]string

func init() {
	monthLunar = map[string]string{
		"1": "正月", "2": "二月", "3": "三月", "4": "四月", "5": "五月", "6": "六月",
		"7": "七月", "8": "八月", "9": "九月", "10": "十月", "11": "十一月", "12": "腊月",
	}
	dayLunar = map[string]string{
		"1": "初一", "2": "初二", "3": "初三", "4": "初四", "5": "初五", "6": "初六",
		"7": "初七", "8": "初八", "9": "初九", "10": "初十", "11": "十一", "12": "十二",
		"13": "十三", "14": "十四", "15": "十五", "16": "十六", "17": "十七", "18": "十八",
		"19": "十九", "20": "廿十", "21": "廿一", "22": "廿二", "23": "廿三", "24": "廿四",
		"25": "廿五", "26": "廿六", "27": "廿七", "28": "廿八", "29": "廿九", "30": "卅十",
	}
}

func Notice(user *User) {
	if strings.EqualFold(user.Type, "农历") {
		ok, err := isNoticeLunar(user.Date, time.Duration(user.Before))
		if err != nil {
			log.Printf("判断是否到农历生日失败,原因是%v\n", err)
			return
		}
		if !ok {
			log.Printf("用户:%s 农历生日还没到\n", user.Name)
			return
		}
		log.Printf("用户:%s 农历生日到了，农历日期为%s\n", user.Name, user.Date)

	} else {
		ok := isNoticeNew(user.Date, time.Duration(user.Before))
		if !ok {
			log.Printf("用户:%s 新历生日还没到\n", user.Name)
			return
		}
		log.Printf("用户:%s 新历生日到了，日期为%s\n", user.Name, user.Date)
	}
}

/*
 * 功能：判断是否按计划快要提醒快到农历生日了
 * params[day] 要提前几天提醒
 * params[date] 生日时间
 */
func isNoticeLunar(date string, day time.Duration) (bool, error) {
	c, err := NewCalendar(getDate(day))
	if err != nil {
		return false, err
	}
	//换算成农历时间
	lunar := getLunarDay(date)
	if strings.EqualFold(lunar, c.Result.Data.Lunar) {
		return true, nil
	} else {
		return false, nil
	}
}

/*
 * 功能：判断是否按计划快要提醒快到新历生日了
 * params[day] 要提前几天提醒
 * params[date] 生日时间
 */

func isNoticeNew(date string, day time.Duration) bool {
	futureDay := getDate(day)
	dates := strings.Split(futureDay, "-")
	var buf bytes.Buffer
	buf.WriteString(dates[1])
	buf.WriteString("-")
	buf.WriteString(dates[2])
	return strings.EqualFold(date, buf.String())
}

//计算在过day天后的时间
func getDate(day time.Duration) (date string) {
	now := time.Now()
	now = now.Add(time.Hour * 24 * day)
	tw := time.Unix(now.Unix(), 0)
	date = tw.Format("2006-1-2")
	return
}

//把新历生日换算成农历生日
func getLunarDay(date string) string {
	dates := strings.Split(date, "-")
	var buf bytes.Buffer
	buf.WriteString(monthLunar[dates[0]])
	buf.WriteString(dayLunar[dates[1]])
	return buf.String()
}