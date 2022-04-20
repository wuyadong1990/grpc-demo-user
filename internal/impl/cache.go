package impl

import (
	"context"

	"github.com/wuyadong1990/grpc-demo-user/cinit"
	"github.com/wuyadong1990/grpc-demo-user/internal/utils"

	"github.com/xiaomeng79/go-log"
)

const (
	CacheIDPrefix = "ucid"
)

func CacheGet(ctx context.Context, id int64) (map[string]string, error) {
	k := getIDKey(CacheIDPrefix, id)
	// 获取全部
	r, err := cinit.RedisCli.HGetAll(k).Result()
	if err != nil {
		log.Info(err.Error(), ctx)
	}

	return r, err
}

func CacheSet(ctx context.Context, id int64, m *User) {
	_m := utils.Struct2Map(*m)
	k := getIDKey(CacheIDPrefix, id)
	log.Debugf("[CacheSet] with key: %s", k)
	err := cinit.RedisCli.HMSet(k, _m).Err()
	if err != nil {
		log.Error(err.Error(), ctx)
		return
	}
	setKeyExpire(ctx, k)
}

func CacheDel(ctx context.Context, id int64) {
	k := getIDKey(CacheIDPrefix, id)
	err := cinit.RedisCli.Del(k).Err()
	if err != nil {
		log.Info(err.Error(), ctx)
	}
}
