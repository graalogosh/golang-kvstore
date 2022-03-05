package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"kvstore/kvstore"
)

var kv kvstore.KVStorage

func main() {
	kv = kvstore.NewKVStorage()

	r := gin.Default()
	r.POST("/:key", saveValue)
	r.GET("/:key", getValue)
	r.DELETE("/:key", deleteValue)
	r.Run()
}

func saveValue(c *gin.Context) {
	key := c.Param("key")
	val := c.Query("value")
	_ = kv.Put(context.TODO(), key, val)
}

func getValue(c *gin.Context) {
	key := c.Param("key")

	val, _ := kv.Get(context.TODO(), key)
	c.JSON(200, map[string]interface{}{
		"value": val,
	})
}

func deleteValue(c *gin.Context) {
	key := c.Param("key")

	_ = kv.Delete(context.TODO(), key)
}
