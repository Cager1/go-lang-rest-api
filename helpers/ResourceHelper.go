package helpers

import (
	"fmt"
)

type ResourceHelper struct {
}

func (r *ResourceHelper) Find(id int, model string) {
	fmt.Printf("Finding %s model with id %d\n", model, id)
}

func (r *ResourceHelper) All(model string) {
	fmt.Printf("Getting all %s model\n", model)
}

func (r *ResourceHelper) Create(model string) {
	fmt.Printf("Creating %s model\n", model)
}

func (r *ResourceHelper) Update(id int, model string) {
	fmt.Printf("Updating %s model with id %d\n", model, id)
}

func (r *ResourceHelper) Delete(id int, model string) {
	fmt.Printf("Deleting %s model with id %d\n", model, id)
}
