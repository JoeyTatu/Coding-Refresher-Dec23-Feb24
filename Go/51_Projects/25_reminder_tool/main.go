package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "GO_CLI_REMINDER"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <hh:mm> <message>", os.Args[0])
	}

	now := time.Now()
	when := when.New(nil)
	when.Add(en.All...)
	when.Add(common.All...)

	t, err := when.Parse(os.Args[1], now)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	if t == nil {
		log.Fatal("Unable to parse time")
	}
	if now.After(t.Time) {
		log.Fatal("Set a time in the future")
	}

	difference := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(difference)
		err = beeep.Alert("Reminder:", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			log.Fatal("Error: ", err)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			log.Fatal("Error: ", err)
		}
		fmt.Println("Reminder will be displayed in", difference.Round(time.Second))
		os.Exit(0)
	}

}
