package main

import (
	documents "chatoverchaiadmin/documentsforadmin"
	"chatoverchaiadmin/storage"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type manager interface {
	Add(title, text string) bool
	Remove(title string) bool
	Change(title, newtitle, newText string) bool
	GetAllDocs() []documents.Doc
}

func GetDdoc(option string, Options map[string]manager) manager {
	ddoc := Options[option]
	return ddoc
}

func main() {

	Options := make(map[string]manager)

	poetryManager := documents.DocumentManager{}
	blogManager := documents.DocumentManager{}
	previousManager := documents.DocumentManager{}
	upcomingManager := documents.DocumentManager{}
	photoDocManager := documents.DocumentManager{}
	vidManager := documents.DocumentManager{}

	Options["blog"] = &blogManager
	Options["poetry"] = &poetryManager
	Options["previous"] = &previousManager
	Options["upcoming"] = &upcomingManager
	Options["photo"] = &photoDocManager
	Options["videos"] = &vidManager

	storage.Load(&poetryManager, "poetry")
	storage.Load(&blogManager, "blog")
	storage.Load(&previousManager, "previous")
	storage.Load(&upcomingManager, "upcoming")
	storage.Load(&photoDocManager, "photo")
	storage.Load(&vidManager, "videos")

	fmt.Println(blogManager)
	fmt.Println(poetryManager)
	fmt.Println(previousManager)
	fmt.Println(upcomingManager)
	fmt.Println(photoDocManager)
	fmt.Println(vidManager)

	username := "ChatOverChai"
	password := "ChatUnderChai@321"

	homeurl := "adminui.html"
	addurl := "admin11.html"
	removeurl := "admin2.html"
	changeurl := "admin33.html"
	loginurl := "login.html"

	gin.SetMode(gin.ReleaseMode)
	var authorised []string
	r := gin.Default()
	r.Delims("[[", "]]")
	r.LoadHTMLFiles("admin2.html", "newadmin.html", "adminui.html", "admin11.html", "admin33.html", "login.html")
	r.Use(static.Serve("/images", static.LocalFile("../chatoverchai/images", true)))
	r.Use(static.Serve("/js", static.LocalFile("./js", true)))

	r.GET("/login1", func(c *gin.Context) {
		auth := false
		for _, v := range authorised {
			if c.ClientIP() == v {
				auth = true
			}
		}
		fmt.Println("authorised")
		if auth {
			c.Redirect(http.StatusMovedPermanently, "/home1")
		} else {
			c.HTML(http.StatusOK, loginurl, nil)
		}
	})

	r.POST("/login1", func(c *gin.Context) {
		UserName := c.PostForm("username")
		Password := c.PostForm("password")
		if UserName == username && Password == password {
			_, err := c.Cookie("auth")

			if err != nil {
				c.SetCookie("auth", "true", 300, "/", "localhost", false, true)
			}
		}
		c.HTML(http.StatusOK, homeurl, nil)
	})

	r.GET("/home1", func(c *gin.Context) {
		auth1, err := c.Cookie("auth")
		if err != nil {
			c.HTML(403, loginurl, Options)
		}
		if auth1 == "true" {
			c.HTML(http.StatusOK, homeurl, Options)
		}
	})

	r.GET("/add1", func(c *gin.Context) {
		auth1, err := c.Cookie("auth")
		if err != nil {
			c.HTML(403, loginurl, Options)
		}
		if auth1 == "true" {
			c.HTML(http.StatusOK, addurl, Options)
		}
	})

	r.GET("/change1", func(c *gin.Context) {
		auth1, err := c.Cookie("auth")
		if err != nil {
			c.HTML(403, loginurl, Options)
		}
		if auth1 == "true" {
			c.HTML(http.StatusOK, changeurl, Options)
		}
	})

	r.GET("/remove1", func(c *gin.Context) {
		auth1, err := c.Cookie("auth")
		if err != nil {
			c.HTML(403, loginurl, Options)
		}
		if auth1 == "true" {
			c.HTML(http.StatusOK, removeurl, Options)
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home1")
	})
	r.POST("/remove1", func(c *gin.Context) {
		contentname := c.PostForm("removename")
		option := c.PostForm("option")

		for k := range Options {
			if option == k {

				ddoc := Options[option]

				ddoc.Remove(contentname)
				fmt.Println(ddoc)
				storage.Remove(option)
				storage.Store(ddoc, option)

				break
			}
		}
		c.HTML(http.StatusOK, "admin2.html", Options)
	})
	r.POST("/change1", func(c *gin.Context) {
		contentname := c.PostForm("changename")
		newname := c.PostForm("newname")
		newtext := c.PostForm("newmessage")
		newlink := c.PostForm("newlink")
		option := c.PostForm("option")
		for k := range Options {
			if option == k {
				if option == "photo" {
					newlink = strings.Replace(newlink, "https://drive.google.com/file/d/", "https://drive.google.com/thumbnail?id=", -1)
					newlink = strings.Replace(newlink, "/view?usp=sharing", "", -1)
				} else if option == "videos" {
					newlink = strings.Replace(newlink, "/view?usp=sharing", "/preview", -1)
				}
				ddoc := GetDdoc(option, Options)

				if option == "videos" || option == "photo" {
					ddoc.Change(contentname, newname, newlink)
				} else {
					ddoc.Change(contentname, newname, newtext)
				}
				storage.Store(ddoc, option)
				break
			}
		}
		c.HTML(http.StatusOK, "admin33.html", Options)
	})
	r.POST("/add1", func(c *gin.Context) {
		contentname := c.PostForm("name")
		PostContent := c.PostForm("message")
		msrc := c.PostForm("msrc")
		option := c.PostForm("optionadd")
		for k := range Options {
			if option == k {
				if option == "photo" {
					msrc = strings.Replace(msrc, "https://drive.google.com/file/d/", "https://drive.google.com/thumbnail?id=", -1)
					msrc = strings.Replace(msrc, "/view?usp=sharing", "", -1)
				} else if option == "videos" {
					msrc = strings.Replace(msrc, "/view?usp=sharing", "/preview", -1)
				}
				ddoc := GetDdoc(option, Options)
				if option == "videos" || option == "photo" {
					ddoc.Add(contentname, msrc)
				} else {
					ddoc.Add(contentname, PostContent)
				}
				fmt.Println(ddoc)
				storage.Store(ddoc, option)
				break
			}
		}
		c.HTML(http.StatusOK, "admin11.html", Options)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
