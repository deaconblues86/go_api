# Go API
Simple Blog API

Includes APIs to:
- Fetch all blog posts, their comments, and their comment's comments
- All CRUD operations on blog posts
- All CRUD operations on comments to a blog post or comments to a comment

## Installation
- Install Go 1.18
- Clone repository into desired location
- Within the newly created folder, execute `go run .` to start server
  - Note:  All data is currently stored in memory
- To execute unit tests, run `go test` from the root of the repository

## Get Posts

```
GET localhost:8080/posts/
```

Get All Posts

## Get Post

```
GET localhost:8080/posts/2
```

Get Specific Post by ID

## Post Post

```
POST localhost:8080/posts
```

Create a new Post

### Request

> 
> **Body**
> 
> ```
> {
>     "title": "A New Post",
>     "author": "Brian H",
>     "content": "I'm a test",
>     "timestamp": "2022-04-27T6:11Z"
> }
> ```
> 

## Add Comment on Post

```
POST localhost:8080/comments
```

Adds a comment to specified Post

### Request

> 
> **Body**
> 
> ```
> {
>     "author": "Brian H",
>     "content": "I'm a comment",
>     "timestamp": "2022-04-27T6:11Z",
>     "ref_id": 2,
>     "ref_type": "post"
> }
> ```
> 

## Add Nested Comment on Comment

```
POST localhost:8080/comments
```

Adds a nested comment on a comment

### Request

> 
> **Body**
> 
> ```
> {
>     "author": "Brian H",
>     "content": "I'm a Nested comment",
>     "timestamp": "2022-04-27T6:11Z",
>     "ref_id": 1,
>     "ref_type": "comment"
> }
> ```
> 

## Add Nested Comment on Comment on Comment

```
POST localhost:8080/comments
```

Adds a nested comment on nested comment

### Request

> 
> **Body**
> 
> ```
> {
>     "author": "Brian H",
>     "content": "I'm a Nested comment",
>     "timestamp": "2022-04-27T6:11Z",
>     "ref_id": 2,
>     "ref_type": "comment"
> }
> ```
> 

## Update Post

```
PUT localhost:8080/posts/2
```

Updates a Post

### Request

> 
> **Body**
> 
> ```
> {
>     "title": "Another Post Modified",
>     "author": "Brian H - Mod",
>     "content": "I'm a test modificiation",
>     "timestamp": "2022-04-27T5:51Z"
> }
> ```
> 

## Delete Post

```
DELETE localhost:8080/posts/2
```

Deletes a Post

### Request

> 

## Get Nested Comment

```
GET localhost:8080/comments/2
```

Gets the specified comment and it's comments

## Update Nested Comment

```
PUT localhost:8080/comments/2
```

Updates a comment

### Request

> 
> **Body**
> 
> ```
> {
>     "author": "Brian H - Modified",
>     "content": "I'm a modified nested comment",
>     "timestamp": "2022-04-27T5:51Z"
> }
> ```
> 

## Delete Nested Comment

```
DELETE localhost:8080/comments/3
```

Deletes a Comment

### Request

> 
