package paymentmethod

type Repository interface {
	Create(paymentMethod *PaymentMethod) error
}
