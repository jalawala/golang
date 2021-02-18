package main

import (
	log1 "github.com/sirupsen/logrus"
	"log"
)

func main() {

	log.Println("line 1")
	log.Println("line 2")
	log.Println("line 3")

	log1.Info("line 1")
	log1.Warn("line 2")
	log1.Fatal("line 3")

}
