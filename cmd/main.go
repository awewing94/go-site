package main

import (
	"fmt"
	"go-site/pkg/endpoints"
	"html/template"
	"log"
	_ "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    fmt.Println("main(): ")
    // Set up
    tmpls, err := template.ParseFiles(
        "./public/views/index.html",
        "./public/views/navbar.html",

        "./public/views/chess/chess.html",

        "./public/views/store/store.html",
    )
    if err != nil {
        log.Fatalf("could not initialize templates: %+v", err)
    }

    // Setup Rendering
    e := echo.New()
    e.Renderer = endpoints.NewTemplateRenderer(tmpls)
    e.Use(middleware.Logger())
    e.Static("/css", "css")

    // Setup Endpoints
    e.GET("/", endpoints.HandleIndex)
    e.GET("/chess", endpoints.HandleChess)
    e.GET("/store", endpoints.HandleStore)

    // Start Server
    e.Logger.Fatal(e.Start(":58008"))
}
