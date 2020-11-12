package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	fmt.Println("hello world")
	sysLog, err := syslog.Dial("", "", syslog.LOG_ERR, "gopher says")
	if err != nil {
		log.Fatal(err)
	}

	sysLog.Emerg("gopher's log")
}
