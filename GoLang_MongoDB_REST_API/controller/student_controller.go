package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sraynitjsr/model"
	"github.com/sraynitjsr/service"
)

func RegisterStudentRoutes(router *mux.Router) {
	router.HandleFunc("/students", addStudent).Methods("POST")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id}", findStudent).Methods("GET")
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/name/{name}", findStudentsByName).Methods("GET")
	router.HandleFunc("/students/roll/{roll}", findStudentByRoll).Methods("GET")
	router.HandleFunc("/students/sort/age", sortStudentsByAge).Methods("GET")
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	err := service.AddStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := service.DeleteStudent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func findStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	student, err := service.FindStudent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	students := service.GetAllStudents()
	json.NewEncoder(w).Encode(students)
}

func findStudentsByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	students := service.FindStudentsByName(name)
	json.NewEncoder(w).Encode(students)
}

func findStudentByRoll(w http.ResponseWriter, r *http.Request) {
	roll := mux.Vars(r)["roll"]
	student, err := service.FindStudentByRoll(roll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func sortStudentsByAge(w http.ResponseWriter, r *http.Request) {
	students := service.SortStudentsByAge()
	json.NewEncoder(w).Encode(students)
}
