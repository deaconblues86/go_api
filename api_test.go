package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/kinbiko/jsonassert"
)


func TestGetPostsRoute(t *testing.T) {
    router := setupRouter()

    expected := `
    [
        {
            "id": 1,
            "Body": {
                "title": "Test Post",
                "author": "Brian H",
                "content": "I'm a test",
                "timestamp": "2022-04-27T5:49Z"
            },
            "Comments": null
        },
        {
            "id": 2,
            "Body": {
                "title": "Another Post",
                "author": "Brian H",
                "content": "I'm a test",
                "timestamp": "2022-04-27T5:51Z"
            },
            "Comments": null
        }
    ]`

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/posts", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestGetSinglePostsRoute(t *testing.T) {
    router := setupRouter()

    expected := `
    {
        "id": 2,
        "Body": {
            "title": "Another Post",
            "author": "Brian H",
            "content": "I'm a test",
            "timestamp": "2022-04-27T5:51Z"
        },
        "Comments": null
    }`

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/posts/2", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestPostPostsRoute(t *testing.T) {
    router := setupRouter()

    new_post := map[string]interface{}{
        "title": "A New Post",
        "author": "Brian H",
        "content": "I'm a test",
        "timestamp": "2022-04-27T6:11Z",
    }

    body, _ := json.Marshal(new_post)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/posts", bytes.NewReader(body))
    router.ServeHTTP(w, req)

    assert.Equal(t, 201, w.Code)
    expected := `
    {
        "id": 3,
        "Body": {
            "title": "A New Post",
            "author": "Brian H",
            "content": "I'm a test",
            "timestamp": "2022-04-27T6:11Z"
        },
        "Comments": null
    }`
    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestDeletePostsRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/posts/3", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

}


