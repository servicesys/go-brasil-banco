package comum

type ContaMovimentacao struct {
	Banco               string
	Agencia             string
	AgenciaDigito       string
	ContaCorrente       string
	ContaCorrenteDigito string
	Carteira            string
	Convenio            string
	NomeTitular         string
	DocTitular          string
	TipoTitular         int
	CodigoEmpresa       string
}


type ContaCobranca struct {
	Banco               string
	Agencia             string
	ContaCorrente       string
	ContaDigito         string
	Carteira            string
}


