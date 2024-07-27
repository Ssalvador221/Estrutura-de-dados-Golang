package Student

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"MichelPilha/pkg/errors"

	"github.com/google/uuid"
)

type Components interface {
	ToStoreStudent() (err error)
	GetAllStudentsFromList() (err error)
	GetStudentById() (m StudentsList, err error)
	RemoveAllStudents() (l StudentsList)
	RemoveStudentByMatricula() (rs StudentsList, err error)
}

type Student struct {
	ID           string
	NomeCompleto string
	Idade        int
	Matricula    int
}

type StudentsList struct {
	addStudentList []Student
}

type implStudent struct {
	Student
}

func NewStudent() Components {
	return &implStudent{}
}

var sl StudentsList

func (e implStudent) ToStoreStudent() (err error) {

	s := Student{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Adicione o nome do aluno: ")
	s.NomeCompleto, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(errors.ErrNullName)
		return
	}

	firstName := strings.Fields(s.NomeCompleto)
	fmt.Printf("Informe-nos a sua idade %v: ", firstName[0])
	_, err = fmt.Scan(&s.Idade)
	if err != nil {
		fmt.Println(errors.ErrNullAge)
		return
	}

	fmt.Printf("Informe-nos a sua matricula %v: ", firstName[0])
	_, err = fmt.Scan(&s.Matricula)
	if err != nil {
		fmt.Println(errors.ErrNullMatricula)
		return
	}

	s.ID = uuid.New().String()

	sl.addStudentList = append(sl.addStudentList, s)

	fmt.Println("Dados do aluno salvo com sucesso!")

	return nil
}

func (e implStudent) GetAllStudentsFromList() (err error) {

	for _, s := range sl.addStudentList {
		fmt.Printf(
			"Id do Aluno: %s"+
				"\nNome do Aluno: %s"+
				"Idade do aluno: %d\n"+
				"Matricula do aluno: %d\n",
			s.ID, s.NomeCompleto, s.Idade, s.Matricula)
	}

	return nil
}

func (e implStudent) GetStudentById() (m StudentsList, err error) {

	var idMatricula int
	var body StudentsList

	fmt.Println("\nInforme-nos a matricula do aluno que deseja procurar: ")
	_, err = fmt.Scan(&idMatricula)
	if err != nil {
		return body, errors.ErrNullMatricula
	}

	matriculaEncontrada := false
	for _, m := range sl.addStudentList {
		if idMatricula == m.Matricula {
			fmt.Printf("Aluno localizado!\nId do Aluno: %s\nNome do Aluno: %sIdade do aluno: %d\nMatricula do aluno: %d", m.ID, m.NomeCompleto, m.Idade, m.Matricula)
			matriculaEncontrada = true
			break
		}
	}

	if !matriculaEncontrada {
		return body, errors.ErrNotFoundMatricula
	}

	return m, nil
}

func (e implStudent) RemoveAllStudents() (l StudentsList) {

	sl.addStudentList = l.addStudentList

	l = StudentsList{}

	fmt.Println("Todos os seus alunos foram removidos", l)

	return l
}

func (e implStudent) RemoveStudentByMatricula() (rs StudentsList, err error) {

	var idMatricula int
	var body StudentsList

	fmt.Println("\nInforme-nos a matricula do aluno que deseja procurar: ")
	_, err = fmt.Scan(&idMatricula)
	if err != nil {
		return body, errors.ErrNullMatricula
	}

	matriculaEncontrada := false
	for _, m := range sl.addStudentList {
		if idMatricula == m.Matricula {
			rs = body
			fmt.Printf("Aluno localizado e removido com sucesso!\nNome do Aluno: %sIdade do aluno: %d\nMatricula do aluno: %d", m.NomeCompleto, m.Idade, m.Matricula)
			matriculaEncontrada = true
		}
	}

	if !matriculaEncontrada {
		return body, errors.ErrNotFoundMatricula
	}

	return rs, nil
}
