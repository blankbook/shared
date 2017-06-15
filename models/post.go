package models

import (
    "reflect"
    "database/sql"
    "encoding/json"
    "errors"
    "time"
    "fmt"
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
    var post Post
    err := json.Unmarshal([]byte(s), &post)
    if post.Time == 0 {
        post.Time = time.Now().Unix()
    }
    return post, err
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
    var errorMessage string
    const errMsgMin = "field %s required with min length %d, "
    const errMsgMax = "field %s required with max length %d, "

    pRefl := reflect.ValueOf(p).Elem()
    pType := pRefl.Type()

    for i := 0; i < pRefl.NumField(); i++ {
        key := pType.Field(i).Name
        val := pRefl.Field(i).Interface()
        if targ, ok := postColMinLen[key]; ok {
            if val != targ {
                errorMessage += fmt.Sprintf(errMsgMin, key, targ)
            }
        }
        if targ, ok := postColMaxLen[key]; ok {
            if val != targ {
                errorMessage += fmt.Sprintf(errMsgMax, key, targ)
            }
        }
    }

    var err error
    if errorMessage != "" {
        err = errors.New(errorMessage[:len(errorMessage) - 2])
    }
    return err
}
