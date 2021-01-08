package bradesco

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"go-brasil-banco/dominio/cobranca"
	"go-brasil-banco/dominio/comum"
	"os"
	"strconv"
	"strings"
	"time"
)

type CobrancaBradescoServico struct {
}

func NewCobrancaBradescoServico() cobranca.CobrancaServico {

	return &CobrancaBradescoServico{}
}

func (CobrancaBradescoServico CobrancaBradescoServico) GerarRemessa(remessaCobranca cobranca.RemessaCobranca) (string, error) {

	source := strings.NewReader(CNAB400CobrancaBRADESCO)

	layout, _ := superfile.NewLayout(source)
	remessa := superfile.NewRemessa(layout)

	remessa.Header["identifica_registro"] = 0
	remessa.Header["identifica_arquivo_remessa"] = 1
	remessa.Header["literal_remessa"] = "REMESSA"
	remessa.Header["codigo_servico"] = "01"
	remessa.Header["literal_servico"] = "COBRANCA"
	remessa.Header["cod_empresa"] = remessaCobranca.Conta.CodigoEmpresa
	remessa.Header["nome_empresa"] = remessaCobranca.Conta.NomeTitular
	remessa.Header["codigo_banco"] = comum.CodigoBRADESCO
	remessa.Header["nome_banco"] = "BRADESCO"
	remessa.Header["data_geracao"] = comum.DataFormatada(remessaCobranca.DataGeracao)
	remessa.Header["identificacao_sistema"] = "MX"
	remessa.Header["numero_sequencial_remessa"] = remessaCobranca.Sequencial
	remessa.Header["numero_sequencial_registro"] = 1

	lote := remessa.NovoLote()

	for idx, registroRemessa := range remessaCobranca.Registros {

		sequencial := idx + 2
		boleto := registroRemessa.Boleto
		detalhe := lote.NovoDetalhe()
		detalhe["tipo_1"]["identifica_registro"] = 1
		detalhe["tipo_1"]["identificacao_empresa_beneficiaria_carteira"] = comum.Zeros(remessaCobranca.Conta.Carteira, 3)
		detalhe["tipo_1"]["identificacao_empresa_beneficiaria_cod_agencia"] = comum.Zeros(remessaCobranca.Conta.Agencia, 5)
		detalhe["tipo_1"]["identificacao_empresa_beneficiaria_conta"] = comum.Zeros(remessaCobranca.Conta.ContaCorrente, 7)
		detalhe["tipo_1"]["identificacao_empresa_beneficiaria_dig_conta"] = remessaCobranca.Conta.ContaCorrenteDigito
		detalhe["tipo_1"]["num_participante_cliente"] = boleto.Pagador.ID

		detalhe["tipo_1"]["nome_pagador"] = boleto.Pagador.Nome
		detalhe["tipo_1"]["valor"] = boleto.Valor
		detalhe["tipo_1"]["cep"], _ = strconv.Atoi(boleto.Pagador.Endereco.Cep[0:5])
		detalhe["tipo_1"]["cep_sufixo"], _ = strconv.Atoi(boleto.Pagador.Endereco.Cep[5:8])
		detalhe["tipo_1"]["endereco_completo"] = comum.StrLimit(boleto.Pagador.Endereco.Rua, 40)
		detalhe["tipo_1"]["codigo_inscricao_pagador"] = boleto.Pagador.TipoPagador
		detalhe["tipo_1"]["numero_inscricao_pagador"] = boleto.Pagador.NumeroDocumento
		detalhe["tipo_1"]["identificacao_titulo_banco"] = comum.Zeros(strconv.FormatUint(boleto.ID, 10), 11)
		dv, errorDV := CalculoDigitoNossoNumero(comum.Zeros(remessaCobranca.Conta.Carteira, 1), comum.Zeros(strconv.FormatUint(boleto.ID, 10), 12))
		if errorDV != nil {
			return "", errorDV
		}
		detalhe["tipo_1"]["digito_auto_conf_num_bancario"] = dv
		detalhe["tipo_1"]["valor_iof"] = boleto.ValorIOF
		detalhe["tipo_1"]["valor_desconto"] = boleto.ValorDesconto
		detalhe["tipo_1"]["data_emissao"] = comum.DataFormatada(boleto.DataProcessamento)
		detalhe["tipo_1"]["vencimento"] = comum.DataFormatada(boleto.DataVencimento)
		detalhe["tipo_1"]["quantidade_pagamentos"] = "01"
		detalhe["tipo_1"]["especie"] = 99 //outros
		detalhe["tipo_1"]["instrucao_1"] = 14
		detalhe["tipo_1"]["numero_sequencial_registro"] = sequencial
		detalhe["tipo_1"]["codigo_ocorrencia"] = 01

		lote.InserirDetalhe(detalhe)
	}
	remessa.InserirLote(lote)

	remessa.Trailer["total_lotes_arquivo"] = remessa.TotalLotes()
	remessa.Trailer["total_registros"] = remessa.TotalRegistros()

	//@TODO - GERAR NOME ARQUIVO
	fileRemessaPath := "CB031220.REM"
	fileRemessa := file.NewRemessaFile(remessa, fileRemessaPath).Write()

	//@TODO - retornar  *os.file
	return fileRemessa.Name(), nil
}

