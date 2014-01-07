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
    tmpl["about.html"] = template.Must(template.ParseFiles("html/about.html", "html/index.html"))
    tmpl["about.html"].ExecuteTemplate(rw, "base", p)
}