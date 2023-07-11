package models

import "fmt"

type ResourceModel struct {
	Model string
}

func (r *ResourceModel) Index() {
	fmt.Println("Indexing " + r.Model + " model")
}

func (r *ResourceModel) Show(id int) {
	fmt.Println("Showing " + r.Model + " model with id " + string(rune(id)))
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
