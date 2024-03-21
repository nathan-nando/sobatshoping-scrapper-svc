package main

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	jakartaTIme, _ := time.LoadLocation("Asia/Jakarta")
	
	logrus.Info(time.Now().Second())
	scheduler := cron.New(cron.WithLocation(jakartaTIme))

	defer scheduler.Stop()

	var counter int

	_, err := scheduler.AddFunc("*/1 * * * *", func() {
		counter++
		logrus.Infof("calling job %d time", counter)
		doSomething("This is job1")
	})

	if err != nil {
		logrus.Error("error :", err)
	}

	logrus.Info(scheduler.Entries())

	logrus.Info("start cron")
	go scheduler.Start()

	time.Sleep(time.Minute * 10)

}

func doSomething(text string) {
	logrus.Info("Hi from cron: ", text)
}
