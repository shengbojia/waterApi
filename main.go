package waterApi

import (
	"github.com/shengbojia/gorouter"
	"log"
	"net/http"
)

func main() {
	db, err := CreateDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}

	app := &App{
		Router: gorouter.New(),
		Database: db,
	}

	defer app.Database.Close()

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
