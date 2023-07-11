package main

import (
	"go-lang-rest-api/models"
)

func main() {
	user := models.NewUser()
	user.Index()
	client := models.NewClient()
	client.Index()
}
