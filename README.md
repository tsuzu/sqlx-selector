<!--
 Copyright (c) 2019 Tsuzu
 
 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

# sqlx-selector
- SELECT columns helper library for sqlx
- Maybe useful for queries using JOIN

## Installation
```
go get github.com/cs3238-tsuzu/sqlx-selector/v2
```

## Usage
with sqlx

```go
type User struct {
    ID string `db:"id"`
    Name string `db:"name"`
    OrganizationID string `db:"org_id"`
}
type Organization struct {
    ID string `db:"id"`
    Name string `db:"name"`
}

var j struct {
    User *User `db:"u"`
    Organization *Organization `db:"org"`
    UserUpdatedAt time.Time `db:"updated_at"`
}

db.QueryRowx(
    `SELECT` + 
        sqlxselect.New(j).
            SelectAs("u.updated_at", "updated_at").
            SelectStructAs("u.*", "user.*", "id", "name"). // select only id and name
            SelectStructAs("org.*", "org.*").
            String() +
        `FROM users AS u INNER JOIN organizations AS org ON u.org_id = org.id LIMIT 1`,
).StructScan(&j)
```
