package interfaces

type ResourceInterface interface {
	Index()
	Show(id int)
	Store()
	Update(id int)
	Destroy(id int)
}
