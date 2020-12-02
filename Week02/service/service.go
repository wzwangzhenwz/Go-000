package service

import (
	"error/dao"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	Id 	 uint
	Name string
}




//处理逻辑
func bizFindUserById(uid uint) (*User,error) {
	return dao.daoFindUserById(uid)
}