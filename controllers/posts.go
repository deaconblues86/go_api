package controllers

import (
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
)


func GetPosts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, posts)
}


func GetPost(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    for _, p := range posts {
        if p.ID == id {
            c.IndentedJSON(http.StatusOK, p)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}


func PutPost(c *gin.Context) {
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


func DeletePost(c *gin.Context) {
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


func PostPosts(c *gin.Context) {
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