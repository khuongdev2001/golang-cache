package cache

import (
	"time"
)

type Cache struct {
	expiration      int64
	purgeExpiration int64
	data            map[string]Data
}

type Data struct {
	value     any
	expiredAt int64
}

func (cache *Cache) Init(defaultExpiration float64, purgeExpiration float64) {
	cache.expiration = int64(defaultExpiration)
	cache.purgeExpiration = int64(purgeExpiration)
	cache.data = make(map[string]Data)
}

func (cache *Cache) Get(key string) interface{} {
	if cache.data[key].expiredAt > time.Now().Unix() {
		return cache.data[key].value
	}
	cache.Delete(key)
	return nil
}

func (cache *Cache) Set(key string, value any) {
	cache.data[key] = Data{value, time.Now().Unix() + cache.expiration}
}

func (cache *Cache) Delete(key string) {
	delete(cache.data, key)
}

//func main() {
//var cache Cache
//cache.Init(2*time.Minute.Seconds(), 2*time.Minute.Seconds())
//cache.Set("Dev", "Ahihi")
//fmt.Println(cache.data)
//fmt.Println(cache.Get("Dev"))
//}
