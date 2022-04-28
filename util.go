package main

import "fmt"


func findMaxPostID(posts []post) int {
    maxID := 0
    for _, post := range posts {
        if post.ID > maxID {
            maxID = post.ID
        }
    }

    return maxID
}


func findMaxCommentID(posts []post) int {
    maxID := 0
    commentIDs := collectCommentIDs(posts)
    for _, id := range commentIDs {
        if id > maxID {
            maxID = id
        }
    }

    return maxID
}


func flattenCommentIDs(currComments []comment, commentIDs []int) []int {
    for _, c := range  currComments {
        commentIDs = append(commentIDs, c.ID)
        commentIDs = append(commentIDs, flattenCommentIDs(c.Comments, commentIDs)...)
    }
    
    return commentIDs
}


func collectCommentIDs(posts []post) []int {
    var commentIDs []int
    for _, post := range posts {
        commentIDs = append(commentIDs, flattenCommentIDs(post.Comments, commentIDs)...)
    }

    return commentIDs
}


func traverseComments(comments []comment, newComment comment) {
    for i := range comments {
        if comments[i].ID == newComment.Body.RefID && newComment.Body.RefType == "comment" {
            comments[i].AddComment(newComment)
            break
        } else {
            traverseComments(comments[i].Comments, newComment)
        }
    }
}


func traverseItems(posts []post, newComment comment) {
    // TODO: return failure if match is not found
    for i := range posts {
        if posts[i].ID == newComment.Body.RefID && newComment.Body.RefType == "post" {
            posts[i].AddComment(newComment)
            fmt.Println(posts[i].Comments)
            break
        } else {
            traverseComments(posts[i].Comments, newComment)
        }
    }
}