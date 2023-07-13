package models

import (
	"fmt"
	"go-lang-rest-api/helpers"
)

type ResourceModel struct {
	Model string
	helpers.ResourceHelper
}

func (r *ResourceModel) Index() {
	fmt.Println(r)
	fmt.Println("Indexing " + r.Model + " model")
}

func (r *ResourceModel) Show(id int) {
	r.Find(id, r.Model)
	fmt.Println("Showing", r.Model, "model with id", id)
}

func (r *ResourceModel) Store() {
	fmt.Println("Creating " + r.Model + " model")
}

func (r *ResourceModel) Update(id int) {
	fmt.Println("Updating " + r.Model + " model with id " + string(rune(id)))
}

func (r *ResourceModel) Destroy(id int) {
	fmt.Println("Destroying " + r.Model + " model with id " + string(rune(id)))
}
