package models

import (
    "encoding/json"
    "errors"
    "time"
    "strconv"
)

const MIN_TITLE_LENGTH = 5

type Post struct {
    Title string // limit - 300 characters
    Content string // limit - 2^31 characters
    ContentType string // limit - 100 characters
    GroupName string // limit - 100 characters
    Time int64 // format: 1900-04-11 00:00:00
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

    if len(p.Title) < MIN_TITLE_LENGTH {
        errorMessage += "field 'Title' required with length at least " +
                        strconv.Itoa(MIN_TITLE_LENGTH) + ", "
    }
    if p.Content == "" {
        errorMessage += "field 'Content' required, "
    }
    if p.ContentType == "" {
        errorMessage += "field 'ContentType' required, "
    }
    if p.Group == "" {
        errorMessage += "field 'Group' required, "
    }

    var err error
    if errorMessage != "" {
        err = errors.New(errorMessage[:len(errorMessage) - 2])
    }
    return err
}
