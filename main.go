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
    FileName   string
    Title      string
    Date       string
    Content    template.HTML
}

func main() {
    http.HandleFunc("/", blogPage)
    http.HandleFunc("/blog", blogPage)
    http.HandleFunc("/blog/", blogPage)

    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/about/", aboutPage)

    fileServer := http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts")))
    http.Handle("/fonts/", fileServer)

    fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    checkError(err)
}

func loadPage(filePath string, fullPost bool) (Post) {
    // Split the filepath to get the filename
    re := regexp.MustCompile("/(.+).txt")
    fileName := re.FindAllStringSubmatch(filePath, -1)[0][1]

    // Read the file's contents
    contentByte, err := ioutil.ReadFile(filePath)
    var contentString bytes.Buffer

    if err != nil {
        return Post{}
    }

    if (!fullPost) {
        var contentByteTrimmed []byte

        for index, element := range contentByte {
            if (index < 120) {
                contentByteTrimmed = append(contentByteTrimmed, element)
            }
        }

        contentString.WriteString(string(contentByteTrimmed))
        contentString.WriteString(".. <a href='/blog/")
        contentString.WriteString(fileName)
        contentString.WriteString("'><small>Read more</small></a></div>")
    } else {
        contentString.WriteString(string(contentByte))
    }

    contentHTML := template.HTML(contentString.String())

    // Split the filepath to get the date
    var dateString bytes.Buffer
    
    re = regexp.MustCompile("/[0-9]{6}([0-9]{2})_.*")
    day := re.FindAllStringSubmatch(filePath, -1)[0][1]
    dateString.WriteString(day)
    dateString.WriteString("-")

    re = regexp.MustCompile("/[0-9]{4}([0-9]{2})[0-9]{2}_.*")
    month := re.FindAllStringSubmatch(filePath, -1)[0][1]
    dateString.WriteString(month)
    dateString.WriteString("-")

    re = regexp.MustCompile("/([0-9]{4})[0-9]{4}_.*")
    year := re.FindAllStringSubmatch(filePath, -1)[0][1]
    dateString.WriteString(year)

    // Split the filepath to get the title
    re = regexp.MustCompile("_(.+).txt")
    title := re.FindAllStringSubmatch(filePath, -1)[0][1]

    checkError(err)

    return Post{FileName: fileName, Title: title, Date: dateString.String(), Content: contentHTML}
}

func blogPage(rw http.ResponseWriter, req *http.Request) {
    var posts []Post
    var filePath bytes.Buffer

    re := regexp.MustCompile("/blog/(.+)")
    matches := re.FindAllStringSubmatch(req.URL.Path, -1)
    
    if (matches == nil) {
        // Grabs all posts in the posts directory, loads them into a Page struct, and appends to the posts array
        postPaths, _ := ioutil.ReadDir("posts")
        
        for i := len(postPaths)-1; i >= 0; i-- {
            element := postPaths[i]
            filePath.Reset()
            filePath.WriteString("posts/")
            filePath.WriteString(element.Name())
            posts = append(posts, loadPage(filePath.String(), false))
        }
    } else {
        // Grabs specific post
        post := matches[0][1]

        filePath.WriteString("posts/")
        filePath.WriteString(post)
        filePath.WriteString(".txt")

        posts = append(posts, loadPage(filePath.String(), true))
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

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}