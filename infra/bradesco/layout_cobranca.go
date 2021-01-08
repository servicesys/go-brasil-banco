package bradesco


const CNAB400CobrancaBRADESCO = `
# FORMATO: BRADESCO - CNAB400
# OBJETIVO DO ARQUIVO: COBRANÇA BANCÁRIA
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 400 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita.
# 21 a 21 -Zero.22 a 24 -Códigosda Carteira.25 a 29 -Códigos da Agência Beneficiários, sem o Dígito.30 a 36 -Contas-Corrente.37 a 37 -Dígitosda Conta.
#https://banco.bradesco/assets/pessoajuridica/pdf/4008-524-0121-layout-cobranca-versao-portugues.pdf

servico: 'cobranca'
versao: 'MARC2020'
layout: 'cnab400'

remessa:
  header_arquivo:
    identifica_registro:
      pos: [ 1,1 ]
      picture: '9(1)'
      default: 0
    identifica_arquivo_remessa:
      pos: [ 2,2 ]
      picture: '9(1)'
      default: 1
    literal_remessa:
      pos: [ 3,9 ]
      picture: 'X(7)'
      default: 'REMESSA'
    codigo_servico:
      pos: [ 10,11 ]
      picture: '9(2)'
      default: '01'
    literal_servico:
      pos: [ 12,26 ]
      picture: 'X(15)'
      default: 'COBRANCA'
    cod_empresa:
      pos: [ 27,46 ]
      picture: '9(20)'
    nome_empresa:
      pos: [ 47,76 ]
      picture: 'X(30)'
    codigo_banco:
      pos: [ 77,79 ]
      picture: '9(3)'
      default: 237
    nome_banco:
      pos: [ 80,94 ]
      picture: 'X(15)'
      default: 'BRADESCO'
    data_geracao:
      pos: [ 95,100 ]
      picture: '9(6)'
    brancos_02:
      pos: [ 101,108 ]
      picture: 'X(8)'
      default: ''
    identificacao_sistema:
      pos: [ 109,110 ]
      picture: 'X(2)'
      default: 'MX'
    numero_sequencial_remessa:
      pos: [ 111,117 ]
      picture: '9(7)'
    brancos_03:
      pos: [ 118,394 ]
      picture: 'X(277)'
      default: ''
    numero_sequencial_registro:
      pos: [ 395,400 ]
      picture: '9(6)'
      default: '000001'

  trailer_arquivo:
    tipo_registro:
      pos: [ 1,1 ]
      picture: '9(1)'
      default: 9
    brancos:
      pos: [ 2,394 ]
      picture: 'X(393)'
      default: ''
    numero_sequencial_registro:
      pos: [ 395,400 ]
      picture: '9(6)'

  detalhes:
    tipo_1:
      identifica_registro:
        pos: [ 1,1 ]
        picture: '9(1)'
        default: 1
      agencia_debito:
        pos: [ 02,06 ]
        picture: '9(5)'
        default: 0
      digito_agencia_debito:
        pos: [ 07,07 ]
        picture: 'X(1)'
        default: '0'
      razao_conta_corrente:
        pos: [ 08,12 ]
        picture: 'X(5)'
        default: '00000'
      conta_corrente:
        pos: [ 13,19 ]
        picture: '9(7)'
        default: 0
      digito_conta_corrente:
        pos: [ 20,20 ]
        picture: 'X(1)'
        default: '0'
      identificacao_empresa_beneficiaria_zero:
        pos: [ 21,21 ]
        picture: 'X(1)'
        default: '0'
      identificacao_empresa_beneficiaria_carteira:
        pos: [ 22,24 ]
        picture: 'X(3)'
        default: 0
      identificacao_empresa_beneficiaria_cod_agencia:
        pos: [ 25,29 ]
        picture: 'X(5)'
        default: 0
      identificacao_empresa_beneficiaria_conta:
        pos: [ 30,36 ]
        picture: 'X(7)'
        default: 0
      identificacao_empresa_beneficiaria_dig_conta:
        pos: [ 37,37 ]
        picture: 'X(1)'
      num_participante_cliente:
        pos: [ 38,62 ]
        picture: 'X(25)'
      cod_banco_compensacao:
        pos: [ 63,65 ]
        picture: '9(3)'
        default: 237
      campo_multa:
        pos: [ 66,66 ]
        picture: '9(1)'
        default: 0
      percentual_multa:
        pos: [ 67,70 ]
        picture: '9(4)'
        default: 0
      identificacao_titulo_banco:
        pos: [ 71,81 ]
        picture: '9(11)'
      digito_auto_conf_num_bancario:
        pos: [ 82,82 ]
        picture: 'X(1)'
      desconto_bonificacao_dia:
        pos: [ 83,92 ]
        picture: '9(10)'
      condicao_emissao_papeleta:
        pos: [ 93,93 ]
        picture: '9(1)'
        default: 2
      identificacao_boleto_deb_automatico:
        pos: [ 94,94 ]
        picture: 'X(1)'
        default: 'N'
      identificacao_operacao_banco:
        pos: [ 95,104 ]
        picture: 'X(10)'
        default: ''
      identificacao_rateio_credito:
        pos: [ 105,105 ]
        picture: 'X(1)'
        default: ''
      enderecamento_aviso_debito:
        pos: [ 106,106 ]
        picture: '9(1)'
      quantidade_pagamentos:
        pos: [ 107,108 ]
        picture: 'X(2)'
      codigo_ocorrencia:
        pos: [ 109,110 ]
        picture: '9(2)'
      numero_documento:
        pos: [ 111,120 ]
        picture: 'X(10)'
      vencimento:
        pos: [ 121,126 ]
        picture: '9(6)'
      valor:
        pos: [ 127,139 ]
        picture: '9(11)V9(2)'
      codigo_banco_cobranca:
        pos: [ 140,142 ]
        picture: '9(3)'
        default: 0
      agencia_cobradora:
        pos: [ 143,147 ]
        picture: '9(5)'
      especie:
        pos: [ 148,149 ]
        picture: '9(2)'
      aceite:
        pos: [ 150,150 ]
        picture: 'X(1)'
        default: 'N'
      data_emissao:
        pos: [ 151,156 ]
        picture: '9(6)'
      instrucao_1:
        pos: [ 157,158 ]
        picture: '9(2)'
      instrucao_2:
        pos: [ 159,160 ]
        picture: '9(2)'
      juros_1_dia:
        pos: [ 161,173 ]
        picture: '9(11)V9(2)'
      desconto_ate:
        pos: [ 174,179 ]
        picture: '9(6)'
      valor_desconto:
        pos: [ 180,192 ]
        picture: '9(11)V9(2)'
      valor_iof:
        pos: [ 193,205 ]
        picture: '9(11)V9(2)'
      abatimento:
        pos: [ 206,218 ]
        picture: '9(11)V9(2)'
      codigo_inscricao_pagador:
        pos: [ 219,220 ]
        picture: '9(2)'
      numero_inscricao_pagador:
        pos: [ 221,234 ]
        picture: '9(14)'
      nome_pagador:
        pos: [ 235,274 ]
        picture: 'X(40)'
      endereco_completo:
        pos: [ 275,314 ]
        picture: 'X(40)'
      messagem_01:
        pos: [ 315,326 ]
        picture: 'X(12)'
        default: ''
      cep:
        pos: [ 327,331 ]
        picture: '9(5)'
      cep_sufixo:
        pos: [ 332,334 ]
        picture: '9(3)'
      messagem_02:
        pos: [ 335,394 ]
        picture: 'X(60)'
        default: ''
      numero_sequencial_registro:
        pos: [ 395,400 ]
        picture: '9(6)'

retorno:
  header_arquivo:
    identifica_registro:
      pos: [ 1,1 ]
      picture: '9(1)'
      default: 0
    identifica_retorno:
      pos: [ 2,2 ]
      picture: '9(1)'
      default: 2
    literal_retorno:
      pos: [ 3,9 ]
      picture: 'X(7)'
      default: 'RETORNO'
    codigo_servico:
      pos: [ 10,11 ]
      picture: '9(2)'
      default: '01'
    literal_servico:
      pos: [ 12,26 ]
      picture: 'X(15)'
      default: 'COBRANCA'
    codigo_empresa:
      pos: [ 27,46 ]
      picture: '9(20)'
    nome_empresa:
      pos: [ 47,76 ]
      picture: 'X(30)'
    num_bradesco_camara_compensacao:
      pos: [ 77,79 ]
      picture: '9(3)'
      default: 237
    nome_banco:
      pos: [ 80,94 ]
      picture: 'X(15)'
      default: 'BRADESCO'
    data_geracao:
      pos: [ 95,100 ]
      picture: '9(6)'
    densidade:
      pos: [ 101,108 ]
      picture: '9(8)'
    numero_aviso_bancario:
      pos: [ 109,113 ]
      picture: '9(5)'
    brancos_01:
      pos: [ 114,379 ]
      picture: 'X(266)'
      default: ''
    data_credito:
      pos: [ 380,385 ]
      picture: '9(6)'
    brancos_02:
      pos: [ 386,394 ]
      picture: 'X(9)'
      default: ''
    numero_sequencial_registro:
      pos: [ 395,400 ]
      picture: '9(6)'
      default: '000001'

  trailer_arquivo:
    identificacao_registro:
      pos: [ 1,1 ]
      picture: '9(1)'
      default: 9
    codigo_retorno:
      pos: [ 2,2 ]
      picture: '9(1)'
      default: 2
    identificacao_tipo_registro:
      pos: [ 3,4 ]
      picture: '9(2)'
      default: '01'
    codigo_banco:
      pos: [ 5,7 ]
      picture: '9(3)'
      default: 237
    brancos_01:
      pos: [ 8,17 ]
      picture: 'X(10)'
      default: ''
    quantidade_titulos_cobranca:
      pos: [ 18,25 ]
      picture: '9(8)'
    valor_total_cobranaca:
      pos: [ 26,39 ]
      picture: '9(12)V9(2)'
    aviso_bancario_01:
      pos: [ 40,47 ]
      picture: 'X(8)'
    brancos_02:
      pos: [ 48,57 ]
      picture: 'X(10)'
      default: ''
    quantidade_titulos_confirmacao_entrada:
      pos: [ 58,62 ]
      picture: '9(5)'
    valor_registros_confirmacao_entrada:
      pos: [ 63,74 ]
      picture: '9(10)V9(2)'
    valor_registros_liquidacao:
      pos: [ 75,86 ]
      picture: '9(10)V9(2)'
    quantidade_titulos_liquidacao:
      pos: [ 87,91 ]
      picture: '9(5)'
    valor_registros_liquidacao_06:
      pos: [ 92,103 ]
      picture: '9(10)V9(2)'
    quantidade_titulos_baixado:
      pos: [ 104,108 ]
      picture: '9(5)'
    valor_registros_baixado:
      pos: [ 109,120 ]
      picture: '9(10)V9(2)'
    quantidade_titulos_abatimento_cancelado:
      pos: [ 121,125 ]
      picture: '9(5)'
    valor_registros_abatimento_cancelado:
      pos: [ 126,137 ]
      picture: '9(10)V9(2)'
    quantidade_titulos_vencimento_alterado:
      pos: [ 138,142 ]
      picture: '9(5)'
    valor_registros_vencimento_alterado:
      pos: [ 143,154 ]
      picture: '9(10)V9(2)'
    quantidade_titulos_abatimento_concedido:
      pos: [ 125,159 ]
      picture: '9(5)'
    valor_registros_abatimento_concedido:
      pos: [ 160,171 ]
      picture: '9(10)V9(2)'
    quantidade_registros_instrucao_protesto:
      pos: [ 172,176 ]
      picture: '9(5)'
    valor_registros_instrucao_protesto:
      pos: [ 177,188 ]
      picture: '9(10)V9(2)'
    brancos_03:
      pos: [ 189,362 ]
      picture: 'X(174)'
      default: ''
    valor_total_rateios:
      pos: [ 363,377 ]
      picture: '9(13)V9(2)'
    quantidade_total_rateios:
      pos: [ 378, 394 ]
      picture: '9(8)'
    brancos_04:
      pos: [ 386,394 ]
      picture: 'X(9)'
      default: ''
    numero_sequencial_registro:
      pos: [ 395,400 ]
      picture: '9(6)'

  detalhes:
    segmento_1:
      identificacao_registro:
        pos: [ 1,1 ]
        picture: '9(1)'
        default: 1
      codigo_inscricao:
        pos: [ 2,3 ]
        picture: '9(2)'
      numero_inscricao:
        pos: [ 4,17 ]
        picture: '9(14)'
      zeros_01:
        pos: [ 18,20 ]
        picture: '9(3)'
        default: 0
      identificacao_empresa_beneficiaria:
        pos: [ 21,37 ]
        picture: 'X(17)'
      cod_participante_empresa:
        pos: [ 38,62 ]
        picture: 'X(25)'
      zeros_02:
        pos: [ 63,70 ]
        picture: '9(8)'
        default: 0
      identificacao_titulo_nosso_numero:
        pos: [ 71,82 ]
        picture: 'X(12)'
      uso_banco_01:
        pos: [ 83,92 ]
        picture: '9(10)'
      uso_banco_02:
        pos: [ 93,104 ]
        picture: '9(12)'
      identificador_rateio:
        pos: [ 105,105 ]
        picture: 'X(1)'
      pagamento_parcial:
        pos: [ 106,107 ]
        picture: '9(2)'
      codigo_carteira:
        pos: [ 108,108 ]
        picture: '9(1)'
      codigo_ocorrencia:
        pos: [ 109,110 ]
        picture: '9(2)'
      data_ocorrencia:
        pos: [ 111,116 ]
        picture: '9(6)'
      numero_documento:
        pos: [ 117,126 ]
        picture: 'X(10)'
      identificacao_banco_nosos_mumero:
        pos: [ 127,146 ]
        picture: 'X(20)'
      vencimento:
        pos: [ 147,152 ]
        picture: '9(6)'
      valor_titulo:
        pos: [ 153,165 ]
        picture: '9(11)V9(2)'
      codigo_banco:
        pos: [ 166,168 ]
        picture: '9(3)'
      agencia_cobradora:
        pos: [ 169,173 ]
        picture: '9(5)'
      especie:
        pos: [ 174,175 ]
        picture: 'X(2)'
      tarifa_cobranca:
        pos: [ 176,188 ]
        picture: '9(11)V9(2)'
      tarifa_outras_despesas:
        pos: [ 189,201 ]
        picture: '9(11)V9(2)'
      juros_operacao_atraso:
        pos: [ 202,214 ]
        picture: '9(11)V9(2)'
      valor_iof:
        pos: [ 215,227 ]
        picture: '9(11)V9(2)'
      valor_abatimento:
        pos: [ 228,240 ]
        picture: '9(11)V9(2)'
      descontos:
        pos: [ 241,253 ]
        picture: '9(11)V9(2)'
      valor_principal:
        pos: [ 254,266 ]
        picture: '9(11)V9(2)'
      juros_mora_multa:
        pos: [ 267,279 ]
        picture: '9(11)V9(2)'
      outros_creditos:
        pos: [ 280,292 ]
        picture: '9(11)V9(2)'
      brancos_06:
        pos: [ 293,294 ]
        picture: 'X(2)'
        default: ''
      motivo_ocorrencia_25:
        pos: [ 295,295 ]
        picture: 'X(1)'
        default: ''
      data_credito:
        pos: [ 296,301 ]
        picture: 'X(6)'
      origem_pagamento:
        pos: [ 302,304 ]
        picture: '9(3)'
      brancos_07:
        pos: [ 305,314 ]
        picture: 'X(10)'
        default: ''
      cheque_bradesco:
        pos: [ 315,318 ]
        picture: '9(4)'
      motivo_rejeicao_ocorrencia:
        pos: [ 319,328 ]
        picture: '9(10)'
      brancos_08:
        pos: [ 329,368 ]
        picture: 'X(40)'
        default: ''
      numero_cartorio:
        pos: [ 369,370 ]
        picture: '9(2)'
      numero_protocolo:
        pos: [ 371,380 ]
        picture: 'X(10)'
      brancos_09:
        pos: [ 381,394 ]
        picture: 'X(14)'
        default: ''
      numero_sequencial_registro:
        pos: [ 395,400 ]
        picture: '9(6)'
`
