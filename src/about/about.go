package about

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "about",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["about.html"] = template.Must(template.ParseFiles("resources/html/about.html", "resources/html/index.html"))
    tmpl["about.html"].ExecuteTemplate(rw, "base", p)
}