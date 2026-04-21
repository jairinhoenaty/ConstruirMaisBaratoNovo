package exchangecodes

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type ExchangeCodeService interface {
	Generate(userID uint, ttl time.Duration) (*ExchangeCode, error)
	Redeem(code string) (*ExchangeCode, error)
}

type exchangeCodeService struct {
	repository ExchangeCodeRepository
}

func NewExchangeCodeService(repository ExchangeCodeRepository) ExchangeCodeService {
	return &exchangeCodeService{
		repository: repository,
	}
}

func (s *exchangeCodeService) Generate(userID uint, ttl time.Duration) (*ExchangeCode, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}

	ec := &ExchangeCode{
		Code:      hex.EncodeToString(bytes),
		UserID:    userID,
		ExpiresAt: time.Now().Add(ttl),
	}

	if err := s.repository.Create(ec); err != nil {
		return nil, err
	}
	return ec, nil
}

func (s *exchangeCodeService) Redeem(code string) (*ExchangeCode, error) {
	return s.repository.Redeem(code)
}
