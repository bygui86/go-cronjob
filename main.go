package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
)

func main() {
	cronjob := cron.New()
	cronjob.AddFunc("@every 30s", job())
	cronjob.Start()

	fmt.Printf("CronJob entries: %v\n", cronjob.Entries())

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	fmt.Println("Termination signal received!")

	cronjob.Stop()
}

func job() func() {
	return func() {
		fmt.Println("CronJob execution")
		time.Sleep(5 * time.Second)
	}
}
