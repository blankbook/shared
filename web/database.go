package web

import (
    "fmt"
    "database/sql"
    _ "github.com/denisenkom/go-mssqldb"
)

func GetMSSqlDatabase(dbUsername, dbPassword, dbServer, dbName string) (*sql.DB, error) {
    db, err := sql.Open("mssql",
                        fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
                                    dbUsername,
                                    dbPassword,
                                    dbServer,
                                    dbName))
    return db, err
}
