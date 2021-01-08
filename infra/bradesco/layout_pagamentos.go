package bradesco



const CNAB240PagamentosBRADESCO = `

# FORMATO: ITAU - CNAB240
# OBJETIVO DO ARQUIVO: PAGAMENTOS Cheque, OP, DOC, TED ou crédito em conta - SISPAG ITAÚ
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 240 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita.
# Pagamento através de Cheque, OP, DOC, TED, ou crédito em conta corrente:
#
# https://banco.bradesco/assets/pessoajuridica/pdf/jun-19-layout-multipag.pdf
servico: pagamento_cheque_op_doc_ted_credito_cc
versao: "04"
layout: "cnab240"

remessa:
  header_arquivo:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
      default: 0000
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: "0"
    brancos_01:
      pos: [ 9, 17 ]
      picture: "X(9)"
      default: ""
    tipo_inscricao:
      pos: [ 18, 18 ]
      picture: "9(1)"
    inscricao_numero:
      pos: [ 19, 32 ]
      picture: "9(14)"
    codigo_convenio_banco:
      pos: [ 33, 52 ]
      picture: "X(20)"
      default: ""
    agencia_debito:
      pos: [ 53, 57 ]
      picture: "9(5)"
    agencia_digito:
      pos: [ 58, 58 ]
      picture: "X(1)"
    conta_debito:
      pos: [ 59, 70 ]
      picture: "9(12)"
    conta_digito:
      pos: [ 71, 71 ]
      picture: "X(1)"
    digito_ver_conta_agencia:
      pos: [ 72, 72 ]
      picture: "X(1)"
    nome_empresa:
      pos: [ 73, 102 ]
      picture: "X(30)"
    nome_banco:
      pos: [ 103, 132 ]
      picture: "X(30)"
    brancos_05:
      pos: [ 133, 142 ]
      picture: "X(10)"
      default: ""
    arquivo_codigo:
      pos: [ 143, 143 ]
      picture: "9(1)"
      default: 1
    data_geracao:
      pos: [ 144, 151 ]
      picture: "9(8)"
    hora_geracao:
      pos: [ 152, 157 ]
      picture: "9(6)"
    numero_sequencial_arquivo:
      pos: [ 158, 163 ]
      picture: "9(6)"
    numero_versao_layout:
      pos: [ 164, 166 ]
      picture: "9(3)"
      default: 089
    densidade_gravacao:
      pos: [ 167, 171 ]
      picture: "9(5)"
      default: 0
    brancos_06:
      pos: [ 172, 240 ]
      picture: "X(69)"
      default: ""

  header_lote:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: 1
    tipo_operacao:
      pos: [ 9, 9 ]
      picture: "X(1)"
      default: "C"
    tipo_servico:
      pos: [ 10, 11 ]
      picture: "9(2)"
    forma_lancamento:
      pos: [ 12, 13 ]
      picture: "9(2)"
    layout_lote:
      pos: [ 14, 16 ]
      picture: "9(3)"
      default: 045
    brancos_01:
      pos: [ 17, 17 ]
      picture: "X(1)"
      default: ""
    tipo_inscricao_empresa:
      pos: [ 18, 18 ]
      picture: "9(1)"
    numero_inscricao:
      pos: [ 19, 32 ]
      picture: "9(14)"
    codigo_convenio_banco:
      pos: [ 33, 52 ]
      picture: "X(20)"
    agencia_debito:
      pos: [ 53, 57 ]
      picture: "9(5)"
    agencia_debito_digito:
      pos: [ 58, 58 ]
      picture: "X(1)"
    conta_debito:
      pos: [ 59, 70 ]
      picture: "9(12)"
    conta_debito_digito:
      pos: [ 71, 71 ]
      picture: "X(1)"
    conta_agencia_debito_digito:
      pos: [ 72, 72 ]
      picture: "X(1)"
    nome_empresa:
      pos: [ 73, 102 ]
      picture: "X(30)"
    mensagem_1:
      pos: [ 103, 142 ]
      picture: "X(40)"
    endereco_empresa:
      pos: [ 143, 172 ]
      picture: "X(30)"
    numero:
      pos: [ 173, 177 ]
      picture: "9(5)"
    complemento:
      pos: [ 178, 192 ]
      picture: "X(15)"
      default: "CASA"
    cidade:
      pos: [ 193, 212 ]
      picture: "X(20)"
    cep:
      pos: [ 213, 217 ]
      picture: "9(5)"
    cep_complemento:
      pos: [ 218, 220 ]
      picture: "9(3)"
    estado:
      pos: [ 221, 222 ]
      picture: "X(2)"
    indicativo_forma_pagamento:
      pos: [ 223, 224 ]
      picture: "9(2)"
      default: '01'
    brancos_02:
      pos: [ 225, 230 ]
      picture: "X(6)"
      default: ""
    codigo_ocorrencias:
      pos: [ 231, 240 ]
      picture: "X(10)"

  trailer_lote:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: 5
    brancos_01:
      pos: [ 9, 17 ]
      picture: "X(9)"
      default: ""
    total_registros_lote:
      pos: [ 18, 23 ]
      picture: "9(6)"
    total_valor_pagtos:
      pos: [ 24, 41 ]
      picture: "9(16)V9(2)"
    total_moedas:
      pos: [ 42, 59 ]
      picture: "9(13)V9(5)"
    numero_aviso_debito:
      pos: [ 60, 65 ]
      picture: "9(6)"
    brancos_02:
      pos: [ 66, 230 ]
      picture: "X(165)"
      default: ""
    codigos_ocorrencias:
      pos: [ 231, 240 ]
      picture: "X(10)"

  trailer_arquivo:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
      default: 9999
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: 9
    brancos_01:
      pos: [ 9, 17 ]
      picture: "X(9)"
      default: ""
    total_lotes_arquivo:
      pos: [ 18, 23 ]
      picture: "9(6)"
    total_registros:
      pos: [ 24, 29 ]
      picture: "9(6)"
    total_contas_cobranca:
      pos: [ 30, 35 ]
      picture: "9(6)"
    brancos_02:
      pos: [ 36, 240 ]
      picture: "X(205)"
      default: ""

  detalhes:
    # segmentos
    # - Segmento A - (Obrigatório)
    # - Segmento Z – (Opcional – Informado somente no arquivo retorno, quando contratado junto ao Banco a opção de Autenticação Eletrônica de Pagamentos por arquivo)
    # -- Segmento J (OBRIGATORIO)
    segmento_a:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "A"
      tipo_movimento:
        pos: [ 15, 15 ]
        picture: "9(1)"
      movimento_codigo:
        pos: [ 16, 17 ]
        picture: "9(2)"
      codigo_camara_centralizadora:
        pos: [ 18, 20 ]
        picture: "9(3)"
      codigo_banco_favorecido:
        pos: [ 21, 23 ]
        picture: "9(3)"
      agencia_favorecido:
        pos: [ 24, 28 ]
        picture: "9(5)"
      agencia_digito_favorecido:
        pos: [ 29, 29 ]
        picture: "X(1)"
      conta_favorecido:
        pos: [ 30, 41 ]
        picture: "9(12)"
      conta_digito_favorecido:
        pos: [ 42, 42 ]
        picture: "X(1)"
      digito_conta_agencia_favorecido:
        pos: [ 43, 43 ]
        picture: "X(1)"
      nome_favorecido:
        pos: [ 44, 73 ]
        picture: "X(30)"
      numero_doc:
        pos: [ 74, 93 ]
        picture: "X(20)"
      data_pagto:
        pos: [ 94, 101 ]
        picture: "9(8)"
      tipo_moeda:
        pos: [ 102, 104 ]
        picture: "X(3)"
        default: "BRL"
      quantidade_moeda:
        pos: [ 105, 119 ]
        picture: "9(15)"
      valor_pagto:
        pos: [ 120, 134 ]
        picture: "9(13)V9(2)"
      nosso_numero:
        pos: [ 135, 154 ]
        picture: "X(20)"
      data_efetiva:
        pos: [ 155, 162 ]
        picture: "9(8)"
      valor_efetivo:
        pos: [ 163, 177 ]
        picture: "9(13)V9(2)"
      info_deposito_judicial_pg_siape:
        pos: [ 178, 217 ]
        picture: "X(40)"
      codigo_tipo_servico:
        pos: [ 218, 219 ]
        picture: "X(2)"
      finalidade_ted:
        pos: [ 220, 224 ]
        picture: "X(5)"
      compl_finalidade_ted:
        pos: [ 225, 226 ]
        picture: "X(2)"
      brancos_03:
        pos: [ 227, 229 ]
        picture: "X(3)"
        default: ""
      aviso:
        pos: [ 230, 230 ]
        picture: "X(1)"
      codigo_ocorrencias:
        pos: [ 231, 240 ]
        picture: "X(10)"

    segmento_b:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "B"
      brancos_01:
        pos: [ 15, 17 ]
        picture: "X(3)"
        default: ""
      codigo_inscricao_favorecido:
        pos: [ 18, 18 ]
        picture: "9(1)"
      numero_inscricao_favorecido:
        pos: [ 19, 32 ]
        picture: "9(14)"
      logadouro:
        pos: [ 33, 62 ]
        picture: "X(30)"
      numero:
        pos: [ 63, 67 ]
        picture: "9(5)"
      complemento:
        pos: [ 68, 82 ]
        picture: "X(15)"
      bairro:
        pos: [ 83, 97 ]
        picture: "X(15)"
      cidade:
        pos: [ 98, 117 ]
        picture: "X(20)"
      cep:
        pos: [ 118, 122 ]
        picture: "9(5)"
      cep_complemento:
        pos: [ 123, 125 ]
        picture: "9(3)"
      estado:
        pos: [ 126, 127 ]
        picture: "X(2)"
      data_vencimento:
        pos: [ 128, 135 ]
        picture: "9(8)"
      valor_documento:
        pos: [ 136, 150 ]
        picture: "9(13)V9(2)"
      valor_abatimento:
        pos: [ 151, 165 ]
        picture: "9(13)V9(2)"
      valor_desconto:
        pos: [ 166, 180 ]
        picture: "9(13)V9(2)"
      valor_mora:
        pos: [ 181, 195 ]
        picture: "9(13)V9(2)"
      valor_multa:
        pos: [ 196, 210 ]
        picture: "9(13)V9(2)"
      cod_documento_favorecido:
        pos: [ 211, 225 ]
        picture: "X(15)"
      aviso_favorecido:
        pos: [ 226, 226 ]
        picture: "9(1)"
      uso_siape:
        pos: [ 227, 232 ]
        picture: "9(6)"
      codigo_ispb:
        pos: [ 233, 240 ]
        picture: "9(8)"

    segmento_z:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      codigo_segmento:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "Z"
      autenticacao_legislacao:
        pos: [ 15, 78 ]
        picture: "X(64)"
      autenticacao_protocolo:
        pos: [ 79, 103 ]
        picture: "X(25)"
      brancos_01:
        pos: [ 104, 230 ]
        picture: "X(127)"
        default: ""
      codigo_ocorrencias_retorno:
        pos: [ 231, 240 ]
        picture: "X(10)"
        default: ""

    segmento_j:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "J"
      tipo_movimento:
        pos: [ 15, 15 ]
        picture: "9(1)"
      movimento_codigo:
        pos: [ 16, 17 ]
        picture: "9(2)"
      codigo_barra:
        pos: [ 18, 61 ]
        picture: "X(44)"
      nome_cedente:
        pos: [ 62, 91 ]
        picture: "X(30)"
      data_vencimento:
        pos: [ 92, 99 ]
        picture: "9(8)"
      valor_titulo:
        pos: [ 100, 114 ]
        picture: "9(13)V9(2)"
      valor_desconto:
        pos: [ 115, 129 ]
        picture: "9(13)V9(2)"
      valor_mora:
        pos: [ 130, 144 ]
        picture: "9(13)V9(2)"
      data_pagamento:
        pos: [ 145, 152 ]
        picture: "9(8)"
      valor_pagamento:
        pos: [ 153, 167 ]
        picture: "9(13)V9(2)"
      quantidade_moeda:
        pos: [ 168, 182 ]
        picture: "9(10)V9(5)"
      codigo_documento_empresa:
        pos: [ 183, 202 ]
        picture: "X(20)"
      codigo_documento_banco:
        pos: [ 203, 222 ]
        picture: "X(20)"
      codigo_moeda:
        pos: [ 223, 224 ]
        picture: "9(2)"
      brancos_01:
        pos: [ 225, 230 ]
        picture: "X(6)"
        default: ""
      codigo_ocorrencias:
        pos: [ 231, 240 ]
        picture: "X(10)"

    segmento_j_52:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "J"
      brancos_01:
        pos: [ 15, 15 ]
        picture: "X(1)"
        default: ""
      movimento_codigo:
        pos: [ 16, 17 ]
        picture: "9(2)"
      codigo_registro_opcional:
        pos: [ 18, 19 ]
        picture: "9(2)"
        default: 52
      sacado_tipo_inscricao:
        pos: [ 20 ,20 ]
        picture: "9(1)"
      sacado_numero_inscricao:
        pos: [ 21 ,35 ]
        picture: "9(15)"
      sacado_nome:
        pos: [ 36 ,75 ]
        picture: "X(40)"
      cedente_tipo_inscricao:
        pos: [ 76 ,76 ]
        picture: "9(1)"
      cedente_numero_inscricao:
        pos: [ 77 ,91 ]
        picture: "9(15)"
      cedente_nome:
        pos: [ 92 ,131 ]
        picture: "X(40)"
      sacador_tipo_inscricao:
        pos: [ 132 ,132 ]
        picture: "9(1)"
      sacador_numero_inscricao:
        pos: [ 133 , 147 ]
        picture: "9(15)"
      sacador_nome:
        pos: [ 148 ,187 ]
        picture: "X(40)"
      brancos_02:
        pos: [ 188, 240 ]
        picture: "X(53)"
        default: ""

retorno:
  header_arquivo:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
      default: 0000
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: "0"
    brancos_01:
      pos: [ 9, 17 ]
      picture: "X(9)"
      default: ""
    tipo_inscricao:
      pos: [ 18, 18 ]
      picture: "9(1)"
    inscricao_numero:
      pos: [ 19, 32 ]
      picture: "9(14)"
    codigo_convenio_banco:
      pos: [ 33, 52 ]
      picture: "X(20)"
      default: ""
    agencia_debito:
      pos: [ 53, 57 ]
      picture: "9(5)"
    agencia_digito:
      pos: [ 58, 58 ]
      picture: "X(1)"
    conta_debito:
      pos: [ 59, 70 ]
      picture: "9(12)"
    conta_digito:
      pos: [ 71, 71 ]
      picture: "X(1)"
    digito_ver_conta_agencia:
      pos: [ 72, 72 ]
      picture: "X(1)"
    nome_empresa:
      pos: [ 73, 102 ]
      picture: "X(30)"
    nome_banco:
      pos: [ 103, 132 ]
      picture: "X(30)"
    brancos_05:
      pos: [ 133, 142 ]
      picture: "X(10)"
      default: ""
    arquivo_codigo:
      pos: [ 143, 143 ]
      picture: "9(1)"
      default: 1
    data_geracao:
      pos: [ 144, 151 ]
      picture: "9(8)"
    hora_geracao:
      pos: [ 152, 157 ]
      picture: "9(6)"
    numero_sequencial_arquivo:
      pos: [ 158, 163 ]
      picture: "9(6)"
    numero_versao_layout:
      pos: [ 164, 166 ]
      picture: "9(3)"
      default: 089
    densidade_gravacao:
      pos: [ 167, 171 ]
      picture: "9(5)"
      default: 0
    brancos_06:
      pos: [ 172, 240 ]
      picture: "X(69)"
      default: ""

  header_lote:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: 1
    tipo_operacao:
      pos: [ 9, 9 ]
      picture: "X(1)"
      default: "C"
    tipo_servico:
      pos: [ 10, 11 ]
      picture: "9(2)"
    forma_lancamento:
      pos: [ 12, 13 ]
      picture: "9(2)"
    layout_lote:
      pos: [ 14, 16 ]
      picture: "9(3)"
      default: 045
    brancos_01:
      pos: [ 17, 17 ]
      picture: "X(1)"
      default: ""
    tipo_inscricao_empresa:
      pos: [ 18, 18 ]
      picture: "9(1)"
    numero_inscricao:
      pos: [ 19, 32 ]
      picture: "9(14)"
    codigo_convenio_banco:
      pos: [ 33, 52 ]
      picture: "X(20)"
    agencia_debito:
      pos: [ 53, 57 ]
      picture: "9(5)"
    agencia_debito_digito:
      pos: [ 58, 58 ]
      picture: "X(1)"
    conta_debito:
      pos: [ 59, 70 ]
      picture: "9(12)"
    conta_debito_digito:
      pos: [ 71, 71 ]
      picture: "X(1)"
    conta_agencia_debito_digito:
      pos: [ 72, 72 ]
      picture: "X(1)"
    nome_empresa:
      pos: [ 73, 102 ]
      picture: "X(30)"
    mensagem_1:
      pos: [ 103, 142 ]
      picture: "X(40)"
    endereco_empresa:
      pos: [ 143, 172 ]
      picture: "X(30)"
    numero:
      pos: [ 173, 177 ]
      picture: "9(5)"
    complemento:
      pos: [ 178, 192 ]
      picture: "X(15)"
      default: "CASA"
    cidade:
      pos: [ 193, 212 ]
      picture: "X(20)"
    cep:
      pos: [ 213, 217 ]
      picture: "9(5)"
    cep_complemento:
      pos: [ 218, 220 ]
      picture: "9(3)"
    estado:
      pos: [ 221, 222 ]
      picture: "X(2)"
    indicativo_forma_pagamento:
      pos: [ 223, 224 ]
      picture: "9(2)"
      default: '01'
    brancos_02:
      pos: [ 225, 230 ]
      picture: "X(6)"
      default: ""
    codigo_ocorrencias:
      pos: [ 231, 240 ]
      picture: "X(10)"

  trailer_lote:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: 5
    brancos_01:
      pos: [ 9, 17 ]
      picture: "X(9)"
      default: ""
    total_registros_lote:
      pos: [ 18, 23 ]
      picture: "9(6)"
    total_valor_pagtos:
      pos: [ 24, 41 ]
      picture: "9(16)V9(2)"
    total_moedas:
      pos: [ 42, 59 ]
      picture: "9(13)V9(5)"
    numero_aviso_debito:
      pos: [ 60, 65 ]
      picture: "9(6)"
    brancos_02:
      pos: [ 66, 230 ]
      picture: "X(165)"
      default: ""
    codigos_ocorrencias:
      pos: [ 231, 240 ]
      picture: "X(10)"

  trailer_arquivo:
    codigo_banco:
      pos: [ 1, 3 ]
      picture: "9(3)"
      default: 237
    codigo_lote:
      pos: [ 4, 7 ]
      picture: "9(4)"
      default: 9999
    tipo_registro:
      pos: [ 8, 8 ]
      picture: "9(1)"
      default: 9
    brancos_01:
      pos: [ 9, 17 ]
      picture: "X(9)"
      default: ""
    total_lotes_arquivo:
      pos: [ 18, 23 ]
      picture: "9(6)"
    total_registros:
      pos: [ 24, 29 ]
      picture: "9(6)"
    total_contas_cobranca:
      pos: [ 30, 35 ]
      picture: "9(6)"
    brancos_02:
      pos: [ 36, 240 ]
      picture: "X(205)"
      default: ""

  detalhes:
    # segmentos
    # - Segmento A - (Obrigatório)
    # - Segmento Z – (Opcional – Informado somente no arquivo retorno, quando contratado junto ao Banco a opção de Autenticação Eletrônica de Pagamentos por arquivo)
    # -- Segmento J (OBRIGATORIO)
    segmento_a:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "A"
      tipo_movimento:
        pos: [ 15, 15 ]
        picture: "9(1)"
      movimento_codigo:
        pos: [ 16, 17 ]
        picture: "9(2)"
      codigo_camara_centralizadora:
        pos: [ 18, 20 ]
        picture: "9(3)"
      codigo_banco_favorecido:
        pos: [ 21, 23 ]
        picture: "9(3)"
      agencia_favorecido:
        pos: [ 24, 28 ]
        picture: "9(5)"
      agencia_digito_favorecido:
        pos: [ 29, 29 ]
        picture: "X(1)"
      conta_favorecido:
        pos: [ 30, 41 ]
        picture: "9(12)"
      conta_digito_favorecido:
        pos: [ 42, 42 ]
        picture: "X(1)"
      digito_conta_agencia_favorecido:
        pos: [ 43, 43 ]
        picture: "X(1)"
      nome_favorecido:
        pos: [ 44, 73 ]
        picture: "X(30)"
      numero_doc:
        pos: [ 74, 93 ]
        picture: "X(20)"
      data_pagto:
        pos: [ 94, 101 ]
        picture: "9(8)"
      tipo_moeda:
        pos: [ 102, 104 ]
        picture: "X(3)"
        default: "BRL"
      quantidade_moeda:
        pos: [ 105, 119 ]
        picture: "9(15)"
      valor_pagto:
        pos: [ 120, 134 ]
        picture: "9(13)V9(2)"
      nosso_numero:
        pos: [ 135, 154 ]
        picture: "X(20)"
      data_efetiva:
        pos: [ 155, 162 ]
        picture: "9(8)"
      valor_efetivo:
        pos: [ 163, 177 ]
        picture: "9(13)V9(2)"
      info_deposito_judicial_pg_siape:
        pos: [ 178, 217 ]
        picture: "X(40)"
      codigo_tipo_servico:
        pos: [ 218, 219 ]
        picture: "X(2)"
      finalidade_ted:
        pos: [ 220, 224 ]
        picture: "X(5)"
      compl_finalidade_ted:
        pos: [ 225, 226 ]
        picture: "X(2)"
      brancos_03:
        pos: [ 227, 229 ]
        picture: "X(3)"
        default: ""
      aviso:
        pos: [ 230, 230 ]
        picture: "X(1)"
      codigo_ocorrencias:
        pos: [ 231, 240 ]
        picture: "X(10)"

    segmento_b:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "B"
      brancos_01:
        pos: [ 15, 17 ]
        picture: "X(3)"
        default: ""
      codigo_inscricao_favorecido:
        pos: [ 18, 18 ]
        picture: "9(1)"
      numero_inscricao_favorecido:
        pos: [ 19, 32 ]
        picture: "9(14)"
      logadouro:
        pos: [ 33, 62 ]
        picture: "X(30)"
      numero:
        pos: [ 63, 67 ]
        picture: "9(5)"
      complemento:
        pos: [ 68, 82 ]
        picture: "X(15)"
      bairro:
        pos: [ 83, 97 ]
        picture: "X(15)"
      cidade:
        pos: [ 98, 117 ]
        picture: "X(20)"
      cep:
        pos: [ 118, 122 ]
        picture: "9(5)"
      cep_complemento:
        pos: [ 123, 125 ]
        picture: "9(3)"
      estado:
        pos: [ 126, 127 ]
        picture: "X(2)"
      data_vencimento:
        pos: [ 128, 135 ]
        picture: "9(8)"
      valor_documento:
        pos: [ 136, 150 ]
        picture: "9(13)V9(2)"
      valor_abatimento:
        pos: [ 151, 165 ]
        picture: "9(13)V9(2)"
      valor_desconto:
        pos: [ 166, 180 ]
        picture: "9(13)V9(2)"
      valor_mora:
        pos: [ 181, 195 ]
        picture: "9(13)V9(2)"
      valor_multa:
        pos: [ 196, 210 ]
        picture: "9(13)V9(2)"
      cod_documento_favorecido:
        pos: [ 211, 225 ]
        picture: "X(15)"
      aviso_favorecido:
        pos: [ 226, 226 ]
        picture: "9(1)"
      uso_siape:
        pos: [ 227, 232 ]
        picture: "9(6)"
      codigo_ispb:
        pos: [ 233, 240 ]
        picture: "9(8)"

    segmento_j:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      segmento_codigo:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "J"
      tipo_movimento:
        pos: [ 15, 15 ]
        picture: "9(1)"
      movimento_codigo:
        pos: [ 16, 17 ]
        picture: "9(2)"
      codigo_barra:
        pos: [ 18, 61 ]
        picture: "X(44)"
      nome_cedente:
        pos: [ 62, 91 ]
        picture: "X(30)"
      data_vencimento:
        pos: [ 92, 99 ]
        picture: "9(8)"
      valor_titulo:
        pos: [ 100, 114 ]
        picture: "9(13)V9(2)"
      valor_desconto:
        pos: [ 115, 129 ]
        picture: "9(13)V9(2)"
      valor_mora:
        pos: [ 130, 144 ]
        picture: "9(13)V9(2)"
      data_pagamento:
        pos: [ 145, 152 ]
        picture: "9(8)"
      valor_pagamento:
        pos: [ 153, 167 ]
        picture: "9(13)V9(2)"
      quantidade_moeda:
        pos: [ 168, 182 ]
        picture: "9(10)V9(5)"
      codigo_documento_empresa:
        pos: [ 183, 202 ]
        picture: "X(20)"
      codigo_documento_banco:
        pos: [ 203, 222 ]
        picture: "X(20)"
      codigo_moeda:
        pos: [ 223, 224 ]
        picture: "9(2)"
      brancos_01:
        pos: [ 225, 230 ]
        picture: "X(6)"
        default: ""
      codigo_ocorrencias:
        pos: [ 231, 240 ]
        picture: "X(10)"

    segmento_z:
      codigo_banco:
        pos: [ 1, 3 ]
        picture: "9(3)"
        default: 237
      codigo_lote:
        pos: [ 4, 7 ]
        picture: "9(4)"
      tipo_registro:
        pos: [ 8, 8 ]
        picture: "9(1)"
        default: 3
      numero_registro:
        pos: [ 9, 13 ]
        picture: "9(5)"
      codigo_segmento:
        pos: [ 14, 14 ]
        picture: "X(1)"
        default: "Z"
      autenticacao_legislacao:
        pos: [ 15, 78 ]
        picture: "X(64)"
      autenticacao_protocolo:
        pos: [ 79, 103 ]
        picture: "X(25)"
      brancos_01:
        pos: [ 104, 230 ]
        picture: "X(127)"
        default: ""
      codigo_ocorrencias_retorno:
        pos: [ 231, 240 ]
        picture: "X(10)"
        default: ""

`