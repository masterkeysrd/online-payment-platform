package customer

type Repository interface {
	Create(customer *Customer) error
}
