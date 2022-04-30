package controllers

import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
)


func GetComment(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    _, reqComment := findComment(posts, id)
    if reqComment.ID == 0 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "comment not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, reqComment)
}


func PutComment(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    _, reqComment := findComment(posts, id)
    if reqComment.ID == 0 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "comment not found"})
        return
    }

    if err := c.BindJSON(&reqComment.Body); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
        return
    }
    c.IndentedJSON(http.StatusOK, reqComment)
    return
}


func DeleteComment(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    delIDX, reqComment := findComment(posts, id)
    if reqComment.ID == 0 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "comment not found"})
        return
    }

    if reqComment.Body.RefType == "post" {
        for i, _ := range posts {
            if posts[i].ID == reqComment.Body.RefID {
                posts[i].Comments = removeComment(posts[i].Comments, delIDX)
                return
            }
        }
    }

    _, parentComment := findComment(posts, reqComment.Body.RefID)
    parentComment.Comments = removeComment(parentComment.Comments, delIDX)
}


func PostComments(c *gin.Context) {
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