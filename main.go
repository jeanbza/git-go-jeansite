package main

import (
    "regexp"
    "bytes"
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
    Paragraphs []template.HTML
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

func loadPage(fileName string) (Post) {
    // Split the filename to get the title
    re := regexp.MustCompile("_([a-zA-Z0-9 ]+)")
    title := re.FindAllStringSubmatch(fileName, -1)[0][1]

    // Read the file's contents
    contentByte, err := ioutil.ReadFile(fileName)
    contentSplitBytes := bytes.Split(contentByte, []byte("\n"))
    var contentSplitString []template.HTML

    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }

    for _, element := range contentSplitBytes {
        contentSplitString = append(contentSplitString, template.HTML(string(element)))
    }

    return Post{Title: title, Paragraphs: contentSplitString}
}

func blogPage(rw http.ResponseWriter, req *http.Request) {
    var posts []Post
    var buffer bytes.Buffer
    postPaths, _ := ioutil.ReadDir("posts")
    
    for _, element := range postPaths {
        buffer.Reset()
        buffer.WriteString("posts/")
        buffer.WriteString(element.Name())
        posts = append(posts, loadPage(buffer.String()))
    }

    p := Page{
        Title: "blog",
        Posts: posts,
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