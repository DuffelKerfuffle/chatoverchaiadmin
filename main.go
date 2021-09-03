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

/*type image struct {
	pic imageupload.Image
}

type source struct {
	otherpage bool
}

type sourceinterface interface {
	Hi(bool)
}*/

type manager interface {
	Add(title, text, imgSrc string) bool
	Remove(title string) bool
	Change(title, newtitle, newText, newimg string) bool
	GetAllDocs() []documents.Doc
}

/*
func addImage(source, contentname, option string, c *gin.Context, y image, Options map[string]manager) manager {
	img, err := imageupload.Process(c.Request, source)

	fmt.Println("webpage recognised")
	if err != nil {
		log.Println(err)
		fmt.Println("hi1")
		vid, err := c.FormFile("source")
		if err != nil {
			log.Println(err)
			fmt.Println("hi2")
			panic(err)
		}
		err = c.SaveUploadedFile(vid, "../chatoverchai/images/"+contentname+".MOV")
		if err != nil {
			log.Println(err)
			fmt.Println("hi3")
			panic(err)
		}
		//panic(err)
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		log.Println(err)
		fmt.Println("hi4")
		panic(err)
	}

	y.pic = *thumb
	ddoc := Options[option]
	err2 := y.pic.Save("../chatoverchai/images/" + contentname + ".png")
	if err2 != nil {
		log.Println(err2)
		panic(err2)
	}
	return ddoc
}


func (s *source) Hi(hello bool) {
	fmt.Println("hi")
}*/

func GetDdoc(option string, Options map[string]manager) manager {
	ddoc := Options[option]
	return ddoc
}

/*
func addVideo(c *gin.Context, option, contentname string, Options map[string]manager) manager {
	vid, err := c.FormFile("mediasrc")
	if err != nil {
		log.Println(err)
		fmt.Println("hi2")
		panic(err)
	}
	err = c.SaveUploadedFile(vid, "../chatoverchai/videos/"+contentname+".mp4")
	if err != nil {
		log.Println(err)
		fmt.Println("hi3")
		panic(err)
	}
	ddoc := Options[option]
	return ddoc
}*/

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

	homeurl := "login.html"
	addurl := "login.html"
	removeurl := "login.html"
	changeurl := "login.html"
	loginurl := "login.html"

	//authenticated := false
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Delims("[[", "]]")
	r.LoadHTMLFiles("admin2.html", "newadmin.html", "adminui.html", "admin11.html", "admin33.html", "login.html")
	r.Use(static.Serve("/images", static.LocalFile("../chatoverchai/images", true)))
	r.Use(static.Serve("/js", static.LocalFile("./js", true)))

	r.POST("/login1", func(c *gin.Context) {
		Usernameinput := c.PostForm("username")
		Passwordinput := c.PostForm("password")
		fmt.Println(Usernameinput)
		if Usernameinput == username && Passwordinput == password {
			//authenticated = true
			homeurl = "adminui.html"
			addurl = "admin11.html"
			removeurl = "admin2.html"
			changeurl = "admin33.html"
			loginurl = "adminui.html"
			fmt.Println("correct!!!")
		}
		c.HTML(http.StatusOK, loginurl, nil)
	})

	r.GET("/login1", func(c *gin.Context) {
		c.HTML(http.StatusOK, loginurl, nil)
	})

	r.GET("/home1", func(c *gin.Context) {
		c.HTML(http.StatusOK, homeurl, Options)
	})

	r.GET("/add1", func(c *gin.Context) {
		c.HTML(http.StatusOK, addurl, Options)
	})

	r.GET("/change1", func(c *gin.Context) {
		c.HTML(http.StatusOK, changeurl, Options)
	})

	r.GET("/remove1", func(c *gin.Context) {
		c.HTML(http.StatusOK, removeurl, Options)
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
				if option != "videos" {
					newlink = strings.Replace(newlink, "https://drive.google.com/file/d/", "https://drive.google.com/thumbnail?id=", -1)
					newlink = strings.Replace(newlink, "/view?usp=sharing", "", -1)
				} else {
					newlink = strings.Replace(newlink, "/view?usp=sharing", "/preview", -1)
				}
				ddoc := GetDdoc(option, Options)

				ddoc.Change(contentname, newname, newtext, newlink)
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
		fmt.Println(contentname, PostContent, msrc, option)
		fmt.Println(msrc)
		for k := range Options {
			if option == k {
				if option != "videos" {
					msrc = strings.Replace(msrc, "https://drive.google.com/file/d/", "https://drive.google.com/thumbnail?id=", -1)
					msrc = strings.Replace(msrc, "/view?usp=sharing", "", -1)
				} else {
					msrc = strings.Replace(msrc, "/view?usp=sharing", "/preview", -1)
				}
				ddoc := GetDdoc(option, Options)
				ddoc.Add(contentname, PostContent, msrc)
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
