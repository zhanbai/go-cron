package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/robfig/cron/v3"
)

type CmdJob struct {
	Cmd string
}

func (c CmdJob) Run() {
	cmd := exec.Command("sh", "-c", c.Cmd)
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

func main() {
	hello := "echo 'hello world'"

	c := cron.New()

	c.AddJob("@every 2s", CmdJob{hello})
	c.AddFunc("@every 5s", func() {
		fmt.Println("hello world")
	})

	go c.Start()
	defer c.Stop()

	select {}
}
