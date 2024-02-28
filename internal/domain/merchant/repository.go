package merchant

type Repository interface {
	Create(merchant *Merchant) error
	GetByApiKey(apiKey string) (Merchant, error)
}
