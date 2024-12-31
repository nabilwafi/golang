package golangredis

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)
}

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)

	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Nabil Wafi", time.Second*3)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Nabil Wafi", result)

	time.Sleep(5 * time.Second)

	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Muhammad")
	client.RPush(ctx, "names", "Nabil")
	client.RPush(ctx, "names", "Wafi")

	assert.Equal(t, "Muhammad", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Nabil", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Wafi", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Muhammad")
	client.SAdd(ctx, "students", "Muhammad")
	client.SAdd(ctx, "students", "Nabil")
	client.SAdd(ctx, "students", "Nabil")
	client.SAdd(ctx, "students", "Wafi")
	client.SAdd(ctx, "students", "Wafi")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Muhammad", "Nabil", "Wafi"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Muhammad"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Nabil"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Wafi"})

	assert.Equal(t, []string{"Nabil", "Wafi", "Muhammad"}, client.ZRange(ctx, "scores", 0, 2).Val())
	assert.Equal(t, "Muhammad", client.ZPopMax(ctx, "scores", 1).Val()[0].Member)
	assert.Equal(t, "Wafi", client.ZPopMax(ctx, "scores", 1).Val()[0].Member)
	assert.Equal(t, "Nabil", client.ZPopMax(ctx, "scores", 1).Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Nabil")
	client.HSet(ctx, "user:1", "email", "nabil@example.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Nabil", user["name"])
	assert.Equal(t, "nabil@example.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 106.822702,
		Latitude:  -6.177590,
	})

	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 106.820889,
		Latitude:  -6.174964,
	})

	distance := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()
	assert.Equal(t, 0.3543, distance)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.820889,
		Latitude:   -6.174964,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()

	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "muhammad", "nabil", "wafi")
	client.PFAdd(ctx, "visitors", "nabil", "jaka", "wahyu")
	client.PFAdd(ctx, "visitors", "rizki", "jaka", "wahyu")

	total := client.PFCount(ctx, "visitors").Val()
	assert.Equal(t, int64(6), total)

	client.Del(ctx, "visitors")
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.SetEx(ctx, "name", "Nabil", 5*time.Second)
		p.SetEx(ctx, "address", "Indonesia", 5*time.Second)

		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Nabil", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.SetEx(ctx, "name", "Joko", 5*time.Second)
		p.SetEx(ctx, "address", "Jakarta", 5*time.Second)

		return nil
	})
	assert.Nil(t, err)

	assert.Equal(t, "Joko", client.Get(ctx, "name").Val())
	assert.Equal(t, "Jakarta", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"Name":    "Nabil-" + strconv.Itoa(i),
				"address": "Indonesia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    time.Second * 5,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.Values)
		}
	}
}

func TestSubscribePubSub(t *testing.T) {
	Subscriber := client.Subscribe(ctx, "channel-1")
	defer Subscriber.Close()
	for i := 0; i < 10; i++ {
		message, err := Subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello-"+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
