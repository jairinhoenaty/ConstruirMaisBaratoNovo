package payment_usecase

import (
	"sync"
	"time"
)

// pollThrottleTTL é o intervalo mínimo entre duas consultas ao MercadoPago
// para o mesmo pagamento.
//
// Fica deliberadamente abaixo do intervalo de polling do cliente (5s): se
// fossem iguais, qualquer variação de rede faria uma consulta legítima cair no
// bloqueio e a confirmação demoraria o dobro do tempo para aparecer na tela.
const pollThrottleTTL = 3 * time.Second

// pollThrottleMaxEntries limita a memória usada pelo mapa antes de uma limpeza.
const pollThrottleMaxEntries = 1000

// PollThrottle evita que o endpoint público de status vire um amplificador de
// requisições contra a API do MercadoPago.
//
// O endpoint não é autenticado: sem isso, alguém repetindo a chamada em laço
// faria o backend consultar o MercadoPago na mesma cadência e o nosso token
// acabaria limitado por rate limit.
//
// É compartilhado entre requisições, então precisa ser seguro para uso
// concorrente. O estado é apenas uma otimização: perdê-lo em um restart não
// causa problema algum.
type PollThrottle struct {
	mu         sync.Mutex
	ttl        time.Duration
	lastChecke map[int64]time.Time
}

func NewPollThrottle() *PollThrottle {
	return NewPollThrottleWithTTL(pollThrottleTTL)
}

// NewPollThrottleWithTTL permite ajustar o intervalo mínimo. TTL zero libera
// todas as consultas — útil em testes.
func NewPollThrottleWithTTL(ttl time.Duration) *PollThrottle {
	return &PollThrottle{
		ttl:        ttl,
		lastChecke: make(map[int64]time.Time),
	}
}

// ShouldCheck informa se já passou tempo suficiente para consultar de novo o
// pagamento, e registra a consulta quando autoriza.
func (p *PollThrottle) ShouldCheck(paymentID int64) bool {
	if p == nil {
		return true
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()
	if last, ok := p.lastChecke[paymentID]; ok && now.Sub(last) < p.ttl {
		return false
	}

	if len(p.lastChecke) >= pollThrottleMaxEntries {
		p.pruneLocked(now)
	}

	p.lastChecke[paymentID] = now
	return true
}

// pruneLocked descarta entradas velhas o bastante para não influenciarem mais
// nenhuma decisão. Deve ser chamada com o mutex já travado.
func (p *PollThrottle) pruneLocked(now time.Time) {
	for paymentID, last := range p.lastChecke {
		if now.Sub(last) >= p.ttl {
			delete(p.lastChecke, paymentID)
		}
	}
}
