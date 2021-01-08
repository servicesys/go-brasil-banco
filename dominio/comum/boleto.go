package comum

import (
	"time"
)

type Boleto struct {
	ID                uint64
	Cedente           ContaMovimentacao
	Pagador           Pagador
	Valor             float64
	ValorIOF          float64
	ValorDesconto     float64
	DataDocumento     time.Time
	DataProcessamento time.Time
	DataVencimento    time.Time
	Numero            int64
	NossoNumero       string
	CampoLivre        string
	CodigoBarras      string
	LinhaDigitavel    string
}
