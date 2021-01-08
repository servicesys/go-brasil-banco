package cobranca

import (
	"github.com/google/uuid"
	"go-brasil-banco/dominio/comum"
	"time"
)

type RemessaCobranca struct {
	UUID        string
	Conta       comum.ContaMovimentacao
	Registros   []RegistroCobrancaRemessa
	DataGeracao time.Time
	Sequencial  int64
}

func (remessaCobranca *RemessaCobranca) AddRegistro(registro RegistroCobrancaRemessa) {

	remessaCobranca.Registros = append(remessaCobranca.Registros, registro)
}

type RegistroCobrancaRemessa struct {
	ID         int64
	Sequencial int
	Boleto     comum.Boleto
}

type RetornoCobranca struct {
	UUID          uuid.UUID
	NomeArquivo   string
	DataGeracao   time.Time
	CodigoEmpresa string
	NomeEmpresa   string
	Registros     []RegistroRetorno
}

type RegistroRetorno struct {
	ID             string
	TipoCobranca   string
	DataOcorrencia time.Time
	Valor          float64
	DataVencimento time.Time
	Numero         string
	NossoNumero    string
	ValorTarifa    float64
	Ocorrencia     string
	TipoOcorrencia string
	EspecieTitulo      string
	OutrasDespesas     float64
	JurosAtraso        float64
	TaxaIOF            float64
	ValorAbatimento    float64
	DescontoConcedido  float64
	ValorRecebido      float64
	JurosMora          float64
	OutrosRecebimentos float64
	DataCredito        time.Time
	Sequencial         int64
	ContaCobranca comum.ContaCobranca
}
