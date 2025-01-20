package service

import (
	"errors"
	"sort"

	"github.com/sraynitjsr/model"
	"github.com/sraynitjsr/repository"
)

func AddStudent(student model.Student) error {
	if repository.RollExists(student.Roll) {
		return errors.New("student with the same roll already exists")
	}
	return repository.AddStudent(student)
}

func DeleteStudent(id string) error {
	return repository.DeleteStudent(id)
}

func FindStudent(id string) (model.Student, error) {
	return repository.FindStudent(id)
}

func GetAllStudents() []model.Student {
	return repository.GetAllStudents()
}

func FindStudentsByName(name string) []model.Student {
	return repository.FindStudentsByName(name)
}

func FindStudentByRoll(roll string) (model.Student, error) {
	return repository.FindStudentByRoll(roll)
}

func SortStudentsByAge() []model.Student {
	students := repository.GetAllStudents()
	sort.Slice(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	return students
}
