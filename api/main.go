package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"path"
	"path/filepath"

	"github.com/lithammer/shortuuid"
)

var ctoas []Ctoa

func adding(c *gin.Context) {

	ctoa := Ctoa{}
	c.BindJSON(&ctoa)

	ctoa.ID = shortuuid.New()

	ctoas = append(ctoas, ctoa)

	log.Printf("%v\n", ctoa)

	c.JSON(200, ctoa)
}

func getting(c *gin.Context) {
	c.JSON(200, ctoas)
}

func updating(c *gin.Context) {
	ctoa := Ctoa{}
	c.BindJSON(&ctoa)

	id := c.Param("id")

	found := false
	for i, j := range ctoas {
		if j.ID == id {
			// Update Ctoa
			found = true
			ctoas[i] = ctoa
			c.JSON(200, ctoa)
		}
	}

	if !found {
		c.JSON(500, nil)
	}

}

func deleting(c *gin.Context) {

	id := c.Param("id")

	found := false
	for i, j := range ctoas {
		if j.ID == id {
			// Remove Ctoa
			found = true
			ctoas[i] = ctoas[len(ctoas)-1]
			ctoas[len(ctoas)-1] = Ctoa{}
			ctoas = ctoas[:len(ctoas)-1]
			c.JSON(200, j)
		}
	}

	if !found {
		c.JSON(500, nil)
	}

}

func main() {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./index.html")
		} else {
			c.File("./" + path.Join(dir, file))
		}

	})

	router.POST("./api/ctoa", adding)
	router.GET("./api/ctoa", getting)
	router.PUT("./api/ctoa/:id", updating)
	router.DELETE("./api/ctoa/:id", deleting)

	err := router.Run(":80")
	if err != nil {
		panic(err)
	}
}
