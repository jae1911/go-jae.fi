package main

import (
    "io/ioutil"
    "net/http"
    "os"
    "log"
    "html/template"
    "errors"
    "strings"
    "regexp"

    "github.com/gin-gonic/gin"
    "github.com/russross/blackfriday"
    "github.com/missdeer/gin-minify"
)

// For content pages
type Post struct {
    Title   string
    Content template.HTML
}

func main() {
    // Web service
    r := gin.New()
    r.Use(gin.Recovery())

    // Streaming
    stream_online := false

    // Compression
    r.Use(minify.Minify(minify.IgnoreHTML()))

    r.Delims("{{", "}}")

    // Personal info endpoint
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
            "stream_online": stream_online,
        })
    })

    // Signatures
    r.GET("/signatures/malware", func(c *gin.Context) {
        signatures, err := ioutil.ReadFile("static/signatures.txt")
        if err != nil {
            log.Fatal(err)
        }

        c.String(http.StatusOK, string(signatures))
    })

    // IP range location feed
    r.GET("/ip/ranges", func(c *gin.Context) {
        message := "2001:67c:2724::/48,FI,,,\n2a0e:8f02:f01f::/48,DE,,,\n2a12:4946:9900::/40,BE,,,\n89.46.97.0/24,FI,,,"

        c.String(http.StatusOK, message)
    })

    // Webring members endpoint
    // TODO: use a database instead of a plain file
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

    // Matrix server delegation
    r.GET("/.well-known/matrix/server", func(c *gin.Context) {
        c.JSON(200, gin.H {
            "m.server": "matrix.jae.fi:443",
        })
    })

    // Matrix client configuration
    r.GET("/.well-known/matrix/client", func(c *gin.Context) {
        c.String(http.StatusOK, "{\"m.homeserver\":{\"base_url\":\"https://matrix.jae.fi/\"},\"m.identity_server\":{\"base_url\":\"https://matrix.jae.fi\"}}")
    })

    // Static routes & templates
    r.LoadHTMLGlob("templates/**/*.tmpl")
    r.Static("/assets", "./static")

    // Main page, cannot be handled by :page
    r.GET("/", func(c *gin.Context) {
        postContent, err := ioutil.ReadFile("content/index.md")
        if err != nil {
            log.Fatal(err)
        }

        postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(postContent)))

        post := Post{Title: "Main Page", Content: postHTML}

        c.HTML(http.StatusOK, "globals/complete.tmpl", gin.H{
            "title":        post.Title,
            "content":      post.Content,
            "currentsite":  "Jae's Website",
            "stream_online": stream_online,
        })
    })

    // Get routes
    r.GET("/:page", func(c *gin.Context) {
        requestedPage := strings.ToLower(c.Param("page"))
        contentLocation := "content/" + requestedPage + ".md"

        // Check if file exists
        if _, err := os.Stat("content/" + requestedPage + ".md"); errors.Is(err, os.ErrNotExist){
            contentLocation = "content/404.md"
        }

        postContent, err := ioutil.ReadFile(contentLocation)
        if err != nil {
            log.Fatal(err)
        }

        if strings.HasPrefix(string(postContent), "redir:") {
            // Get redir URI
            re, _ := regexp.Compile(`^redir:`)

            raw_uri := re.ReplaceAllString(string(postContent), "")

            c.Redirect(http.StatusMovedPermanently, raw_uri)
        } else {
            postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(postContent)))

            post := Post{Title: requestedPage, Content: postHTML}

            c.HTML(http.StatusOK, "globals/complete.tmpl", gin.H{
                "title": post.Title,
                "content": post.Content,
                "stream_online": stream_online,
                "path": "/" + requestedPage,
            })
        }    

    })

    r.Run(":2021")
}
