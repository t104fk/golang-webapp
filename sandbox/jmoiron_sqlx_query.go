package sandbox

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// selectStatement is select
const selectInStatement = `
SELECT
  id,
  name,
  created_at
WHERE
  id in (:ids)
`

func namedExec() {
	arg := map[string]interface{}{
		"ids": []int{1, 2, 3, 4},
	}
	query, args, err := sqlx.Named(selectInStatement, arg)
	if err != nil {
		panic(err)
	}
	/*
		SELECT
		  id,
		  name,
		  created_at
		WHERE
		  id in (?)
	*/
	fmt.Println(query)
	/*
		[[1 2 3 4]]
	*/
	fmt.Println(args)
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		panic(err)
	}
	/*
		SELECT
		  id,
		  name,
		  created_at
		WHERE
		  id in (?, ?, ?, ?)
	*/
	fmt.Println(query)
	/*
		[1 2 3 4]
	*/
	fmt.Println(args)
}
