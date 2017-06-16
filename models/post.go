package models

import (
    "database/sql"
    "encoding/json"
    "time"
)

const PostSQLColumns = "ID, Title, EditTitle, Content, EditContent, ContentType, GroupName, Time, Color"

var postColMinLen = map[string]int {
    "Title" : 5,
    "EditTitle": 5,
    "Color": 6,
}

var postColMaxLen = map[string]int {
    "Title" : 300,
    "EditTitle": 300,
    "Content": 100000000,
    "EditContent": 100000000,
    "ContentType":100,
    "GroupName":100,
    "Color": 6,
}

type Post struct {
    ID int64
    Title string
    EditTitle string
    Content string
    EditContent string
    ContentType string
    GroupName string
    Time int64
    Color string
}

func ParsePost(s string) (Post, error) {
    var p Post
    err := json.Unmarshal([]byte(s), &p)
    if p.Time == 0 {
        p.Time = time.Now().Unix()
    }
    return p, err
}

func GetPostFromRow(r *sql.Row) (Post, error) {
    var p Post
    err := r.Scan(
        &p.ID,
        &p.Title,
        &p.EditTitle,
        &p.Content,
        &p.EditContent,
        &p.ContentType,
        &p.GroupName,
        &p.Time,
        &p.Color)
    return p, err
}


func GetPostsFromRows(r *sql.Rows) ([]Post, error) {
    defer r.Close()
    posts := make([]Post, 0, 10)
    var err error
    for r.Next() {
        var p Post
        err = r.Scan(
            &p.ID,
            &p.Title,
            &p.EditTitle,
            &p.Content,
            &p.EditContent,
            &p.ContentType,
            &p.GroupName,
            &p.Time,
            &p.Color)
        if err != nil {
            break
        }
        posts = append(posts, p)
    }
    return posts, err
}

func (p *Post) Validate() error {
    return ValidateRanges(p, postColMinLen, postColMaxLen)
}
