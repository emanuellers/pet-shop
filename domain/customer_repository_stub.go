package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

// Mock function
func (stub CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return stub.customers, nil
}

func NewCustomersRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"001", "Example Name", "Active"},
	}

	return CustomerRepositoryStub{customers: customers}
}
