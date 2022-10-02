package main

import (
	"github.com/pratikdaigavane/emoji-hash/resources"
)

func main() {
	resources.Connect()
	//err := resources.Err
	//session := resources.Session
	defer resources.Close()
}
