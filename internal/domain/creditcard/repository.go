package creditcard

type Repository interface {
	Create(creditCard *CreditCard) error
	GetByPaymentMethod(merchantID, paymentMethodID string) (*CreditCard, error)
}
