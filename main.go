package main

import (
	"fmt"
	"net/http"

	//"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var profile []Profile

func GetProfile(c *gin.Context) {
 c.JSON(http.StatusOK, gin.H{"list": profile})
}
func main() {
	crislana := Profile{Name: "Crislana", School: "ASU", Year: "Sophomore", Goals: "To be the very best like no one ever was"}
	fmt.Printf("%+v\n", crislana)

	profile = append(profile, Profile{Name: "Crislana", School: "ASU", Year: "Sophomore", Goals: "To be the very best like no one ever was"})

 	r := gin.Default()
 	r.GET("/api/profile", GetProfile)
 	r.Run(":8090")
	/*
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./codehouse-vue/dist", false)))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Go and Gin")
	})
	r.Run(":8090")
	*/
}