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
            case "ember_treetable":
                loadEmberTreetableShowcase(rw)
            case "jquery_treetable":
                loadJqueryTreetableShowcase(rw)
            case "d3_concerts":
                loadD3ConcertsShowcase(rw)
            case "circle_wave":
                loadCircleWaveShowcase(rw)
        }
    }
}

func loadJqueryTreetableShowcase(rw http.ResponseWriter) {
    type ShowCase struct {}

    type Page struct {
        Title               string
        CurrentShowcase     string
        ShowCase            ShowCase
    }

    s := ShowCase{}
    p := Page{Title: "showcase", CurrentShowcase: "jquery_treetable", ShowCase: s}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("resources/html/showcase.html", "resources/html/index.html", "resources/showcases/jquery_treetable/showcase.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}

func loadD3ConcertsShowcase(rw http.ResponseWriter) {
    type ShowCase struct {}

    type Page struct {
        Title               string
        CurrentShowcase     string
        ShowCase            ShowCase
    }

    s := ShowCase{}
    p := Page{Title: "showcase", CurrentShowcase: "d3_concerts", ShowCase: s}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("resources/html/showcase.html", "resources/html/index.html", "resources/showcases/d3_concerts/showcase.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}

func loadEmberTreetableShowcase(rw http.ResponseWriter) {
    type ShowCase struct {
        Additional  template.HTML
    }

    type Page struct {
        Title               string
        CurrentShowcase     string
        ShowCase            ShowCase
    }

    filePaths := []string{
        "resources/showcases/ember_treetable/templates/app.html",
        "resources/showcases/ember_treetable/templates/tree.html",
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
    p := Page{Title: "showcase", CurrentShowcase: "ember_treetable", ShowCase: s}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("resources/html/showcase.html", "resources/html/index.html", "resources/showcases/ember_treetable/showcase.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
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

func loadCircleWaveShowcase(rw http.ResponseWriter) {
    type ShowCase struct {}

    type Page struct {
        Title               string
        CurrentShowcase     string
        ShowCase            ShowCase
    }

    s := ShowCase{}
    p := Page{Title: "showcase", CurrentShowcase: "circle_wave", ShowCase: s}

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("resources/html/showcase.html", "resources/html/index.html", "resources/showcases/circle_wave/showcase.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}