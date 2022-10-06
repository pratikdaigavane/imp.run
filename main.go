package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pratikdaigavane/emoji-hash/models"
	re "github.com/pratikdaigavane/emoji-hash/resources"
	"log"
	"net/http"
	"time"
)

func insertUrl(c *gin.Context) {
	req := models.URL{
		CreatedAt: time.Now(),
	}
	if err := c.BindJSON(&req); err != nil {
		return
	}

	q := models.UrlsTable.InsertBuilder().Unique().Query(re.Session).BindStruct(req)

	applied, err := q.ExecCASRelease()

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if !applied {
		log.Println("Short Code Already Present")
		c.Status(http.StatusConflict)
		return
	}

	c.JSON(http.StatusCreated, req)
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

	if len(dbObj.Url) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusMovedPermanently, dbObj.Url)
}

func main() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://app.imp.run"}
	router := gin.Default()
	re.Connect()
	defer re.Close()
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://app.imp.run")
	})
	router.GET("/:sc", getUrl)
	router.POST("/insert", insertUrl)
	router.Use(cors.New(config))
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
