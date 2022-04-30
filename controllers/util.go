package controllers

import "fmt"

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


var posts = []post {
    {ID: 1, Body: postBody{Title: "Test Post", Author: "Brian H", Content: "I'm a test", Timestamp: "2022-04-27T5:49Z"}},
    {ID: 2, Body: postBody{Title: "Another Post", Author: "Brian H", Content: "I'm a test", Timestamp: "2022-04-27T5:51Z"}},
}


func removePost(orig []post, index int) []post {
    ret := make([]post, 0)
    origLength := len(orig)

    if index < 0 || index > origLength - 1 {
        return nil
    }
    ret = append(ret, orig[:index]...)

    if origLength < index + 1 {
        return ret
    }

    return append(ret, orig[index+1:]...)
}


func removeComment(orig []comment, index int) []comment {
    ret := make([]comment, 0)
    origLength := len(orig)

    if index < 0 || index > origLength - 1 {
        return nil
    }
    ret = append(ret, orig[:index]...)

    if origLength < index + 1 {
        return ret
    }

    return append(ret, orig[index+1:]...)
}



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


func checkNestedComments(comments []comment, id int) (int, *comment) {
    fmt.Println("Checking comments of comment")
    for i, _ := range comments {
        if comments[i].ID == id {
            return i, &comments[i]
        }
        nIDX, nComment := checkNestedComments(comments[i].Comments, id)
        if nComment.ID != 0 {
            return nIDX, nComment
        }
    }
    return 0, &comment{}
}


func findComment(posts []post, id int) (int, *comment) {
    fmt.Println("Checking posts for comment")
    for pIdx, _ := range posts {
        for cIdx, _ := range posts[pIdx].Comments {
            if posts[pIdx].Comments[cIdx].ID == id {
                return cIdx, &posts[pIdx].Comments[cIdx]
            }
            nIDX, nComment := checkNestedComments(posts[pIdx].Comments[cIdx].Comments, id)
            if nComment.ID != 0 {
                return nIDX, nComment
            }
        }
    }
    return 0, &comment{}
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
            break
        } else {
            traverseComments(posts[i].Comments, newComment)
        }
    }
}