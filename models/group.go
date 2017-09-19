package models

import (
    "database/sql"
    "encoding/json"
    "errors"
)

const GroupSQLColumns = "Name, Protected, Salt"

type Group struct {
    Name string
    Protected bool
    Salt string
}

func ParseGroup(s string) (Group, error) {
    var g Group
    err := json.Unmarshal([]byte(s), &g)
    return g, err
}

func GetGroupFromRow(r *sql.Row) (Group, error) {
    var g Group
    err := r.Scan(
        &g.Name,
        &g.Protected,
        &g.Salt)
    return g, err
}


func GetGroupsFromRows(r *sql.Rows) ([]Group, error) {
    defer r.Close()
    groups := make([]Group, 0, 10)
    var err error
    for r.Next() {
        var g Group
        err := r.Scan(
            &g.Name,
            &g.Protected,
            &g.Salt)
        if err != nil {
            break
        }
        groups = append(groups, g)
    }
    return groups, err
}

func (g *Group) Validate() error {
    if g.Name == "" {
        return errors.New("Name cannot be empty")
    }
    return nil
}
