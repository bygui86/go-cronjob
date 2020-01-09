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
	cronjob.AddFunc("@every 5s", job)
	cronjob.Start()

	fmt.Printf("CronJob entries: %d\n", len(cronjob.Entries()))

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	fmt.Println("Termination signal received!")

	cronjob.Stop()
}

func job() {
	fmt.Println("CronJob execution started")
	time.Sleep(3 * time.Second)
	fmt.Println("CronJob execution completed")
}
