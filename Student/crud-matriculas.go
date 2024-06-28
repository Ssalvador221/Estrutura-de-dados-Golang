package Student

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"MichelPilha/pkg/errors"
)

type Components interface {
	ToStoreStudent() (err error)
	GetAllStudentsFromList() (err error)
	GetStudentById() (m StudentsList, err error)
	RemoveAllStudents() (l StudentsList) 
	RemoveStudentByMatricula() (rs StudentsList, err error)
}

type Student struct {
	nomeCompleto string
	age          int
	matricula    int
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
	s.nomeCompleto, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(errors.ErrNullName)
		return
	}

	firstName := strings.Fields(s.nomeCompleto)
	fmt.Printf("Informe-nos a sua idade %v: ", firstName[0])
	_, err = fmt.Scan(&s.age)
	if err != nil {
		fmt.Println(errors.ErrNullAge)
		return
	}

	fmt.Printf("Informe-nos a sua matricula %v: ", firstName[0])
	_, err = fmt.Scan(&s.matricula)
	if err != nil {
		fmt.Println(errors.ErrNullMatricula)
		return
	}

	sl.addStudentList = append(sl.addStudentList, s)
	fmt.Println("Dados do aluno salvo com sucesso!")

	return err
}

func (e implStudent) GetAllStudentsFromList() (err error) {

	for _, s := range sl.addStudentList{
		fmt.Printf(
			"\nNome do Aluno: %s"+
				"Idade do aluno: %d\n"+
				"Matricula do aluno: %d\n",
			s.nomeCompleto, s.age, s.matricula)
	}

	return nil
}

func (e implStudent) GetStudentById() (m StudentsList, err error) {

	var idMatricula int

	fmt.Println("\nInforme-nos a matricula do aluno que deseja procurar: ")
	_, err = fmt.Scan(&idMatricula)
	if err != nil {
		return StudentsList{}, errors.ErrNullMatricula
	}

	matriculaEncontrada := false
	for _, m := range sl.addStudentList {
		if idMatricula == m.matricula {
			fmt.Printf("Aluno localizado!\nNome do Aluno: %sIdade do aluno: %d\nMatricula do aluno: %d", m.nomeCompleto, m.age, m.matricula)
			matriculaEncontrada = true
			break
		}
	}

	if !matriculaEncontrada {
		return StudentsList{}, errors.ErrNotFoundMatricula
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

	fmt.Println("\nInforme-nos a matricula do aluno que deseja procurar: ")
	_, err = fmt.Scan(&idMatricula)
	if err != nil {
		return StudentsList{}, errors.ErrNullMatricula
	}

	matriculaEncontrada := false
	for _, m := range sl.addStudentList {
		if idMatricula == m.matricula {
			rs = StudentsList{}
			fmt.Printf("Aluno localizado e removido com sucesso!\nNome do Aluno: %sIdade do aluno: %d\nMatricula do aluno: %d", m.nomeCompleto, m.age, m.matricula)
			matriculaEncontrada = true
		}
	}

	if !matriculaEncontrada {
		return StudentsList{}, errors.ErrNotFoundMatricula
	}

	return rs, nil
}
