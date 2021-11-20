package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// For the future
type Website struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Owner string `json:"owner"`
}

type Members struct {
	Members []Website `json:"members"`
}

func main() {
	// Web service
	r := gin.New()
	r.Use(gin.Recovery())

	// Compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Delims("{{", "}}")

	r.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":       "Jae",
			"familyname": "Lo Presti",
			"job":        "Back-end Developer",
			"birthday":   "2001-04-28",
			"location":   "Helsinki, Finland, Europe, Earth, Alpha Quadrant",
			"email":      "me@jae.fi",
			"matrix":     "@me:jae.fi",
			"fediverse":  "@jae@mastodon.tedomum.net",
			"pronouns":   "She/Her",
		})
	})

	r.GET("/webring/members", func(c *gin.Context) {
		jsonfile, err := ioutil.ReadFile("./webring/members.json")
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to read members",
			})
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		c.Data(http.StatusOK, "application/json", jsonfile)
	})

	r.LoadHTMLGlob("templates/**/*.tmpl")
	r.Static("/assets", "./static")
	r.Static("/.well-known", "./wellknown")

	r.GET("/", func(c *gin.Context) {
		jsonfile, err := os.Open("./webring/members.json")
		if err != nil {
			fmt.Println(err)
		}
		var members Members

		byteValue, _ := ioutil.ReadAll(jsonfile)

		defer jsonfile.Close()

		json.Unmarshal(byteValue, &members)

		var previousSite, nextSite, randomsite string
		currentSite := "https://jae.fi"
		randSite := rand.Intn(len(members.Members) - 1)

		for i := 0; i < len(members.Members); i++ {
			if members.Members[i].Url == currentSite {
				if i-1 < 0 {
					previousSite = members.Members[len(members.Members)-1].Url
				} else {
					previousSite = members.Members[i-1].Url
				}

				if i+1 > len(members.Members) {
					nextSite = members.Members[i-1].Url
				} else {
					nextSite = members.Members[i+1].Url
				}

				randomsite = members.Members[randSite].Url
			}
		}

		c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
			"title":        "Main page",
			"currentsite":  "Jae's Website",
			"currentowner": "Jae Lo Presti",
			"prevsite":     previousSite,
			"nextsite":     nextSite,
			"randomsite":   randomsite,
			"path":         c.FullPath(),
		})
	})

	// Gallery
	r.GET("/gallery", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/gallery.tmpl", gin.H{
			"title": "Gallery",
			"path":  c.FullPath(),
		})
	})

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "home/404.tmpl", gin.H{
			"title": "404",
		})
	})

	// Timeline
	r.GET("/timeline", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/timeline.tmpl", gin.H{
			"title": "Timeline",
		})
	})

	// Contact
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/contact.tmpl", gin.H{
			"title": "Contact",
		})
	})

	// Donations
	r.GET("/donation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/donation.tmpl", gin.H{
			"title": "About Donations",
		})
	})

	// Stuff
	r.GET("/stuff", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/stuff.tmpl", gin.H{
			"title": "Current Hardware",
		})
	})

	// Stream
	r.GET("/stream", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/stream.tmpl", gin.H {
			"title": "Livestream",
		})
	})

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.Run(":2021")
}
