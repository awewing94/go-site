package endpoints

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Header struct {
    Title string
}

type IndexPage struct {
    Header
    Data string
}

func HandleIndex(c echo.Context) error {
    fmt.Println("HandleIndex(): ")
    return c.Render(http.StatusOK, "index.html", IndexPage{
        Header: Header{
            Title: "My website",
        },
        Data: "some data here",
    })
}

