package models

type Client struct {
	ResourceModel
}

func NewClient() Client {
	return Client{
		ResourceModel: ResourceModel{
			Model: "clients",
		},
	}
}
