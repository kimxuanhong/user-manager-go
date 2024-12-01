package dao

import (
	"github.com/kimxuanhong/user-manager-go/internal/infra/entity"
	"github.com/kimxuanhong/user-manager-go/internal/infra/sql"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"log"
	"sync"
)

type UserDao interface {
	FindUserByPartnerId(ctx *app.Context, partnerId string, whenDone app.Handler[[]entity.User])
	FindAllUser(ctx *app.Context, whenDone app.Handler[[]entity.User])
}

type userDao struct {
	db *sql.Database
}

var instanceUserDao *userDao
var userDaoOnce sync.Once

func NewUserDao(db *sql.Database) UserDao {
	userDaoOnce.Do(func() {
		instanceUserDao = &userDao{db: db}
	})
	return instanceUserDao
}

func (r *userDao) FindUserByPartnerId(ctx *app.Context, partnerId string, whenDone app.Handler[[]entity.User]) {
	sql.QueryWithParams(ctx, r.db, sql.Params{Query: sql.GetUserByPartnerId, Values: []interface{}{partnerId}}, func(users []entity.User, err error) {
		if err != nil {
			log.Println("Query was error!")
			whenDone(nil, err)
			return
		}
		whenDone(users, nil)
	})
}

func (r *userDao) FindAllUser(ctx *app.Context, whenDone app.Handler[[]entity.User]) {
	sql.QueryWithoutParams(ctx, r.db, sql.FinALlUser, func(users []entity.User, err error) {
		if err != nil {
			log.Println("Query was error!")
			whenDone(nil, err)
			return
		}
		whenDone(users, nil)
	})
}
