package main

import (
	"birthdayNotice/lib"
)

func main() {
	users := lib.GetUsers()
	for _, user := range users {
		lib.Notice(&user)
	}
}
