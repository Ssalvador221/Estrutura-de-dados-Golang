package main

import (
	"MichelPilha/Student"
)

func main() {

	Student.NewStudent().ToStoreStudent()
	Student.NewStudent().ToStoreStudent()

	Student.NewStudent().GetAllStudentsFromList()
	Student.NewStudent().GetStudentById()

	Student.NewStudent().RemoveStudentByMatricula()
	Student.NewStudent().RemoveAllStudents()

	//fazer menu para chamar todas as ações!

}
