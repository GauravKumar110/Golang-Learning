package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

type DbData struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"date"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Pin     string    `json:"pin"`
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("get").
		Path("/society/{id}").
		HandlerFunc(app.getSocietyDetail)
}

func (app *App) getSocietyDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, name,address,pin,created_at FROM `society` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Name, &dbdata.Address, &dbdata.Pin, &dbdata.Date)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	log.Println("You fetched a thing!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}
