package pagamento

type PagamentoServico interface {
	GerarRemessaBoleto(pagamento RemessaPagamentoBoleto) (string, error)
	GerarRemessaTransferencia(transferencia RemessaTransferencia) (string, error)
	LeituraRetorno(pathArquivo string) (RetornoPagamento, error)
}
