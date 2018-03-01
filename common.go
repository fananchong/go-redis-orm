package go_redis_orm

type IClient interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	Del(key string) error
}

type NewRedisType func(dbName string, addrs []string, password string, dbindex int) (IClient, error)

func SetNewRedisHandler(handler NewRedisType) {
	g_redismgr.SetNewRedisHandler(handler)
}

func CreateDB(dbName string, addrs []string, password string, dbindex int) {
	g_redismgr.Create(dbName, addrs, password, dbindex)
}

func GetDB(dbName string) IClient {
	return g_redismgr.Get(dbName)
}
