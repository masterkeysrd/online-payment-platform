package paymentmethod

type Repository interface {
	Create(paymentMethod *PaymentMethod) error
	Get(merchantID, customerID string, paymentMethodID string) (*PaymentMethod, error)
}
