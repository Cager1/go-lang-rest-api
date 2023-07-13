package main

import (
	"go-lang-rest-api/models"
)

func main() {
	user := models.NewUser()
	user.Show(1)
	user.Index()
}
