package showcase

import (
    "net/http"
    "html/template"
    "regexp"
    "bytes"
    "io/ioutil"
    "git-go-jeansite/src/common"
)

func GetPage(rw http.ResponseWriter, req *http.Request) {
    re := regexp.MustCompile("/showcase/(.+)")
    matches := re.FindAllStringSubmatch(req.URL.Path, -1)
    
    if (matches == nil) {
        type Page struct {
            Title               string
            CurrentShowcase     string
            ShowCase            bool
        }

        p := Page{Title: "showcase", CurrentShowcase: "default", ShowCase: false}

        tmpl := make(map[string]*template.Template)
        tmpl["showcase.html"] = template.Must(template.ParseFiles("resources/html/showcase.html", "resources/html/index.html", "resources/showcases/default/showcase.html"))
        tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
    } else {
        showcaseTitle := matches[0][1]
        
        switch showcaseTitle {
            case "ember_widget":
                loadEmberWidgetShowcase(rw)
        }
    }
}

func loadEmberWidgetShowcase(rw http.ResponseWriter) {
    type ShowCase struct {
        Additional  template.HTML
    }

    type Page struct {
        Title               string
        CurrentShowcase     string
        ShowCase            ShowCase
    }

    filePaths := []string{
        "resources/showcases/ember_widget/templates/app.html",
        "resources/showcases/ember_widget/templates/widget.html",
    }

    var contentString bytes.Buffer

    for _, filePath := range filePaths {
        // Read the file's contents
        contentByte, err := ioutil.ReadFile(filePath)
        common.CheckError(err)

        contentString.WriteString(string(contentByte))
    }

    contentHTML := template.HTML(contentString.String())

    s := ShowCase{Additional: contentHTML}
    p := Page{Title: "showcase", CurrentShowcase: "ember_widget", ShowCase: s}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("resources/html/showcase.html", "resources/html/index.html", "resources/showcases/ember_widget/showcase.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}