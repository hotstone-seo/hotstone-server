package cachekit_test

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/hotstone-seo/hotstone-seo/pkg/cachekit"
	"github.com/stretchr/testify/require"
)

func ExampleCache() {
	var (
		data   string
		server *miniredis.Miniredis
		err    error
	)

	// run redis server
	if server, err = miniredis.Run(); err != nil {
		log.Fatal(err.Error())
	}
	defer server.Close()

	// create redis client
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})

	// define key and refresh function for your cache
	cache := cachekit.New("some-key", func() (interface{}, error) {
		return "fresh-data", nil
	})

	// execute cache to get the data
	if err = cache.Execute(client, &data, pragmaWithCacheControl("")); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(data)

	// Output:
	// fresh-data

}

func TestCache(t *testing.T) {
	var target bean

	testRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer testRedis.Close()

	client := redis.NewClient(&redis.Options{Addr: testRedis.Addr()})

	t.Run("GIVEN cache not available", func(t *testing.T) {
		t.Run("WHEN refresh failed", func(t *testing.T) {
			cache := cachekit.New("key", func() (interface{}, error) {
				return nil, errors.New("some-refresh-error")
			})
			require.EqualError(t, cache.Execute(client, &target, pragmaWithCacheControl("")), "Cache: RefreshFunc: some-refresh-error")
		})
		t.Run("WHEN marshal failed", func(t *testing.T) {
			cache := cachekit.New("key", func() (interface{}, error) {
				return make(chan int), nil
			})
			require.EqualError(t, cache.Execute(client, &target, pragmaWithCacheControl("")), "Cache: Marshal: json: unsupported type: chan int")
		})
		t.Run("WHEN failed to save to redis", func(t *testing.T) {
			badClient := redis.NewClient(&redis.Options{Addr: "wrong-addr"})
			cache := cachekit.New("key", func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})
			require.EqualError(t, cache.Execute(badClient, &target, pragmaWithCacheControl("")), "Cache: Set: dial tcp: address wrong-addr: missing port in address")
		})
		t.Run("", func(t *testing.T) {
			// monkey patch time.Now
			defer monkey.Patch(time.Now, func() time.Time {
				return time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
			}).Unpatch()

			cache := cachekit.New("key", func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})

			pragma := pragmaWithCacheControl("")
			require.NoError(t, cache.Execute(client, &target, pragma))

			// check target
			require.Equal(t, bean{Name: "new-name"}, target)

			// check data in redis
			require.Equal(t, `{"Name":"new-name"}`, client.Get("key").Val())
			require.Equal(t, 30*time.Second, client.TTL("key").Val())

			// check pragma
			require.Equal(t, "Thu, 16 Feb 2017 00:00:30 UTC", pragma.ResponseHeaders()[cachekit.HeaderExpires])
		})
	})

	t.Run("GIVEN cache available", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			// monkey patching time.Now
			defer monkey.Patch(time.Now, func() time.Time {
				return time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
			}).Unpatch()

			// set cache n redis
			testRedis.Set("key", `{"name":"cached"}`)
			testRedis.SetTTL("key", 10*time.Second)

			cache := cachekit.New("key", func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})

			pragma := pragmaWithCacheControl("")
			require.NoError(t, cache.Execute(client, &target, pragma))

			// Check target
			require.Equal(t, bean{Name: "cached"}, target)

			// check pragma
			require.Equal(t, "Thu, 16 Feb 2017 00:00:10 UTC", pragma.ResponseHeaders()[cachekit.HeaderExpires])
		})
		t.Run("WHEN cache-control: no-cache", func(t *testing.T) {
			testRedis.Set("key", `{"name":"cached"}`)
			cache := cachekit.New("key", func() (interface{}, error) {
				return &bean{Name: "new-name"}, nil
			})

			require.NoError(t, cache.Execute(client, &target, pragmaWithCacheControl("no-cache")))

			require.Equal(t, bean{Name: "new-name"}, target)

			require.Equal(t, `{"Name":"new-name"}`, client.Get("key").Val())
			require.Equal(t, 30*time.Second, client.TTL("key").Val())
		})
	})
}

type bean struct {
	Name string
}
