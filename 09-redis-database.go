package main

import (
	"fmt"
)

// COURSE 9: REDIS - IN-MEMORY DATA STORE
// Topics covered:
// 1. Redis basics
// 2. Data structures (strings, lists, sets, hashes, sorted sets)
// 3. Key-value operations
// 4. Expiration and TTL
// 5. Transactions
// 6. Pub/Sub
// 7. Connection pooling
// 8. Best practices

// Note: Requires "github.com/redis/go-redis/v9"

// ============ REDIS CONNECTION PATTERN ============

// func connectRedis(addr string) (*redis.Client, error) {
//	client := redis.NewClient(&redis.Options{
//		Addr: addr,
//	})
//	
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	
//	_, err := client.Ping(ctx).Result()
//	return client, err
// }

// ============ COURSE NINE MAIN FUNCTION ============
func courseNine() {
	fmt.Println("=== REDIS - IN-MEMORY DATA STORE ===\n")

	fmt.Println("REDIS SETUP:")
	fmt.Println("---\n")

	fmt.Println("Docker Redis:")
	fmt.Println(`docker run --name redis -d -p 6379:6379 redis:latest`)
	fmt.Println()

	fmt.Println("Connection:")
	fmt.Println(`client := redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})`)
	fmt.Println()

	fmt.Println("STRING OPERATIONS:")
	fmt.Println("---")
	fmt.Println(`
// SET key value
client.Set(ctx, "name", "Alice", 0)

// SET with expiration
client.Set(ctx, "session", "token123", 24*time.Hour)

// GET value
val, err := client.Get(ctx, "name").Result()

// EXISTS - check if key exists
exists, err := client.Exists(ctx, "name").Result()

// DEL - delete key
deleted, err := client.Del(ctx, "name").Result()

// INCR/DECR - atomic increment/decrement
count, err := client.Incr(ctx, "counter").Result()
count, err := client.Decr(ctx, "counter").Result()
count, err := client.IncrBy(ctx, "counter", 5).Result()

// APPEND - append to string
length, err := client.Append(ctx, "name", "Smith").Result()

// STRLEN - get string length
length, err := client.StrLen(ctx, "name").Result()

// GETRANGE - get substring
substr, err := client.GetRange(ctx, "name", 0, 2).Result()

// SETRANGE - set substring
length, err := client.SetRange(ctx, "name", 0, "Bob").Result()
`)
	fmt.Println()

	fmt.Println("LIST OPERATIONS:")
	fmt.Println("---")
	fmt.Println(`
// LPUSH - push to left
length, err := client.LPush(ctx, "queue", "task1", "task2").Result()

// RPUSH - push to right
length, err := client.RPush(ctx, "queue", "task3").Result()

// LPOP - pop from left
value, err := client.LPop(ctx, "queue").Result()

// RPOP - pop from right
value, err := client.RPop(ctx, "queue").Result()

// LLEN - get list length
length, err := client.LLen(ctx, "queue").Result()

// LRANGE - get range
values, err := client.LRange(ctx, "queue", 0, -1).Result()

// LTRIM - trim list
err := client.LTrim(ctx, "queue", 0, 10).Err()

// LINDEX - get element at index
value, err := client.LIndex(ctx, "queue", 0).Result()

// LREM - remove elements
removed, err := client.LRem(ctx, "queue", 1, "task1").Result()
`)
	fmt.Println()

	fmt.Println("SET OPERATIONS:")
	fmt.Println("---")
	fmt.Println(`
// SADD - add to set
added, err := client.SAdd(ctx, "tags", "go", "rust", "python").Result()

// SMEMBERS - get all members
members, err := client.SMembers(ctx, "tags").Result()

// SCARD - set cardinality (size)
count, err := client.SCard(ctx, "tags").Result()

// SISMEMBER - check membership
is, err := client.SIsMember(ctx, "tags", "go").Result()

// SREM - remove from set
removed, err := client.SRem(ctx, "tags", "go").Result()

// SUNION - union of sets
union, err := client.SUnion(ctx, "set1", "set2").Result()

// SINTER - intersection
inter, err := client.SInter(ctx, "set1", "set2").Result()

// SDIFF - difference
diff, err := client.SDiff(ctx, "set1", "set2").Result()

// SPOP - remove and return random member
member, err := client.SPop(ctx, "tags").Result()
`)
	fmt.Println()

	fmt.Println("HASH OPERATIONS:")
	fmt.Println("---")
	fmt.Println(`
// HSET - set hash fields
created, err := client.HSet(ctx, "user:1", "name", "Alice", "age", 30).Result()

// HGET - get single field
name, err := client.HGet(ctx, "user:1", "name").Result()

// HGETALL - get all fields
all, err := client.HGetAll(ctx, "user:1").Result()
// Returns map[string]string

// HEXISTS - check if field exists
exists, err := client.HExists(ctx, "user:1", "name").Result()

// HDEL - delete field
deleted, err := client.HDel(ctx, "user:1", "age").Result()

// HLEN - get number of fields
count, err := client.HLen(ctx, "user:1").Result()

// HKEYS - get all keys
keys, err := client.HKeys(ctx, "user:1").Result()

// HVALS - get all values
values, err := client.HVals(ctx, "user:1").Result()

// HINCRBY - increment field
newAge, err := client.HIncrBy(ctx, "user:1", "age", 1).Result()
`)
	fmt.Println()

	fmt.Println("SORTED SET OPERATIONS:")
	fmt.Println("---")
	fmt.Println(`
// ZADD - add to sorted set
added, err := client.ZAdd(ctx, "leaderboard", redis.Z{
	Score:  100,
	Member: "alice",
}, redis.Z{
	Score:  95,
	Member: "bob",
}).Result()

// ZRANGE - get range by rank
members, err := client.ZRange(ctx, "leaderboard", 0, -1).Result()

// ZREVRANGE - get range in reverse
members, err := client.ZRevRange(ctx, "leaderboard", 0, -1).Result()

// ZRANGE with scores
items, err := client.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()

// ZCARD - get sorted set size
count, err := client.ZCard(ctx, "leaderboard").Result()

// ZSCORE - get score of member
score, err := client.ZScore(ctx, "leaderboard", "alice").Result()

// ZRANK - get rank of member
rank, err := client.ZRank(ctx, "leaderboard", "alice").Result()

// ZREM - remove member
removed, err := client.ZRem(ctx, "leaderboard", "bob").Result()

// ZINCRBY - increment score
newScore, err := client.ZIncrBy(ctx, "leaderboard", 5, "alice").Result()

// ZCOUNT - count in score range
count, err := client.ZCount(ctx, "leaderboard", "90", "100").Result()
`)
	fmt.Println()

	fmt.Println("KEY OPERATIONS:")
	fmt.Println("---")
	fmt.Println(`
// KEYS - find keys by pattern
keys, err := client.Keys(ctx, "user:*").Result()

// SCAN - iterative key scan (better for large databases)
var keys []string
iter := client.Scan(ctx, 0, "user:*", 100).Iterator()
for iter.Next(ctx) {
	keys = append(keys, iter.Val())
}

// EXPIRE - set key expiration
ok, err := client.Expire(ctx, "key", 1*time.Hour).Result()

// PEXPIRE - expire in milliseconds
ok, err := client.PExpire(ctx, "key", 1000).Result()

// TTL - get remaining time to live
ttl, err := client.TTL(ctx, "key").Result()

// PERSIST - remove expiration
ok, err := client.Persist(ctx, "key").Result()

// TYPE - get key type
keyType, err := client.Type(ctx, "key").Result()

// RENAME - rename key
ok, err := client.Rename(ctx, "old", "new").Result()
`)
	fmt.Println()

	fmt.Println("TRANSACTIONS:")
	fmt.Println("---")
	fmt.Println(`
// MULTI/EXEC
pipe := client.Pipeline()

pipe.Set(ctx, "key1", "value1", 0)
pipe.Set(ctx, "key2", "value2", 0)
pipe.Get(ctx, "key1")

cmds, err := pipe.Exec(ctx)

// WATCH/UNWATCH - optimistic locking
err := client.Watch(ctx, func(tx *redis.Tx) error {
	// Get current value
	current, _ := tx.Get(ctx, "counter").Result()
	
	// Use transaction
	_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, "counter", "newValue", 0)
		return nil
	})
	
	return err
}, "counter")
`)
	fmt.Println()

	fmt.Println("PUB/SUB:")
	fmt.Println("---")
	fmt.Println(`
// SUBSCRIBE
pubsub := client.Subscribe(ctx, "events")
ch := pubsub.Channel()

for msg := range ch {
	fmt.Printf("Channel: %s, Message: %s\\n", msg.Channel, msg.Payload)
}

// PUBLISH
err := client.Publish(ctx, "events", "event data").Err()

// UNSUBSCRIBE
err := pubsub.Unsubscribe(ctx, "events")

// Pattern subscribe
pubsub := client.PSubscribe(ctx, "events:*")
`)
	fmt.Println()

	fmt.Println("PIPELINING (Batch Operations):")
	fmt.Println("---")
	fmt.Println(`
// Send multiple commands at once
pipe := client.Pipeline()

for i := 0; i < 100; i++ {
	pipe.Set(ctx, fmt.Sprintf("key:%d", i), i, 0)
}

_, err := pipe.Exec(ctx)

// Significantly faster than individual commands
`)
	fmt.Println()

	fmt.Println("SCRIPTING:")
	fmt.Println("---")
	fmt.Println(`
// Lua scripting for atomic operations
script := redis.NewScript(` + "`" + `
if redis.call('exists', KEYS[1]) == 1 then
	return redis.call('incr', KEYS[1])
else
	return redis.call('set', KEYS[1], 1)
end
` + "`" + `)

result, err := script.Run(ctx, client, []string{"counter"}).Result()
`)
	fmt.Println()

	fmt.Println("USE CASES:")
	fmt.Println("---")
	fmt.Println("✓ Session storage")
	fmt.Println("✓ Caching")
	fmt.Println("✓ Rate limiting")
	fmt.Println("✓ Job queues")
	fmt.Println("✓ Real-time leaderboards")
	fmt.Println("✓ Pub/Sub messaging")
	fmt.Println("✓ Counters and analytics")
	fmt.Println("✓ Distributed locks")
	fmt.Println("✓ Full-text search (with modules)")
	fmt.Println()

	fmt.Println("BEST PRACTICES:")
	fmt.Println("---")
	fmt.Println("✓ Use connection pooling")
	fmt.Println("✓ Set appropriate expiration times")
	fmt.Println("✓ Use pipelining for batch operations")
	fmt.Println("✓ Monitor memory usage")
	fmt.Println("✓ Use appropriate data structure for each task")
	fmt.Println("✓ Implement fallback if Redis unavailable")
	fmt.Println("✓ Set maxmemory and eviction policy")
	fmt.Println("✓ Use AOF or RDB for persistence")
	fmt.Println("✓ Replicate for high availability")
	fmt.Println()

	fmt.Println("COMMON LIBRARIES:")
	fmt.Println("---")
	fmt.Println("github.com/redis/go-redis/v9  - Official Redis client")
	fmt.Println("github.com/go-redis/cache     - Caching wrapper")
	fmt.Println()

	fmt.Println("=== END OF REDIS ===")
}

// KEY TAKEAWAYS:
// 1. Redis is an in-memory key-value store - extremely fast
// 2. Data is stored in RAM - not persistent by default
// 3. SET/GET for simple string values
// 4. Lists for queues and stacks
// 5. Sets for unique items and membership checks
// 6. Hashes for objects (like structs)
// 7. Sorted sets for leaderboards and rankings
// 8. Expiration (TTL) auto-deletes keys
// 9. Transactions ensure atomic operations
// 10. Pub/Sub for real-time messaging
// 11. Pipelining batches commands for performance
// 12. Use context for timeout handling
// 13. Always close connections when done
// 14. Redis is single-threaded - very predictable
// 15. Good for caching, sessions, rate limiting
// 16. Lua scripts allow complex atomic operations
// 17. WATCH for optimistic locking
// 18. Use appropriate data structure for each use case
// 19. Monitor memory - Redis stores everything in RAM
// 20. Use Redis Cluster or Sentinel for high availability
