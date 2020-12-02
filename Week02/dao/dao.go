package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	Id 	 uint
	Name string
}

//执行数据库查询的函数
func getUserFromDb() (*User,error) {
	//TODO
	//省略数据库逻辑，直接返回sql.ErrNoRows
	return &User{}, sql.ErrNoRows
}


//dao层处理逻辑
func daoFindUserById(uid uint) (*User, error) {
	user,err := getUserFromDb()
	if err != nil{
		return user,errors.Wrap(err, fmt.Sprintf("dao error: find user by id=%+v", uid))
	}
	return user, nil
}
