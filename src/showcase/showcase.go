package showcase

import (
    "net/http"
    "html/template"
    "regexp"
    "bytes"
    "io/ioutil"
    "github/git-go-jeansite/src/common"
)

func GetPage(rw http.ResponseWriter, req *http.Request) {
    re := regexp.MustCompile("/showcase/(.+)")
    matches := re.FindAllStringSubmatch(req.URL.Path, -1)
    
    if (matches == nil) {
        type Page struct {
            Title       string
            ShowCase    bool
        }

        p := Page{Title: "default", ShowCase: false}

        tmpl := make(map[string]*template.Template)
        tmpl["showcase.html"] = template.Must(template.ParseFiles("html/showcase.html", "html/index.html", "showcaselib/default/showcase.html"))
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
        Title       string
        ShowCase    ShowCase
    }

    filePaths := []string{
        "showcaselib/ember_widget/templates/app.html",
        "showcaselib/ember_widget/templates/widget.html",
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
    p := Page{Title: "ember_widget", ShowCase: s}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("html/showcase.html", "html/index.html", "showcaselib/ember_widget/showcase.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}