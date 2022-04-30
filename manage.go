package main


import (
    "fmt"
    "strconv"
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


func getPost(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for _, p := range posts {
        if p.ID == id {
            c.IndentedJSON(http.StatusOK, p)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}


func putPost(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for i, p := range posts {
        if p.ID == id {
            if err := c.BindJSON(& posts[i].Body); err != nil {
                return
            }
            c.IndentedJSON(http.StatusOK, posts[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}


func deletePost(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    reqIndex := -1
    for i, p := range posts {
        if p.ID == id {
            reqIndex = i
            break
        }
    }

    if reqIndex == -1 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
        return
    }
    newPosts := removePost(posts, reqIndex)
    posts = newPosts
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


func setupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/posts", getPosts)
    router.POST("/posts", postPosts)

    router.GET("/posts/:id", getPost)
    router.PUT("/posts/:id", putPost)
    router.DELETE("/posts/:id", deletePost)

    router.POST("/comments", postComments)

    return router
}


var posts = []post {
    {ID: 1, Body: postBody{Title: "Test Post", Author: "Brian H", Content: "I'm a test", Timestamp: "2022-04-27T5:49Z"}},
    {ID: 2, Body: postBody{Title: "Another Post", Author: "Brian H", Content: "I'm a test", Timestamp: "2022-04-27T5:51Z"}},
}

func main() {
    r := setupRouter()
    r.Run("localhost:8080")
}