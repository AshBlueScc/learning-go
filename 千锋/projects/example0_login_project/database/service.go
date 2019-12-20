package database

import (
	"github.com/go-xorm/xorm"
	"time"
)

type Service interface {
	GetByAdminNameAndPassword(username, password string)(*User, bool)

}

var service Service

func GetService() Service {
	if service == nil {
		service = &dbService{
			db:      GetDB(),
			timeout: time.Duration(dbTimeout) *time.Second,
		}
	}
	return service
}

type dbService struct {
	db	*xorm.Engine
	timeout	time.Duration
}

func (s *dbService) GetByAdminNameAndPassword(username, password string)(*User, bool){

	var user User

	has, err := s.db.Where("user_name = ? and password = ?", username, password).Get(&user)
	if has && err == nil{
		return &user, has
	}

	return nil, false
}

