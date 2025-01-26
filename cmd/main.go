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
    tmplt, err := template.ParseFiles(
        "./public/views/index.html",
    )
    if err != nil {
        log.Fatalf("could not initialize templates: %+v", err)
    }

    // Setup Server Endpoints
    e := echo.New()
    e.Renderer = endpoints.NewTemplateRenderer(tmplt)
    e.Use(middleware.Logger())
    e.Static("/css", "css")

    e.GET("/", endpoints.HandleIndex)

    // Start Server
    e.Logger.Fatal(e.Start(":58008"))
}
