package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MikaSRahwono/initialize-go-project/infrastructure/db"
	"github.com/MikaSRahwono/initialize-go-project/infrastructure/router"
	"github.com/joho/godotenv"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

// func getController() controllers.AuthorController {
// 	authorRepo := repository.NewAuthorRepo(dbHandler)
// 	authorInteractor := usecases.NewAuthorInteractor(authorRepo)
// 	authorController := controllers.NewAuthorController(authorInteractor)
// 	return *authorController
// }

func SetupRouter() {
	// bookController := getBookController()
	// authorController := getAuthorController()

	// httpRouter.POST("/book/add", bookController.Add)
	// httpRouter.GET("/book", bookController.FindAll)
	// httpRouter.POST("/author/add", authorController.Add)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})

	var err error
	dbHandler, err = db.NewDBHandler(mongoURI, dbName)
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}

	httpRouter.SERVE(":" + port)
}
