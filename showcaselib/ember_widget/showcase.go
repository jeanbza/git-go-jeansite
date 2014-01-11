type Page struct {
    Title      string
    ShowCase   ShowCase
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := loadPage()

    tmpl := make(map[string]*template.Template)
    tmpl["showcase.html"] = template.Must(template.ParseFiles("html/showcase.html", "html/index.html"))
    tmpl["showcase.html"].ExecuteTemplate(rw, "base", p)
}

func loadPage() (Page) {
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

    return Page{Title: "Showcase", Additional: contentHTML}
}