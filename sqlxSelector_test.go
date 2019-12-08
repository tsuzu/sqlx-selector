package sqlxselect

import (
	"sort"
	"strings"
	"testing"
)

func TestSqlxSelector(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		selector := New(struct{}{})

		str, err := selector.
			Select("column-1").
			Select("column-2").
			StringWithError()

		if err != nil {
			t.Fatalf("failed to select: %v", err)
		}

		if want := "`column-1`,`column-2`"; str != want {
			t.Fatalf("want(%s) != got(%s)", want, str)
		}
	})
	t.Run("select_as", func(t *testing.T) {
		selector := New(struct{}{})

		str, err := selector.
			SelectAs("column-1", "as-1").
			SelectAs("column-2", `as-2\`).
			StringWithError()

		if err != nil {
			t.Fatalf("failed to select: %v", err)
		}

		if want := "`column-1` AS \"as-1\",`column-2` AS \"as-2\\\\\""; str != want {
			t.Fatalf("want(%s) != got(%s)", want, str)
		}
	})

	t.Run("select_struct_as_1", func(t *testing.T) {
		type dataType struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}

		selector := New(&dataType{})

		str, err := selector.
			SelectStructAs("users.*", "*").
			StringWithError()

		if err != nil {
			t.Fatalf("failed to select: %v", err)
		}

		s := strings.Split(str, ",")
		sort.Strings(s)
		str = strings.Join(s, ",")

		if want := "`users`.`id` AS \"id\",`users`.`name` AS \"name\""; str != want {
			t.Fatalf("want(%s) != got(%s)", want, str)
		}
	})
	t.Run("select_struct_as_2", func(t *testing.T) {
		type user struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}
		type group struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}

		type dataType struct {
			User  *user  `db:"user"`
			Group *group `db:"group"`
		}

		selector := New(&dataType{})

		str, err := selector.
			SelectStructAs("users.*", "user.*").
			SelectStructAs("groups.*", "group.*").
			StringWithError()

		if err != nil {
			t.Fatalf("failed to select: %v", err)
		}

		s := strings.Split(str, ",")
		sort.Strings(s)
		str = strings.Join(s, ",")

		if want := "`groups`.`id` AS \"group.id\",`groups`.`name` AS \"group.name\",`users`.`id` AS \"user.id\",`users`.`name` AS \"user.name\""; str != want {
			t.Fatalf("want(%s) != got(%s)", want, str)
		}
	})

	t.Run("select_struct", func(t *testing.T) {
		type user struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}
		type group struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}

		type dataType struct {
			User  *user  `db:"users"`
			Group *group `db:"groups"`
		}

		selector := New(&dataType{})

		str, err := selector.
			SelectStruct("users.*").
			SelectStruct("groups.*").
			StringWithError()

		if err != nil {
			t.Fatalf("failed to select: %v", err)
		}

		s := strings.Split(str, ",")
		sort.Strings(s)
		str = strings.Join(s, ",")

		if want := "`groups`.`id` AS \"groups.id\",`groups`.`name` AS \"groups.name\",`users`.`id` AS \"users.id\",`users`.`name` AS \"users.name\""; str != want {
			t.Fatalf("want(%s) != got(%s)", want, str)
		}
	})

	t.Run("select_struct_limit", func(t *testing.T) {
		type user struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}
		type group struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}

		type dataType struct {
			User  *user  `db:"users"`
			Group *group `db:"groups"`
		}

		selector := New(&dataType{})

		str, err := selector.
			SelectStruct("users.*", "id").
			SelectStruct("groups.*", "name").
			StringWithError()

		if err != nil {
			t.Fatalf("failed to select: %v", err)
		}

		s := strings.Split(str, ",")
		sort.Strings(s)
		str = strings.Join(s, ",")

		if want := "`groups`.`name` AS \"groups.name\",`users`.`id` AS \"users.id\""; str != want {
			t.Fatalf("want(%s) != got(%s)", want, str)
		}
	})

	t.Run("new_for_non_struct_value", func(t *testing.T) {
		selector := New("")

		_, err := selector.StringWithError()

		if err == nil {
			t.Error("New for string should return an error")
		}
	})
}
