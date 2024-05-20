package routes

import (
	"log"
	"myapp/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // default port if not specified
	}

	router := mux.NewRouter()

	// Student routes (boys)
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{enrollment_no}", controller.GetStudent).Methods("GET")
	router.HandleFunc("/student/{enrollment_no}", controller.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{enrollment_no}", controller.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students", controller.GetAllStudents).Methods("GET")

	// Student routes (girls)
	router.HandleFunc("/studentg", controller.AddStudentG).Methods("POST")
	router.HandleFunc("/studentg/{enrollment_no}", controller.GetStudentG).Methods("GET")
	router.HandleFunc("/studentg/{enrollment_no}", controller.UpdateStudentG).Methods("PUT")
	router.HandleFunc("/studentg/{enrollment_no}", controller.DeleteStudentG).Methods("DELETE")
	router.HandleFunc("/studentsg", controller.GetAllStudentsG).Methods("GET")

	// Room routes (boys and girls)
	router.HandleFunc("/room", controller.CreateRoom).Methods("POST")
	router.HandleFunc("/room/{roomNo}", controller.GetRoom).Methods("GET")

	router.HandleFunc("/roomg", controller.CreateRoomG).Methods("POST")      // Assuming you have a CreateRoomG controller method for girls
	router.HandleFunc("/roomg/{roomNo}", controller.GetRoomG).Methods("GET") // Assuming you have a GetRoomG controller method for girls

	// Authentication routes
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")

	// Serve static files
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
