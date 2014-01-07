package main

import (
    "github/git-go-jeansite/src/blog"
    "github/git-go-jeansite/src/about"
    "github/git-go-jeansite/src/showcase"
    "github/git-go-jeansite/src/common"
    "net/http"
)

func main() {
    http.HandleFunc("/", blogPage)
    http.HandleFunc("/blog", blogPage)
    http.HandleFunc("/blog/", blogPage)

    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/about/", aboutPage)

    http.HandleFunc("/showcase", showcasePage)
    http.HandleFunc("/showcase/", showcasePage)

    fileServer := http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts")))
    http.Handle("/fonts/", fileServer)

    fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    common.CheckError(err)
}

func blogPage(rw http.ResponseWriter, req *http.Request) {
    blog.GetPage(rw, req)
}

func aboutPage(rw http.ResponseWriter, req *http.Request) {
    about.GetPage(rw, req)
}

func showcasePage(rw http.ResponseWriter, req *http.Request) {
    showcase.GetPage(rw, req)
}