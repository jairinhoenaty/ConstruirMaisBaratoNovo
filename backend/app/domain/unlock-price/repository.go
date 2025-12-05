package unlockprice

type UnlockPriceRepository interface {
	FindByUserType(userType UserType) (*UnlockPrice, error)
	FindActiveByUserType(userType UserType) (*UnlockPrice, error)
	FindAll() ([]*UnlockPrice, error)
	Save(unlockPrice UnlockPrice) (*UnlockPrice, error)
	Update(unlockPrice *UnlockPrice) error
}
