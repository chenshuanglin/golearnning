package notice

import (
	"strings"
	"testing"
)

func TestGetAppid(t *testing.T) {
	appid := GetAppid()
	if !strings.EqualFold(appid, "169eedea0b95976f7da065ca002a9805") {
		t.Error("没有获取正确的appid")
	}
}

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if !strings.EqualFold(version, "1.0.0") {
		t.Error("没有获取正确的version")
	}
}

func TestGetUsers(t *testing.T) {
	user := GetUsers()
	if !strings.EqualFold(user[0].Name, "舒文静") {
		t.Error("没有获取到用户名称")
	}
}
