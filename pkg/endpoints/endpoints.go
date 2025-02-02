package endpoints

import (
	"fmt"
	"net/http"
    "go-site/pkg/database"
    "go-site/pkg/global"

	"github.com/labstack/echo/v4"
)

type Header struct {
    Title string
}

type IndexPage struct {
    Header
    Data string
}

type ChessPage struct {
    Header
    Data string
}

type StorePage struct {
    Header
    Products []database.Product
}

func HandleIndex(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleIndex():")
    }

    return c.Render(http.StatusOK, "index.html", IndexPage{
        Header: Header{
            Title: "My website",
        },
        Data: "some data here",
    })
}

func HandleChess(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleChess():")
    }

    return c.Render(http.StatusOK, "chess", ChessPage{
        Header: Header{
            Title: "Chess",
        },
        Data: "Here's some data having to do with chess",
    })
}

func HandleStore(c echo.Context) error {
    if global.Debug {
        fmt.Println("HandleStore():")
    }

    return c.Render(http.StatusOK, "store", StorePage{
        Header: Header{
            Title: "Store",
        },
        Products: []database.Product{},
    })
}

