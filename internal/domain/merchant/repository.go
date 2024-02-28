package merchant

type Repository interface {
	Create(merchant *Merchant) error
}
