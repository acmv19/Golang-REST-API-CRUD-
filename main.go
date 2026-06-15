package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// task define la estructura de una tarea
type task struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

// allTask es un slice de tasks
type allTask []task

// tasks es la lista de tareas en memoria
var tasks = allTask{
	{
		ID:      "1",
		Name:    "Task One",
		Content: "some task",
	},
}

// indexRoute maneja la ruta raíz "/"
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to my api")
}

// getTasks devuelve todas las tareas en formato JSON
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// createTask crea una nueva tarea a partir del body de la request
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task

	// leer el body de la request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "insert a valid task")
		return // detener si hay error
	}

	// convertir el JSON del body a la estructura task
	json.Unmarshal(reqBody, &newTask)

	// asignar ID automáticamente basado en la cantidad de tareas
	newTask.ID = fmt.Sprintf("%d", len(tasks)+1)

	// agregar la nueva tarea a la lista
	tasks = append(tasks, newTask)

	// responder con la tarea creada en JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // status 201
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	// crear el router
	router := mux.NewRouter().StrictSlash(true)

	// registrar las rutas
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")   // obtener todas las tareas
	router.HandleFunc("/tasks", createTask).Methods("POST") // crear una tarea

	log.Println("server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}