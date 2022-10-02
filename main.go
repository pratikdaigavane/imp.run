package main

import (
	"github.com/gin-gonic/gin"
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

	err := re.Session.Query(`INSERT INTO url_shortener.urls (short_code, created_at, url) VALUES (?, toUnixTimestamp(now()), ?)`,
		req.ShortCode, req.Url).Exec()

	if err != nil {
		log.Println(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, req)
}

func getUrl(c *gin.Context) {
	shortCode := c.Param("sc")
	var url string
	sleepTimeOutput := re.Session.Query("SELECT url FROM url_shortener.urls WHERE short_code = ? LIMIT 1", shortCode).Iter()
	sleepTimeOutput.Scan(&url)
	c.Redirect(http.StatusPermanentRedirect, url)
}

func main() {
	router := gin.Default()
	re.Connect()
	defer re.Close()
	router.GET("/:sc", getUrl)
	router.POST("/insert", insertUrl)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
