package main_test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}