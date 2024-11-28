package main

import (
	"os"

	"github.com/rezairfanwijaya/go-autoscale.git/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	cli := cmd.New()
	_ = cli.Execute()
}
