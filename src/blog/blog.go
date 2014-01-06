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
    FileName    string
    Title       string
    Date        string
    Content     template.HTML
    OneOfMany   bool
}

var trimByteLength = 200

func GetPage(rw http.ResponseWriter, req *http.Request) {
    re := regexp.MustCompile("/blog/(.+)")
    matches := re.FindAllStringSubmatch(req.URL.Path, -1)
    fullPost := matches == nil

    if (fullPost) {
        // Grabs all posts in the posts directory, loads them into a Page struct, and appends to the posts array
        loadPosts(rw)
    } else {
        // Grabs specific post
        postTitle := matches[0][1]
        loadPost(rw, postTitle)
    }
}

func loadPost(rw http.ResponseWriter, postTitle string) {
    var filePath bytes.Buffer

    type Page struct {
        Title       string
        Post        Post
        SinglePost  bool
    }

    filePath.WriteString("posts/")
    filePath.WriteString(postTitle)
    filePath.WriteString(".txt")

    post := loadPage(filePath.String(), true)

    p := Page{
        Title: "blog",
        Post: post,
        SinglePost: true,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["blog.html"] = template.Must(template.ParseFiles("html/post.html", "html/blog.html", "html/index.html"))
    tmpl["blog.html"].ExecuteTemplate(rw, "base", p)
}

func loadPosts(rw http.ResponseWriter) {
    var filePath bytes.Buffer
    var posts []Post

    type Page struct {
        Title       string
        Posts       []Post
        SinglePost  bool
    }

    postPaths, _ := ioutil.ReadDir("posts")
    
    for i := len(postPaths)-1; i >= 0; i-- {
        element := postPaths[i]
        filePath.Reset()
        filePath.WriteString("posts/")
        filePath.WriteString(element.Name())
        posts = append(posts, loadPage(filePath.String(), false))
    }

    p := Page{
        Title: "blog",
        Posts: posts,
        SinglePost: false,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["blog.html"] = template.Must(template.ParseFiles("html/post.html", "html/blog.html", "html/index.html"))
    tmpl["blog.html"].ExecuteTemplate(rw, "base", p)
}

func loadPage(filePath string, fullPost bool) (Post) {
    var contentString bytes.Buffer
    fileName :=  parseFilePathForFileName(filePath)

    // Read the file's contents
    contentByte, err := ioutil.ReadFile(filePath)
    common.CheckError(err)

    if (!fullPost && len(contentByte) > trimByteLength) {
        var contentByteTrimmed []byte

        for index, element := range contentByte {
            if (index < trimByteLength) {
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

    return Post{FileName: fileName, Title: parseFilePathForTitle(filePath), Date: parseFilePathForDate(filePath), Content: contentHTML, OneOfMany: !fullPost}
}

func parseFilePathForFileName(filePath string) (string) {
    re := regexp.MustCompile("/(.+).txt")
    fileName := re.FindAllStringSubmatch(filePath, -1)[0][1]

    return fileName
}

func parseFilePathForTitle(filePath string) (string) {
    re := regexp.MustCompile("_(.+).txt")
    title := re.FindAllStringSubmatch(filePath, -1)[0][1]

    return title
}

/**
 * Splits the file path for its date
 */
func parseFilePathForDate(filePath string) (string) {
    var dateString bytes.Buffer
    
    re := regexp.MustCompile("/[0-9]{6}([0-9]{2})_.*")
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

    return dateString.String()
}