package dao

import (
	"fmt"
	"gitee.com/qianxunke/book-ticket-common/plugins/redis"
	ticketProto "gitee.com/qianxunke/book-ticket-common/proto/ticket"
	r "github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	dao              *ticketDaoIml
	m                sync.Mutex
	redisClient      *r.Client
	tokenExpiredDate = 5 * time.Second //5秒更新一次
)

type ticketDaoIml struct {
}

type TicketDao interface {
	FindById(secretStr string) (product *ticketProto.Train, err error)

	Insert(product []*ticketProto.Train) (err error)

	SimpleQuery(findFrom string, findTo string, trainDate string, purposeCodes string) (value string, err error)

	Delete(ids []string) (err error)

	Update(product *ticketProto.Train) (err error)
}

func GetTicketDao() (TicketDao, error) {
	if dao == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return dao, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if dao != nil {
		return
	}
	dao = &ticketDaoIml{}
	redisClient = redis.Redis()
}
