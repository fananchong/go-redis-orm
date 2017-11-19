package go_redis_orm

type RD_Base interface {
	Key() string
	Value() ([]byte, error)
	Load(dbName string) error
	Save(dbName string) error
	Delete(dbName string) error
}

func CreateDB(dbName string, addrs []string) {
	g_redismgr.Create(dbName, addrs)
}

func GetDB(dbName string) *RedisClient {
	return g_redismgr.Get(dbName)
}
