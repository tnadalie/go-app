package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/tnadalie/go-app/app/handler"
	"github.com/tnadalie/go-app/app/model"
	"github.com/tnadalie/go-app/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the clients
	a.Get("/clients", a.GetAllClients)
	a.Post("/clients", a.CreateClient)
	a.Get("/clients/{title}", a.GetClient)
	a.Put("/clients/{title}", a.UpdateClient)
	a.Delete("/clients/{title}", a.DeleteClient)
	a.Put("/clients/{title}/archive", a.ArchiveClient)
	a.Delete("/clients/{title}/archive", a.RestoreClient)

	// Routing for handling the tasks
	// a.Get("/clients/{title}/tasks", a.GetAllTasks)
	// a.Post("/clients/{title}/tasks", a.CreateTask)
	// a.Get("/clients/{title}/tasks/{id:[0-9]+}", a.GetTask)
	// a.Put("/clients/{title}/tasks/{id:[0-9]+}", a.UpdateTask)
	// a.Delete("/clients/{title}/tasks/{id:[0-9]+}", a.DeleteTask)
	// a.Put("/clients/{title}/tasks/{id:[0-9]+}/complete", a.CompleteTask)
	// a.Delete("/clients/{title}/tasks/{id:[0-9]+}/complete", a.UndoTask)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// GetAllclients Handlers
func (a *App) GetAllclients(w http.ResponseWriter, r *http.Request) {
	handler.GetAllclients(a.DB, w, r)
}

func (a *App) CreateClient(w http.ResponseWriter, r *http.Request) {
	handler.CreateClient(a.DB, w, r)
}

func (a *App) GetClient(w http.ResponseWriter, r *http.Request) {
	handler.GetClient(a.DB, w, r)
}

func (a *App) UpdateClient(w http.ResponseWriter, r *http.Request) {
	handler.UpdateClient(a.DB, w, r)
}

func (a *App) DeleteClient(w http.ResponseWriter, r *http.Request) {
	handler.DeleteClient(a.DB, w, r)
}

func (a *App) ArchiveClient(w http.ResponseWriter, r *http.Request) {
	handler.ArchiveClient(a.DB, w, r)
}

func (a *App) RestoreClient(w http.ResponseWriter, r *http.Request) {
	handler.RestoreClient(a.DB, w, r)
}

//
// /*
// ** Tasks Handlers
//  */
// func (a *App) GetAllTasks(w http.ResponseWriter, r *http.Request) {
// 	handler.GetAllTasks(a.DB, w, r)
// }
//
// func (a *App) CreateTask(w http.ResponseWriter, r *http.Request) {
// 	handler.CreateTask(a.DB, w, r)
// }
//
// func (a *App) GetTask(w http.ResponseWriter, r *http.Request) {
// 	handler.GetTask(a.DB, w, r)
// }
//
// func (a *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateTask(a.DB, w, r)
// }
//
// func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	handler.DeleteTask(a.DB, w, r)
// }
//
// func (a *App) CompleteTask(w http.ResponseWriter, r *http.Request) {
// 	handler.CompleteTask(a.DB, w, r)
// }
//
// func (a *App) UndoTask(w http.ResponseWriter, r *http.Request) {
// 	handler.UndoTask(a.DB, w, r)
// }

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
