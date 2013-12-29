package main

import (
    "io/ioutil"
    "html/template"
    "net/http"
    "os"
    "fmt"
)

type Page struct {
    Title string
    Posts []Post
}

type Post struct {
    Title   string
    Content string
}

func main() {
    http.HandleFunc("/", blogPage)
    http.HandleFunc("/blog", blogPage)
    http.HandleFunc("/blog/", blogPage)

    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/about/", aboutPage)

    http.HandleFunc("/contact", contactPage)
    http.HandleFunc("/contact/", contactPage)

    fileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    checkError(err)
}

func loadPage(title string) (Post) {
    filename := title + ".txt"
    contentByte, err := ioutil.ReadFile(filename)
    contentString := string(contentByte)
    
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }

    return Post{Title: title, Content: contentString}
}

func blogPage(rw http.ResponseWriter, req *http.Request) {
    somePost := loadPage("posts/20131228_Testpost")
    someOtherPost := loadPage("posts/20131228_Testpost2")

    p := Page{
        Title: "blog",
        Posts: []Post{
            somePost,
            someOtherPost,
        },
    }

    tmpl := make(map[string]*template.Template)
    tmpl["blog.html"] = template.Must(template.ParseFiles("html/blog.html", "html/index.html"))
    tmpl["blog.html"].ExecuteTemplate(rw, "base", p)
}

func aboutPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "about",
        Posts: nil,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["about.html"] = template.Must(template.ParseFiles("html/about.html", "html/index.html"))
    tmpl["about.html"].ExecuteTemplate(rw, "base", p)
}

func contactPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "contact",
        Posts: nil,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["contact.html"] = template.Must(template.ParseFiles("html/contact.html", "html/index.html"))
    tmpl["contact.html"].ExecuteTemplate(rw, "base", p)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}