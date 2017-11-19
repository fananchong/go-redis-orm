package go_redis_orm

var g_redismgr *RedisMgr

type RedisMgr struct {
	dbs map[string]*RedisClient
}

func NewRedisMgr() *RedisMgr {
	return &RedisMgr{dbs: make(map[string]*RedisClient)}
}

func (this *RedisMgr) Create(dbName string, addrs []string) {
	db := NewRedisClient(dbName, addrs)
	this.dbs[dbName] = db
}

func (this *RedisMgr) Get(dbName string) *RedisClient {
	if db, ok := this.dbs[dbName]; ok {
		return db
	}
	return nil
}

func init() {
	g_redismgr = NewRedisMgr()
}
