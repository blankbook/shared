package models

import (
    "database/sql"
    "encoding/json"
)

const PostSQLColumns = "Rank, ID, Score, Title, EditTitle, Content, EditContent, ContentType, GroupName, Time, Color"

var pColMin = map[string]int64 {
    "Title" : 5,
    "EditTitle": 5,
    "Color": 6,
}

var pColMax = map[string]int64 {
    "Title" : 300,
    "EditTitle": 300,
    "Content": 100000000,
    "EditContent": 100000000,
    "ContentType":100,
    "GroupName":100,
    "Color": 6,
}

var pRequiredCols = map[string]bool {
    "ID" : false,
    "Score" : false,
    "Title" : true,
    "EditTitle" : false,
    "Content" : true,
    "EditContent" : false,
    "ContentType" : true,
    "GroupName" : true,
    "Time" : true,
    "Color" : true,
}

type Post struct {
    Rank int64
    ID int64
    Score int32
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
    return p, err
}

func GetPostFromRow(r *sql.Row) (Post, error) {
    var p Post
    err := r.Scan(
        &p.Rank,
        &p.ID,
        &p.Score,
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
            &p.Rank,
            &p.ID,
            &p.Score,
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
    return ValidateRanges(p, pRequiredCols, pColMin, pColMax)
}
