package examples

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var r = redis.NewClient(&redis.Options{})

// инкремент KeepTTL
func Incr(ctx context.Context, client *redis.Client, key string) error {

	script := `
		redis.call(‘INCR’, KEYS[1]) local ttl = redis.call(‘ttl’, KEYS[1]) if ttl < 0 then return redis.call(‘EXPIRE’, KEYS[1],  ARGV[1]) end
			`
	return client.Eval(ctx, script, []string{key}, []string{"300"}).Err()
}
