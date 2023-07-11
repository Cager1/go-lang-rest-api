package models

import "fmt"

type User struct {
	ResourceModel
}

func NewUser() User {
	return User{
		ResourceModel: ResourceModel{
			Model: "users",
		},
	}
}

func (r *User) Index() {
	fmt.Println("indexing before using main function " + r.Model + " model")
	r.ResourceModel.Index()
}
