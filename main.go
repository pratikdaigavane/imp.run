package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pratikdaigavane/emoji-hash/models"
	re "github.com/pratikdaigavane/emoji-hash/resources"
	"log"
	"net/http"
)

type URLReq struct {
	Url       string `cql:"url" json:"url"`
	ShortCode string `cql:"short_code" json:"short-code"`
}

func insertUrl(c *gin.Context) {
	var req URLReq
	if err := c.BindJSON(&req); err != nil {
		return
	}

	dbObj := models.URL{
		req.ShortCode,
		req.Url,
		"",
	}

	q := re.Session.Query(models.UrlsTable.Insert()).BindStruct(dbObj)
	if err := q.ExecRelease(); err != nil {
		log.Println(err)
	}

	c.IndentedJSON(http.StatusCreated, req)
}

func getUrl(c *gin.Context) {
	shortCode := c.Param("sc")
	dbObj := models.URL{
		ShortCode: shortCode,
	}
	q := re.Session.Query(models.UrlsTable.Get()).BindStruct(dbObj)
	if err := q.GetRelease(&dbObj); err != nil {
		log.Println(err)
	}
	c.Redirect(http.StatusMovedPermanently, dbObj.Url)
}

func main() {
	router := gin.Default()
	re.Connect()
	defer re.Close()
	router.GET("/:sc", getUrl)
	router.POST("/insert", insertUrl)
	err := router.Run(":8081")
	if err != nil {
		return
	}
}
