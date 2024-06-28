package errors

import "fmt"


var  (
	ErrNullName          = fmt.Errorf("por favor insira um nome valido")
	ErrNullAge           = fmt.Errorf("por favor insira uma idade valida")
	ErrNullMatricula     = fmt.Errorf("por favor insira uma matricula valida")
	ErrNotFoundMatricula = fmt.Errorf("Matricula do aluno não encontrado")
	ErrNullValues        = fmt.Errorf("não pudemos salvar os seus dados")
)

