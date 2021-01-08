package pagamento

type PagamentoServico interface {
	GerarRemessaBoleto(pagamento RemessaPagamentoBoleto) (string, error)
	//GerarRemessaTransferencia(pagamento RemessaPagamento) (string, error)
	LeituraRetorno(pathArquivo string) (RetornoPagamento, error)
}
