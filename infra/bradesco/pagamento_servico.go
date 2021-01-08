package bradesco

import (
	"fmt"
	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"go-brasil-banco/dominio/comum"
	"go-brasil-banco/dominio/pagamento"
	"log"
	"os"
	"strings"
	"time"
)

type PagamentoBradescoServico struct {
}

func NewPagamentoBradescoServico() pagamento.PagamentoServico {

	return &PagamentoBradescoServico{}
}

func (pagamentoBradescoServico PagamentoBradescoServico) GerarRemessaBoleto(remessaPagamentoBoleto pagamento.RemessaPagamentoBoleto) (string, error) {

	source := strings.NewReader(CNAB240PagamentosBRADESCO)
	layout, _ := superfile.NewLayout(source)
	remessa := superfile.NewRemessa(layout)

	remessa.Header["codigo_banco"] = remessaPagamentoBoleto.ContaMovimentacao.Banco

	remessa.Header["codigo_lote"] = 0000
	remessa.Header["tipo_registro"] = 0

	remessa.Header["tipo_inscricao"] = remessaPagamentoBoleto.ContaMovimentacao.TipoTitular
	remessa.Header["inscricao_numero"] = remessaPagamentoBoleto.ContaMovimentacao.DocTitular
	remessa.Header["codigo_convenio_banco"] = remessaPagamentoBoleto.ContaMovimentacao.Convenio // "348828"
	remessa.Header["agencia_debito"] = remessaPagamentoBoleto.ContaMovimentacao.Agencia
	remessa.Header["agencia_digito"] = remessaPagamentoBoleto.ContaMovimentacao.AgenciaDigito

	remessa.Header["conta_debito"] = remessaPagamentoBoleto.ContaMovimentacao.ContaCorrente
	remessa.Header["conta_digito"] = remessaPagamentoBoleto.ContaMovimentacao.ContaCorrenteDigito
	//remessa.Header["digito_ver_conta_agencia"] = "0"

	remessa.Header["nome_empresa"] = remessaPagamentoBoleto.ContaMovimentacao.NomeTitular
	remessa.Header["nome_banco"] = "BANCO BRADESCO S.A"
	remessa.Header["data_geracao"] = 5122020
	remessa.Header["hora_geracao"] = 150007
	remessa.Header["numero_sequencial_arquivo"] = 700048
	remessa.Header["densidade_gravacao"] = 1600 //Densidade de Gravação do Arquivo pag 84

	lote := remessa.NovoLote()

	lote.Header["codigo_lote"] = 1 //- a cada novo lote +1

	lote.Header["tipo_registro"] = 1

	//lote.Header["tipo_operacao"] = 'C'
	lote.Header["tipo_servico"] = 20
	//20'= Pagamento Fornecedor - pag 85 manual

	lote.Header["forma_lancamento"] = 31
	lote.Header["layout_lote"] = 40

	//'03' = DOC/TED(1) pag 85 manual
	//'03' = DOC/TED(1) pag 85 manual
	//11’ = Pagamento de Contas e Tributos comCódigo de Barras

	lote.Header["tipo_inscricao_empresa"] = remessaPagamentoBoleto.ContaMovimentacao.TipoTitular
	lote.Header["numero_inscricao"] = remessaPagamentoBoleto.ContaMovimentacao.DocTitular
	lote.Header["codigo_convenio_banco"] = remessaPagamentoBoleto.ContaMovimentacao.Convenio

	lote.Header["identificacao_lancamento"] = "0"
	lote.Header["agencia_debito"] = remessaPagamentoBoleto.ContaMovimentacao.Agencia
	lote.Header["agencia_debito_digito"] = remessaPagamentoBoleto.ContaMovimentacao.AgenciaDigito
	lote.Header["conta_debito"] = remessaPagamentoBoleto.ContaMovimentacao.ContaCorrente
	lote.Header["conta_debito_digito"] = remessaPagamentoBoleto.ContaMovimentacao.ContaCorrenteDigito

	lote.Header["nome_empresa"] = remessaPagamentoBoleto.ContaMovimentacao.NomeTitular
	lote.Header["finalidade_lote"] = "0"
	lote.Header["endereco_empresa"] = "Vinicius de Moraes"
	lote.Header["numero"] = 145
	lote.Header["cidade"] = "VICOSA"
	lote.Header["cep"] = 36570
	lote.Header["cep_complemento"] = 19
	lote.Header["estado"] = "MG"

	lote.Header["numero_sequencial_arquivo"] = 1

	//BOLETO PAGAR 01

	//pagamento de boleto
	//boleto  J

	detalhe2 := lote.NovoDetalhe()
	detalhe2["segmento_j"]["codigo_banco"] = 237
	detalhe2["segmento_j"]["codigo_lote"] = 1
	detalhe2["segmento_j"]["numero_registro"] = 1
	detalhe2["segmento_j"]["tipo_movimento"] = 000
	detalhe2["segmento_j"]["nome_cedente"] = "CENDENTE DO BOLETO "

	var codigoBarra = comum.CodigoBarras{Banco: "756",
		Valor:          269.21,
		DataVencimento: time.Date(2020, time.Month(12), 10, 0, 0, 0, 0, time.UTC),
		CampoLivre:     "0000002962087000000000117",
	}

	println(codigoBarra.String())
	detalhe2["segmento_j"]["codigo_barra"] = codigoBarra.String()
	detalhe2["segmento_j"]["valor_titulo"] = 269.21
	detalhe2["segmento_j"]["valor_desconto"] = 0
	detalhe2["segmento_j"]["valor_pagamento"] = 269.21
	detalhe2["segmento_j"]["data_vencimento"] = 10122020
	detalhe2["segmento_j"]["data_pagamento"] = 10122020
	detalhe2["segmento_j"]["codigo_moeda"] = "09"

	//J52
	detalhe2["segmento_j_52"]["codigo_lote"] = 1
	detalhe2["segmento_j_52"]["numero_registro"] = 2
	//detalhe2["segmento_j_52"]["movimento_codigo"] = 2

	//sacado
	detalhe2["segmento_j_52"]["sacado_tipo_inscricao"] = 1
	detalhe2["segmento_j_52"]["sacado_numero_inscricao"] = 88967140525
	detalhe2["segmento_j_52"]["sacado_nome"] = "VALTER DE O. LOBO"

	//cedente
	detalhe2["segmento_j_52"]["cedente_tipo_inscricao"] = 2
	detalhe2["segmento_j_52"]["cedente_numero_inscricao"] = 9987222000178
	detalhe2["segmento_j_52"]["cedente_nome"] = "Centro de Educação Fenix Ltda-ME"

	//sacador
	detalhe2["segmento_j_52"]["sacador_tipo_inscricao"] = 2
	detalhe2["segmento_j_52"]["sacador_numero_inscricao"] = 9987222000178
	detalhe2["segmento_j_52"]["sacador_nome"] = "Centro de Educação Fenix Ltda-ME"

	lote.InserirDetalhe(detalhe2)

	//BOLETO 01

	lote.Trailer["codigo_lote"] = lote.Sequencial
	lote.Trailer["total_registros_lote"] = 4 //numero de registros  no lote
	lote.Trailer["total_valor_pagtos"] = 269.21

	remessa.InserirLote(lote)

	println(remessa.TotalRegistros())
	println(remessa.TotalLotes())
	println(remessa.Lotes[0].Segmentos())
	remessa.Trailer["total_lotes_arquivo"] = remessa.TotalLotes()
	remessa.Trailer["total_registros"] = remessa.TotalRegistros() - 1
	remessa.Trailer["total_contas_cobranca"] = 0
	remessaFile := file.NewRemessaFile(remessa, "bradesco-pagamentos-cnab240.rem")

	arquivo := remessaFile.Write()
	log.Println(arquivo)
	return "path bradesco remesssa", nil
}

