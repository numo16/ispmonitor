package main

import (
	"fmt"
	"log"

	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/parnurzeal/gorequest"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	home, _ := homedir.Dir()
	logDir := fmt.Sprintf("%s/log/ispmonitor/", home)

	log.SetOutput(&lumberjack.Logger{
		Dir:       logDir,
		MaxSize:   10 * lumberjack.Megabyte,
		MaxAge:    28,
		LocalTime: true,
	})

	request := gorequest.New().Timeout(10 * time.Second)

	for {
		resp, _, err := request.Get("http://google.com").End()
		if err != nil {
			log.Print("ERROR: The system is down, the system is down")
			fmt.Print("ERROR: The system is down, the system is down\n")
		} else {
			log.Print("SUCCESS: Internet is up")
			fmt.Printf("SUCCESS: Internet is up %s\n", resp.Status)
		}

		time.Sleep(time.Minute)
	}
}
