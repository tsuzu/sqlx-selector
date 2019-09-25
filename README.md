<!--
 Copyright (c) 2019 Tsuzu
 
 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

# sqlx-selector
- SELECT columns helper library for sqlx
- Maybe useful for queries using JOIN

## Usage
with sqlx

```go
type User struct {
    ID string `db:"id"`
    Name string `db:"name"`
    GroupID string `db:"group_id"`
}
type Group struct {
    ID string `db:"id"`
    Name string `db:"name"`
}

type join struct {
    User *User `db:"user"`
    Group *Group `db:"group"`
    UserUpdatedAt time.Time `db:"user_updated_at"`
}

var j join
db.QueryRowx(
    `SELECT` + 
        sqlxselect.New(j).
            SelectAs("users.updated_at", "user_updated_at").
            SelectStructAs("users.*", "user.*", "id". "name"). // select only id and name
            SelectStructAs("groups.*", "group.*").
            String() +
    `FROM users INNER JOIN groups ON users.group_id = groups.id LIMIT 1`
).StructScan(&j)
```