func (pagamentoBradescoServico PagamentoBradescoServico) GerarRemessaTransferencia(transferencia pagamento.RemessaTransferencia) (string, error) {
	return "path bradesco remesssa", nil
}

func (pagamentoBradescoServico PagamentoBradescoServico) LeituraRetorno(pathArquivo string) (pagamento.RetornoPagamento, error) {

	source := strings.NewReader(CNAB240PagamentosBRADESCO)
	layout, _ := superfile.NewLayout(source)

	f, _ := os.Open("PI081000.RST")
	defer f.Close()
	arquivo, _ := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	fmt.Println(retorno.Header)

	totalLotes := retorno.TotalLotes()
	fmt.Println(totalLotes)
	var idx int64 = 0
	for ; idx < totalLotes; idx++ {

		segmentos := retorno.Lotes[idx].Segmentos()

		//detalhe :=retorno.Lotes[idx].NovoDetalhe()

		//retorno.Lotes[idx].InserirDetalhe(detalhe)

		for _, element := range segmentos {
			fmt.Println("----------------")
			//fmt.Println(index)
			fmt.Println(element.Valores["segmento_codigo"])

			if "A" == element.Valores["segmento_codigo"] {

				fmt.Println(element.Valores)

				fmt.Println(element.Valores["codigo_ocorrencias"])
				fmt.Println(element.Valores["data_pagto"])
				fmt.Println(element.Valores["valor_pagto"])
				fmt.Println(element.Valores["nosso_numero"])
				fmt.Println(element.Valores["numero_registro"])
				fmt.Println(element.Valores["nome_favorecido"])
				fmt.Println(element.Valores["nosso_numero"])
				fmt.Println(element.Valores["numero_doc"])
				fmt.Println(element.Valores["nome_favorecido"])
			}

			if "J" == element.Valores["segmento_codigo"] {

				fmt.Println(element.Valores)
				fmt.Println(element.Valores["codigo_ocorrencias"])
				fmt.Println(element.Valores["valor_titulo"])
				fmt.Println(element.Valores["valor_pagamento"])
				fmt.Println(element.Valores["nome_cedente"])
				fmt.Println(element.Valores["data_pagamento"])
				fmt.Println(element.Valores["data_vencimento"])
				fmt.Println(element.Valores["codigo_documento_empresa"])
				fmt.Println(element.Valores["codigo_barra"])

			}

			//fmt.Println(element.Valores["numero_sequencial_registro"])
			/*for k , v  := range element.Valores {
				fmt.Println(k)
				fmt.Println(v)
			}*/

			fmt.Println("----------------")
		}

		//trailerr lote
		traillerLoteRegistro := retorno.Lotes[idx].Trailer
		fmt.Println(traillerLoteRegistro)

	}
	return pagamento.RetornoPagamento{}, nil
}
