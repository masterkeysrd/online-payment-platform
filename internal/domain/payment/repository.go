package payment

type Repository interface {
	Get(merchantID string, paymentID string) (Payment, error)
	Create(request *Payment) error
	Update(request *Payment) error
}
