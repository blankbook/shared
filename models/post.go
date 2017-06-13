package models

import (
    "encoding/json"
    "errors"
    "time"
    "fmt"
)

const minTitleLength = 5
const maxTitleLength = 300
const maxContentLength = 2147483647
const maxContentTypeLength = 100
const maxGroupNameLength = 100

type Post struct {
    Title string
    Content string
    ContentType string
    GroupName string
    Time int64
}

func ParsePost(s string) (Post, error) {
    var post Post
    err := json.Unmarshal([]byte(s), &post)
    if post.Time == 0 {
        post.Time = time.Now().Unix()
    }
    return post, err
}

func (p *Post) Validate() error {
    var errorMessage string
    const errMsgTemplate = "field %s required with length from %d to %d"

    if len(p.Title) < minTitleLength || len(p.Title) > maxTitleLength {
        errorMessage += fmt.Sprintf(errMsgTemplate, "Title", minTitleLength,
                                    maxTitleLength)
    }
    if p.Content == "" || len(p.Content) > maxContentLength {
        errorMessage += fmt.Sprintf(errMsgTemplate, "Content", 1,
                                    maxContentLength)
    }
    if p.ContentType == "" || len(p.ContentType) > maxContentTypeLength {
        errorMessage += fmt.Sprintf(errMsgTemplate, "ContentType", 1,
                                    maxContentTypeLength)
    }
    if p.GroupName == "" || len(p.GroupName) > maxGroupNameLength {
        errorMessage += fmt.Sprintf(errMsgTemplate, "GroupName", 1,
                                    maxGroupNameLength)
    }

    var err error
    if errorMessage != "" {
        err = errors.New(errorMessage[:len(errorMessage) - 2])
    }
    return err
}
