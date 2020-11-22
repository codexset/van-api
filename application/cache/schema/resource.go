package schema

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"taste-api/application/model"
)

type Resource struct {
	key string
	dep Dependency
}

func NewResource(dep Dependency) *Resource {
	c := new(Resource)
	c.key = "system:resource"
	c.dep = dep
	return c
}

func (c *Resource) Clear() {
	c.dep.Redis.Del(context.Background(), c.key)
}

func (c *Resource) Get() (result []map[string]interface{}, err error) {
	ctx := context.Background()
	var exists int64
	exists, err = c.dep.Redis.Exists(ctx, c.key).Result()
	if err != nil {
		return
	}
	if exists == 0 {
		var resourceLists []map[string]interface{}
		c.dep.Db.Model(&model.Resource{}).
			Select([]string{"`key`", "parent", "name", "nav", "router", "policy", "icon"}).
			Where("status = ?", 1).
			Order("sort desc").
			Scan(&resourceLists)
		var buf []byte
		buf, err = jsoniter.Marshal(resourceLists)
		if err != nil {
			return
		}
		err = c.dep.Redis.Set(ctx, c.key, string(buf), 0).Err()
		if err != nil {
			return
		}
	}
	var raw []byte
	raw, err = c.dep.Redis.Get(ctx, c.key).Bytes()
	if err != nil {
		return
	}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return
	}
	return
}