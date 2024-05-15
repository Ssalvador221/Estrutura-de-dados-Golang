package Student

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Funcs interface {
	ToStoreStudent()
	GetAllStudentsFromList()
	GetStudentById()
	RemoveAllStudents()
	RemoveStudentByMatricula()
}

type Student struct {
	nomeCompleto string
	age          int
	matricula    int
}

type StudentsList struct {
	addStudentList []Student
}

var (
	ErrNullName          = fmt.Errorf("por favor insira um nome valido")
	ErrNullAge           = fmt.Errorf("por favor insira uma idade valida")
	ErrNullMatricula     = fmt.Errorf("por favor insira uma matricula valida")
	ErrNotFoundMatricula = fmt.Errorf("Matricula do aluno não encontrado")
	ErrNullValues        = fmt.Errorf("não pudemos salvar os seus dados")
)

func (sl *StudentsList) ToStoreStudent() (err error) {

	s := Student{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Adicione o nome do aluno: ")
	s.nomeCompleto, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(ErrNullName)
		return
	}

	firstName := strings.Fields(s.nomeCompleto)
	fmt.Printf("Informe-nos a sua idade %v: ", firstName[0])
	_, err = fmt.Scan(&s.age)
	if err != nil {
		fmt.Println(ErrNullAge)
		return
	}

	fmt.Printf("Informe-nos a sua matricula %v: ", firstName[0])
	_, err = fmt.Scan(&s.matricula)
	if err != nil {
		fmt.Println(ErrNullMatricula)
		return
	}

	sl.addStudentList = append(sl.addStudentList, s)
	fmt.Println("Dados do aluno salvo com sucesso!")

	return err
}

func (sl *StudentsList) GetAllStudentsFromList() (s StudentsList) {

	for _, s := range sl.addStudentList {
		fmt.Printf(
			"\nNome do Aluno: %s"+
				"Idade do aluno: %d\n"+
				"Matricula do aluno: %d\n",
			s.nomeCompleto, s.age, s.matricula)
	}

	return s
}

func (sl *StudentsList) GetStudentById() (m StudentsList, err error) {

	var idMatricula int

	fmt.Println("\nInforme-nos a matricula do aluno que deseja procurar: ")
	_, err = fmt.Scan(&idMatricula)
	if err != nil {
		return StudentsList{}, ErrNullMatricula
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
		return StudentsList{}, ErrNotFoundMatricula
	}

	return m, nil
}

func (sl *StudentsList) RemoveAllStudents() (l StudentsList) {

	sl.addStudentList = l.addStudentList

	l = StudentsList{}

	fmt.Println("Todos os seus alunos foram removidos", l)

	return l
}

func (sl *StudentsList) RemoveStudentByMatricula() (rs StudentsList, err error) {

	var idMatricula int

	fmt.Println("\nInforme-nos a matricula do aluno que deseja procurar: ")
	_, err = fmt.Scan(&idMatricula)
	if err != nil {
		return StudentsList{}, ErrNullMatricula
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
		return StudentsList{}, ErrNotFoundMatricula
	}

	return rs, nil
}
