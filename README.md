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
package main

import (
	"fmt"
	"time"

	sqlxselect "github.com/cs3238-tsuzu/sqlx-selector/v2"
)

func main() {
	type User struct {
		ID             string `db:"id"`
		Name           string `db:"name"`
		OrganizationID string `db:"org_id"`
	}
	type Organization struct {
		ID   string `db:"id"`
		Name string `db:"name"`
	}

	var joined struct {
		User          *User         `db:"u"`
		Organization  *Organization `db:"org"`
		UserUpdatedAt time.Time     `db:"updated_at"`
	}

	fmt.Println(
		`SELECT ` +
			sqlxselect.New(&joined).
				SelectAs("u.updated_at", "updated_at").
				SelectStructAs("u.*", "u.*", "id", "name"). // select only id and name
				SelectStructAs("org.*", "org.*").
				String() +
			` FROM users AS u INNER JOIN organizations AS org ON u.org_id = org.id LIMIT 1`,
	)
}
```

The output will be:
```
SELECT `u`.`updated_at` AS "updated_at",`u`.`id` AS "u.id",`u`.`name` AS "u.name",`org`.`name` AS "org.name",`org`.`id` AS "org.id" FROM users AS u INNER JOIN organizations AS org ON u.org_id = org.id LIMIT 1
```