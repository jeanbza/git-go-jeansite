package main

import (
    "git-go-jeansite/src/blog"
    "git-go-jeansite/src/about"
    "git-go-jeansite/src/showcase"
    "git-go-jeansite/src/common"
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

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    fileServer = http.StripPrefix("/resources/", http.FileServer(http.Dir("resources")))
    http.Handle("/resources/", fileServer)

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