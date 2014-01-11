package showcase

import (
    "net/http"
    "html/template"
    "fmt"
    "regexp"
    "github/git-go-jeansite/src/common"
)

type Page struct {
    Title       string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    re := regexp.MustCompile("/showcase/(.+)")
    matches := re.FindAllStringSubmatch(req.URL.Path, -1)
    
    if (matches == nil) {
        fmt.Println("Matches is null")
    } else {
        showcaseTitle := matches[0][1]
        importPath := common.StrCat("showcaselib/", showcaseTitle)
        importPath = common.StrCat(showcaseTitle, "/showcase.go")

        fmt.Println(importPath)
    }

    p := Page{Title: "Showcase"}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("html/showcase.html", "html/index.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}