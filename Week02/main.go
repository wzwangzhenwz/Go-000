package main

import (
	"database/sql"
	"error/service"
	"fmt"
	"github.com/pkg/errors"
)


type User struct {
	Id 	 uint
	Name string
}

func main() {
	user, err := service.bizFindUserById(1234567)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("user not exists: %+v\n", err)
			return
		}

		fmt.Printf("query user failed: %+v\n", err)
		return
	}

	fmt.Printf("query user success: user=%+v\n", user)

}