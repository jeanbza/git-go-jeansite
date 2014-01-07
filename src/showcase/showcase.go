package showcase

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "showcase",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("html/showcase.html", "html/index.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}