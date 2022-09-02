// github.com/go-redis/redis/v8
package redis

import (
	r "github.com/go-redis/redis/v8"
)

func NewClient(options *r.Options) *r.Client {
	return r.NewClient(options)
}
