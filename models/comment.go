package models

import (
    "database/sql"
    "encoding/json"
    "time"
)

const CommentSQLColumns = "ID, ParentPost, ParentComment, Content, EditContent, Time, Color"

var commentColMinLen = map[string]int {
    "Content" : 5,
    "EditContent" : 5,
    "Color" : 6,
}

var commentColMaxLen = map[string]int {
    "Content" : 100000000,
    "EditContent" : 100000000,
    "Color" : 6,
}

type Comment struct {
    ID int64
    ParentPost int64
    ParentComment int64
    Content string
    EditContent string
    Time int64
    Color string
}

func ParseComment(s string) (Comment, error) {
    var c Comment
    err := json.Unmarshal([]byte(s), &c)
    if c.Time == 0 {
        c.Time = time.Now().Unix()
    }
    return c, err
}

func GeFromRow(r *sql.Row) (Comment, error) {
    var c Comment
    err := r.Scan(
        &c.ID,
        &c.ParentPost,
        &c.ParentComment,
        &c.Content,
        &c.EditContent,
        &c.Time,
        &c.Color)
    return c, err
}


func GetCommentsFromRows(r *sql.Rows) ([]Comment, error) {
    defer r.Close()
    comments := make([]Comment, 0, 10)
    var err error
    for r.Next() {
        var c Comment
        err = r.Scan(
            &c.ID,
            &c.ParentPost,
            &c.ParentComment,
            &c.Content,
            &c.EditContent,
            &c.Time,
            &c.Color)
        if err != nil {
            break
        }
        comments = append(comments, c)
    }
    return comments, err
}

func (c *Comment) Validate() error {
    return ValidateRanges(c, commentColMinLen, commentColMaxLen)
}
