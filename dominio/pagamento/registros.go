package pagamento

import (
	"github.com/google/uuid"
	"go-brasil-banco/dominio/comum"
	"time"
)

type RemessaPagamentoBoleto struct {
	UUID uuid.UUID
	ContaMovimentacao  comum.ContaMovimentacao
	DataGeracao time.Time
	Registros []RegistroPagamentoRemessa
}

type RegistroPagamentoRemessa struct {
	ID                 string
	Sequencial         int
}


type RetornoPagamento struct {
	comum.ContaMovimentacao
	Registros []RegistroPagamentoRetorno
}

type RegistroPagamentoRetorno struct {
	ID                 string
	Sequencial         int
}