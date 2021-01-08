package comum

type Pagador struct {
	ID               string
	Nome            string
	Endereco        Endereco
	NumeroDocumento string
	TipoPagador     int
}