func TestPostCommentsRoute(t *testing.T) {
    router := setupRouter()

    new_post := map[string]interface{}{
        "author": "Brian H",
        "content": "I'm a comment",
        "timestamp": "2022-04-27T6:11Z",
        "ref_id": 2,
        "ref_type": "post",
    }

    body, _ := json.Marshal(new_post)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/comments", bytes.NewReader(body))
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    get_req, _ := http.NewRequest("GET", "/posts", nil)
    router.ServeHTTP(w, get_req)

    assert.Equal(t, 200, w.Code)
    expected := `
    [
        {
            "id": 1,
            "Body": {
                "title": "Test Post",
                "author": "Brian H",
                "content": "I'm a test",
                "timestamp": "2022-04-27T5:49Z"
            },
            "Comments": null
        },
        {
            "id": 2,
            "Body": {
                "title": "Another Post",
                "author": "Brian H",
                "content": "I'm a test",
                "timestamp": "2022-04-27T5:51Z"
            },
            "Comments": [
                {
                    "id": 1,
                    "Body": {
                        "author": "Brian H",
                        "content": "I'm a comment",
                        "timestamp": "2022-04-27T6:11Z",
                        "ref_id": 2,
                        "ref_type": "post"
                    },
                    "Comments": null
                }
            ]
        }
    ]`
    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestPostNestedCommentsRoute(t *testing.T) {
    router := setupRouter()

    new_post := map[string]interface{}{
        "author": "Brian H",
        "content": "I'm a Nested comment",
        "timestamp": "2022-04-27T6:11Z",
        "ref_id": 1,
        "ref_type": "comment",
    }

    body, _ := json.Marshal(new_post)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/comments", bytes.NewReader(body))
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    get_req, _ := http.NewRequest("GET", "/posts", nil)
    router.ServeHTTP(w, get_req)

    assert.Equal(t, 200, w.Code)
    expected := `
    [
        {
            "id": 1,
            "Body": {
                "title": "Test Post",
                "author": "Brian H",
                "content": "I'm a test",
                "timestamp": "2022-04-27T5:49Z"
            },
            "Comments": null
        },
        {
            "id": 2,
            "Body": {
                "title": "Another Post",
                "author": "Brian H",
                "content": "I'm a test",
                "timestamp": "2022-04-27T5:51Z"
            },
            "Comments": [
                {
                    "id": 1,
                    "Body": {
                        "author": "Brian H",
                        "content": "I'm a comment",
                        "timestamp": "2022-04-27T6:11Z",
                        "ref_id": 2,
                        "ref_type": "post"
                    },
                    "Comments": [
                        {
                            "id": 2,
                            "Body": {
                                "author": "Brian H",
                                "content": "I'm a Nested comment",
                                "timestamp": "2022-04-27T6:11Z",
                                "ref_id": 1,
                                "ref_type": "comment"
                            },
                            "Comments": null
                        }
                    ]
                }
            ]
        }
    ]`
    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestPutPostRoute(t *testing.T) {
    router := setupRouter()

    new_post := map[string]interface{}{
        "title": "Another Post Modified",
        "author": "Brian H - Mod",
        "content": "I'm a test modificiation",
        "timestamp": "2022-04-27T5:51Z",
    }

    body, _ := json.Marshal(new_post)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/posts/2", bytes.NewReader(body))
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    expected := `
    {
        "id": 2,
        "Body": {
            "title": "Another Post Modified",
            "author": "Brian H - Mod",
            "content": "I'm a test modificiation",
            "timestamp": "2022-04-27T5:51Z"
        },
        "Comments": [
            {
                "id": 1,
                "Body": {
                    "author": "Brian H",
                    "content": "I'm a comment",
                    "timestamp": "2022-04-27T6:11Z",
                    "ref_id": 2,
                    "ref_type": "post"
                },
                "Comments": [
                    {
                        "id": 2,
                        "Body": {
                            "author": "Brian H",
                            "content": "I'm a Nested comment",
                            "timestamp": "2022-04-27T6:11Z",
                            "ref_id": 1,
                            "ref_type": "comment"
                        },
                        "Comments": null
                    }
                ]
            }
        ]
    }`
    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestPostNestedCommentOnNestedCommentRoute(t *testing.T) {
    router := setupRouter()

    new_post := map[string]interface{}{
        "author": "Brian H",
        "content": "I'm a Nested comment on a nested comment",
        "timestamp": "2022-04-27T6:11Z",
        "ref_id": 2,
        "ref_type": "comment",
    }

    body, _ := json.Marshal(new_post)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/comments", bytes.NewReader(body))
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    expected := `
    {
        "id": 2,
        "Body": {
            "title": "Another Post Modified",
            "author": "Brian H - Mod",
            "content": "I'm a test modificiation",
            "timestamp": "2022-04-27T5:51Z"
        },
        "Comments": [
            {
                "id": 1,
                "Body": {
                    "author": "Brian H",
                    "content": "I'm a comment",
                    "timestamp": "2022-04-27T6:11Z",
                    "ref_id": 2,
                    "ref_type": "post"
                },
                "Comments": [
                    {
                        "id": 2,
                        "Body": {
                            "author": "Brian H",
                            "content": "I'm a Nested comment",
                            "timestamp": "2022-04-27T6:11Z",
                            "ref_id": 1,
                            "ref_type": "comment"
                        },
                        "Comments": [
                            {
                                "id": 3,
                                "Body": {
                                    "author": "Brian H",
                                    "content": "I'm a Nested comment on a nested comment",
                                    "timestamp": "2022-04-27T6:11Z",
                                    "ref_id": 2,
                                    "ref_type": "comment"
                                },
                                "Comments": null
                            }
                        ]
                    }
                ]
            }
        ]
    }`

    get_req, _ := http.NewRequest("GET", "/posts/2", nil)
    router.ServeHTTP(w, get_req)

    assert.Equal(t, 200, w.Code)

    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestGetNestedCommentRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/comments/2", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    expected := `
    {
        "id": 2,
        "Body": {
            "author": "Brian H",
            "content": "I'm a Nested comment",
            "timestamp": "2022-04-27T6:11Z",
            "ref_id": 1,
            "ref_type": "comment"
        },
        "Comments": [
            {
                "id": 3,
                "Body": {
                    "author": "Brian H",
                    "content": "I'm a Nested comment on a nested comment",
                    "timestamp": "2022-04-27T6:11Z",
                    "ref_id": 2,
                    "ref_type": "comment"
                },
                "Comments": null
            }
        ]
    }`
    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}

func TestPutNestedCommentOnNestedCommentRoute(t *testing.T) {
    router := setupRouter()

    new_post := map[string]interface{}{
        "author": "Brian H - Modified",
        "content": "I'm a Modified Nested comment on a nested comment",
        "timestamp": "2022-04-27T6:11Z",
    }

    body, _ := json.Marshal(new_post)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("PUT", "/comments/3", bytes.NewReader(body))
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    putExpected := `
    {
        "id": 3,
        "Body": {
            "author": "Brian H - Modified",
            "content": "I'm a Modified Nested comment on a nested comment",
            "timestamp": "2022-04-27T6:11Z",
            "ref_id": 2,
            "ref_type": "comment"
        },
        "Comments": null
    }`

    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), putExpected)

    // Requesting Post #2 to ensure changes were actually propagated
    wFinal := httptest.NewRecorder()
    get_req, _ := http.NewRequest("GET", "/posts/2", nil)
    router.ServeHTTP(wFinal, get_req)

    assert.Equal(t, 200, wFinal.Code)

    expected := `
    {
        "id": 2,
        "Body": {
            "title": "Another Post Modified",
            "author": "Brian H - Mod",
            "content": "I'm a test modificiation",
            "timestamp": "2022-04-27T5:51Z"
        },
        "Comments": [
            {
                "id": 1,
                "Body": {
                    "author": "Brian H",
                    "content": "I'm a comment",
                    "timestamp": "2022-04-27T6:11Z",
                    "ref_id": 2,
                    "ref_type": "post"
                },
                "Comments": [
                    {
                        "id": 2,
                        "Body": {
                            "author": "Brian H",
                            "content": "I'm a Nested comment",
                            "timestamp": "2022-04-27T6:11Z",
                            "ref_id": 1,
                            "ref_type": "comment"
                        },
                        "Comments": [
                            {
                                "id": 3,
                                "Body": {
                                    "author": "Brian H - Modified",
                                    "content": "I'm a Modified Nested comment on a nested comment",
                                    "timestamp": "2022-04-27T6:11Z",
                                    "ref_id": 2,
                                    "ref_type": "comment"
                                },
                                "Comments": null
                            }
                        ]
                    }
                ]
            }
        ]
    }`

    ja.Assertf(wFinal.Body.String(), expected)
}

func TestDeleteNestedCommenRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/comments/2", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)

    expected := `
    {
        "id": 2,
        "Body": {
            "title": "Another Post Modified",
            "author": "Brian H - Mod",
            "content": "I'm a test modificiation",
            "timestamp": "2022-04-27T5:51Z"
        },
        "Comments": [
            {
                "id": 1,
                "Body": {
                    "author": "Brian H",
                    "content": "I'm a comment",
                    "timestamp": "2022-04-27T6:11Z",
                    "ref_id": 2,
                    "ref_type": "post"
                },
                "Comments": [
                ]
            }
        ]
    }`

    get_req, _ := http.NewRequest("GET", "/posts/2", nil)
    router.ServeHTTP(w, get_req)

    assert.Equal(t, 200, w.Code)

    ja := jsonassert.New(t)
    ja.Assertf(w.Body.String(), expected)
}