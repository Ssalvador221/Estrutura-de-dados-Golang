package main

import (
	"MichelPilha/Student"
)

func main() {
	student := Student.StudentsList{}

	student.ToStoreStudent()
	student.RemoveAllStudents()
	student.RemoveStudentByMatricula()
	student.GetAllStudentsFromList()
	student.GetStudentById()

	//fazer menu para chamar todas as ações!

}
