package exchangecodes

type ExchangeCodeRepository interface {
	Create(ec *ExchangeCode) error
	Redeem(code string) (*ExchangeCode, error)
}
