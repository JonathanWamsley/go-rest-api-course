package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/tutorialedge/production-ready-api/internal/comment"
	"github.com/tutorialedge/production-ready-api/internal/database"
	transportHTTP "github.com/tutorialedge/production-ready-api/internal/transport/http"
)

// App - the struct which contains pointers to the database connection
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to set up database connection")
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}

func init() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading env variables")
	}
}
