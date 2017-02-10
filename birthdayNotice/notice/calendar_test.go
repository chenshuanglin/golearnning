package notice

import (
	"strings"
	"testing"
)

func TestNewCalendar(t *testing.T) {
	date := "2017-1-10"
	c, _ := NewCalendar(date)
	if !strings.EqualFold(c.Result.Data.Lunar, "腊月十三") {
		t.Error("返回万年历信息失败")
	}
}