func (CobrancaBradescoServico CobrancaBradescoServico) LeituraRetorno(pathArquivo string, nomeArquivo string) (cobranca.RetornoCobranca, error) {

	retornoCobranca := cobranca.RetornoCobranca{}

	source := strings.NewReader(CNAB400CobrancaBRADESCO)
	layout, _ := superfile.NewLayout(source)

	f, _ := os.Open(pathArquivo + nomeArquivo)
	defer f.Close()
	arquivo, _ := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	var dataGeracao time.Time = time.Time{}
	if retorno.Header["data_geracao"] != nil {
		strData := fmt.Sprintf("%v", retorno.Header["data_geracao"])
		dataGeracao = comum.DateStringTime(strData, 2000)
	}

	retornoCobranca.UUID = uuid.New()
	retornoCobranca.DataGeracao = dataGeracao
	retornoCobranca.NomeArquivo = nomeArquivo
	retornoCobranca.CodigoEmpresa = fmt.Sprintf("%v", retorno.Header["codigo_empresa"])
	retornoCobranca.NomeEmpresa = fmt.Sprintf("%v", retorno.Header["nome_empresa"])
	retornoCobranca.Registros = []cobranca.RegistroRetorno{}

	totalLotes := retorno.TotalLotes()

	var idx int64 = 0

	for ; idx < totalLotes; idx++ {

		segmentos := retorno.Lotes[idx].Segmentos()

		var valorRecebido float64
		for _, element := range segmentos {

			var dataCredito time.Time
			if element.Valores["data_credito"] != nil {
				strData := fmt.Sprintf("%v", element.Valores["data_credito"])
				dataCredito = comum.DateStringTime(strData, 2000)
			}

			var dataOcorrencia time.Time = time.Time{}
			if element.Valores["data_ocorrencia"] != nil {
				strData := fmt.Sprintf("%v", element.Valores["data_ocorrencia"])
				dataOcorrencia = comum.DateStringTime(strData, 2000)
			}

			var dataVencimento time.Time
			if element.Valores["vencimento"] != nil {
				strData := fmt.Sprintf("%v", element.Valores["vencimento"])
				dataVencimento = comum.DateStringTime(strData, 2000)
			}

			valorRecebido = 0.0
			if element.Valores["valor_principal"] != nil {
				valorRecebido = element.Valores["valor_principal"].(float64)
			}

			identificaoEmpresa := fmt.Sprintf("%v", element.Valores["identificacao_empresa_beneficiaria"])
			contaCobranca := comum.ContaCobranca{
				Banco:         fmt.Sprintf("%v", element.Valores["codigo_banco"]),
				Agencia:       identificaoEmpresa[5:9],
				ContaCorrente: identificaoEmpresa[10:16],
				ContaDigito:   identificaoEmpresa[17:17],
				Carteira:      identificaoEmpresa[2:4],
			}

			var registro = cobranca.RegistroRetorno{
				ID:                 "",
				TipoCobranca:       fmt.Sprintf("%v", element.Valores["identificacao_tipo_registro"]),
				DataOcorrencia:     dataOcorrencia,
				Valor:              element.Valores["valor_titulo"].(float64),
				DataVencimento:     dataVencimento,
				Numero:             fmt.Sprintf("%v", element.Valores["numero_documento"]),
				NossoNumero:        fmt.Sprintf("%v", element.Valores["identificacao_titulo_nosso_numero"]),
				ValorTarifa:        element.Valores["tarifa_cobranca"].(float64),
				Ocorrencia:         fmt.Sprintf("%v", element.Valores["codigo_ocorrencia"]),
				TipoOcorrencia:     fmt.Sprintf("%v", element.Valores["motivo_rejeicao_ocorrencia"]),
				DataCredito:        dataCredito,
				EspecieTitulo:      fmt.Sprintf("%v", element.Valores["especie"]),
				OutrasDespesas:     element.Valores["tarifa_outras_despesas"].(float64),
				JurosAtraso:        element.Valores["juros_operacao_atraso"].(float64),
				TaxaIOF:            element.Valores["valor_iof"].(float64),
				ValorAbatimento:    element.Valores["valor_abatimento"].(float64),
				DescontoConcedido:  element.Valores["descontos"].(float64),
				ValorRecebido:      valorRecebido,
				JurosMora:          element.Valores["juros_mora_multa"].(float64),
				OutrosRecebimentos: element.Valores["outros_creditos"].(float64),
				Sequencial:         element.Valores["numero_sequencial_registro"].(int64),
				ContaCobranca:      contaCobranca,
			}

			retornoCobranca.Registros = append(retornoCobranca.Registros, registro)
		}
	}

	return retornoCobranca, nil
}
