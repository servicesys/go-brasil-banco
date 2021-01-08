package cobranca

type CobrancaServico interface {
	GerarRemessa(remessaCobranca RemessaCobranca) (string, error)
	LeituraRetorno(pathArquivo string, nomeArquivo string) (RetornoCobranca, error)
}
