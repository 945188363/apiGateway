package HandlersCache

import (
	"apiGateway/DBModels"
	"apiGateway/Utils/ComponentUtil"
	"apiGateway/Utils/RedisUtil"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

// 保存Api缓存
func SaveCatchApi(api DBModels.Api) {
	bytes, err := json.Marshal(api)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("marshal api cache error", err)
		return
	}
	// 保存到redis中
	_ = RedisUtil.Set(api.ApiUrl, bytes)
}

// 删除缓存的Api
func GetCatchApi(k string) (apiCached DBModels.Api) {

	// 获取保存到redis中的handler.api bytes
	getBytes, err := redis.Bytes(RedisUtil.Get(k))
	if err != nil {
		ComponentUtil.RuntimeLog().Info("get api cache error", err)
		return DBModels.Api{}
	}
	if len(getBytes) == 0 {
		ComponentUtil.RuntimeLog().Info("api cache bytes is empty", err)
		return DBModels.Api{}
	}
	// 反序列化为api
	json.Unmarshal(getBytes, &apiCached)
	return
}

// 删除Api缓存
func DelCatchApi(k string) bool {

	// 删除redis数据
	err := RedisUtil.Del(k)
	if err != nil {
		ComponentUtil.RuntimeLog().Info("delete api cache error", err)
		return false
	}
	return true
}
