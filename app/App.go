package backend

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"local.package.ytbdtc"
)

func NewApp() *App {
	engine := gin.Default()
	client := &http.Client{}
	return &App{Engine: engine, Client: client}
}

type App struct {
	Engine *gin.Engine
	Client *http.Client
	Key    string
}

func (app *App) Run(port int) {
	log.Printf("Proxy Server Running at\n\t http://localhost:%d", port)
	if err := app.Engine.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatalf("Server failed: %v", err)
		return
	}
}

func (app *App) SetKey(key string) {
	app.Key = key
}

func (app *App) SetRoute() {
	app.Engine.GET("/search", func(c *gin.Context) {
		status, res := app.searchResponse(c)
		c.JSON(status, res)
	})
}

func (app *App) searchResponse(c *gin.Context) (int, *Response) {
	sp, err := app.contextToSearchParam(c)
	if err != nil {
		return http.StatusBadRequest, &Response{Success: false, Message: "No Query"}
	}
	sp.SetKey(app.Key)
	req, err := http.NewRequest("GET", sp.ToURL(), nil)
	if err != nil {
		return http.StatusBadRequest, &Response{Success: false, Message: "Error While Creating Request"}
	}
	resp, err := app.Client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, &Response{Success: false, Message: "Error While Getting Response"}
	}
	var responseItem ytbdtc.SearchResponse
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, &Response{Success: false, Message: "Error While Reading All Response"}
	}
	err = json.Unmarshal(text, &responseItem)
	if err != nil {
		return http.StatusInternalServerError, &Response{Success: false, Message: "Error While Unmarshaling"}
	}
	return http.StatusOK, &Response{Items: responseItem.Items, NextPageToken: responseItem.NextPageToken, Success: true}
}

func (app *App) contextToSearchParam(c *gin.Context) (*ytbdtc.SearchParam, error) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		return &ytbdtc.SearchParam{}, ErrNoQuery
	}
	typ := c.DefaultQuery("type", "")
	nextPageToken := c.DefaultQuery("nextPageToken", "")
	sp := ytbdtc.CreateNewSearchParam(q, typ)
	sp.SetNextPageToken(nextPageToken)

	return sp, nil
}
