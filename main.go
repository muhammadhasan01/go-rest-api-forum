package main

import (
	"backend-forum/router"
	"backend-forum/utils"
)

func main() {
	utils.PrepareLog()
	router.StartAPI()
}
