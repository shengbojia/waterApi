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

	defer db.Close()

	router := gorouter.New()
	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func setupRouter(router *gorouter.Router) {
	router.POST("/drinks", createDrink)
	router.GET("/drinks", getDrinks)
}

func createDrink(w http.ResponseWriter, r *http.Request) {

}

func getDrinks(w http.ResponseWriter, r *http.Request) {

}
