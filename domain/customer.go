package domain

type Customer struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
