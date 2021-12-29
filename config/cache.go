package config

type RedisCache struct {
	Host    string
	Db      int64
	Expires int
	Port    string
}
