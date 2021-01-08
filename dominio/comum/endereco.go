package comum

import "fmt"

type Endereco struct {
      Rua string
      Complemento string
      Numero string
      Cep  string
      Bairro string
      Cidade string
      Estado string
}


func ( endereco Endereco )  EnderecoCompleto () string  {

   var strEndereco = fmt.Sprintf("%s,%s %s %s,%s-%s" , endereco.Rua ,endereco.Numero, endereco.Complemento,
   	                           endereco.Bairro , endereco.Cidade , endereco.Estado)
   fmt.Println(len(strEndereco))
	return strEndereco
}
