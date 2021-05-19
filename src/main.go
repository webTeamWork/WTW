package main

import (
	"forum/src/model"
	"forum/src/router"
)

func main() {
	model.OpenDatabase()
	router.RunAPP()
}
