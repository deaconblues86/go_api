package main


import (
    "fmt"
    "rsc.io/quote"
    "net/http"
    "github.com/gin-gonic/gin"
)

type commentBody struct {
    Author string `json:"author"`
    Content string `json:"content"`
    Timestamp string `json:"timestamp"`
    RefID int `json:"ref_id"`
    RefType string `json:"ref_type"`
}

type comment struct {
    ID int `json:"id"`
    Body commentBody `json:body`
    Comments []comment `json:comments`
}

type postBody struct {
    Title string `json:"title"`
    Author string `json:"author"`
    Content string `json:"content"`
    Timestamp string `json:"timestamp"`
}

type post struct {
    ID int `json:"id"`
    Body postBody `json:body`
    Comments []comment `json:comments`
}


func (p *post) AddComment(newComment comment) {
        p.Comments = append(p.Comments, newComment)
}


func (c *comment) AddComment(newComment comment) {
        c.Comments = append(c.Comments, newComment)
}


func getPosts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, posts)
}

func postPosts(c *gin.Context) {
    var newPost postBody

    // Call BindJSON to bind the received JSON to
    // newPost.
    if err := c.BindJSON(&newPost); err != nil {
        return
    }

    newID := findMaxPostID(posts)
    newID += 1

    // Add the new post to the slice.
    realizedPost := post{ID: newID, Body: newPost}
    posts = append(posts, realizedPost)
    c.IndentedJSON(http.StatusCreated, realizedPost)
}

func postComments(c *gin.Context) {
    var newComment commentBody

    if err := c.BindJSON(&newComment); err != nil {
        fmt.Println(err)
        return
    }

    newID := findMaxCommentID(posts)
    newID += 1

    realizedComment := comment{ID: newID, Body:newComment}
    traverseItems(posts, realizedComment)
}

var posts = []post {
    {ID: 1, Body: postBody{Title: "Test Post", Author: "Brian H", Content: "I'm a test", Timestamp: "2022-04-27T5:49Z"}},
    {ID: 2, Body: postBody{Title: "Another Post", Author: "Brian H", Content: "I'm a test", Timestamp: "2022-04-27T5:51Z"}},
}

func main() {
    fmt.Println(quote.Go())

    router := gin.Default()

    router.GET("/posts", getPosts)
    router.POST("/posts", postPosts)
    router.POST("/comments", postComments)

    router.Run("localhost:8080")
}