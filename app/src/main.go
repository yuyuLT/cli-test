package main

import (
	"mvc_test/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run()
}