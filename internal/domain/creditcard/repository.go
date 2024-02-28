package creditcard

type Repository interface {
	Create(creditCard *CreditCard) error
}
