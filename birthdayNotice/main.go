package main

import "birthdayNotice/notice"

func main() {
	users := notice.GetUsers()
	for _, user := range users {
		notice.Notice(&user)
	}
}
