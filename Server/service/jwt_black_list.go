package service

import (
	"context"
	"github.com/vmihailenco/msgpack"
	"go_admin/Server/global"
	"go_admin/Server/model"
)

func JsonInBlacklist(blacklist model.JwtBlacklist) error {
	return global.GVA_DB.Create(&blacklist).Error
}
func IsBlackList(jwt string, jwtList model.JwtBlacklist) bool {

	isNotFound := global.GVA_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()

	return !isNotFound
}

var ctx = context.Background()

func GetRedisJWT(userName string) (err error, redisJWT string) {
	//fmt.Println("redis")
	//fmt.Println(global.GVA_REDIS)
	redisJWT, err = global.GVA_REDIS.Get(ctx, userName).Result()
	return err, redisJWT
}

func SetRedisJWT(jwtList model.JwtBlacklist, userName string) error {

	return global.GVA_REDIS.Set(ctx, userName, &jwtList, 60*60*1).Err()
}

type something struct {
	ID   int
	Name string
}

func (s *something) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}

func (s *something) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
