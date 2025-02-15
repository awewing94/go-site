package main

import (
	"fmt"
	"go-site/pkg/endpoints"
    "go-site/pkg/chess"
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
        "./public/views/header.html",

        "./public/views/chess/chess.html",
        "./public/views/chess/board.html",

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
    e.GET("/store", endpoints.HandleStore)
    
    // Chess Endpoints
    e.GET("/chess", chess.HandleChess)
    e.POST("/chess/move", chess.HandleChessMove)

    // Start Server
    e.Logger.Fatal(e.Start(":58008"))
}

