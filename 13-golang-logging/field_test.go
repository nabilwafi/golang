package main_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "nabilwafi",
		"name":     "Nabil Wafi",
	}).Info("Hello Logger")
}
