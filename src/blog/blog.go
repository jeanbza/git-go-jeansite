package blog

import (
    "github/git-go-jeansite/src/common"
    "regexp"
    "bytes"
    "io/ioutil"
    "html/template"
    "net/http"
)

type Post struct {
    FileName   string
    Title      string
    Date       string
    Content    template.HTML
}

type Page struct {
    Title string
    Posts []Post
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
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

    common.CheckError(err)

    return Post{FileName: fileName, Title: title, Date: dateString.String(), Content: contentHTML}
}