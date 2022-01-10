package main

import (
    "io/ioutil"
    "os"
    "html/template"
    "fmt"
    "path"
    "strings"

    "github.com/russross/blackfriday"
)

// Page structure
type Page struct {
    Title string
    Content template.HTML
}

func main() {
    // Create dist
    _ = os.Mkdir("./dist", os.ModePerm)

    // Read content folder
    files, err := ioutil.ReadDir("./content")
    check(err)

    for _, f := range files {
        fmt.Println("Converting", f.Name(), "to", convert_ext_to_html(f.Name()))

        old_path := "content/" + f.Name()
        new_path := "dist/" + convert_ext_to_html(f.Name())

        dat, err := os.ReadFile(old_path)
        check(err)
        
        html := blackfriday.MarkdownCommon([]byte(dat))

        err = os.WriteFile(new_path, html, 0644)
        check(err)
    }
}

func convert_ext_to_html(fn string) string {
    return strings.TrimSuffix(fn, path.Ext(fn)) + ".html"
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
