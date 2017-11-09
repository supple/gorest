package core

import "fmt"

func Log(msg ...string) {
	if len(msg) == 1 {
		fmt.Println(GetJodaTime(), msg[0])
	} else {
		fmt.Println(GetJodaTime(), msg)
	}
}
