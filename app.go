package waterApi

import (
	"database/sql"
	"github.com/shengbojia/gorouter"
	"net/http"
)

type App struct {
	Router *gorouter.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.POST("/drinks", app.createDrink)
	app.Router.GET("/drinks", app.getDrinks)
}

func (app *App) createDrink(w http.ResponseWriter, r *http.Request) {

}

func (app *App) getDrinks(w http.ResponseWriter, r *http.Request) {

}