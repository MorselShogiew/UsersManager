package redis

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/MorselShogiew/UsersManager/config"
	"github.com/MorselShogiew/UsersManager/consts"
	"github.com/avast/retry-go"
	"github.com/go-redis/redis"
)

type RedisManager struct {
	Ok     bool
	client *redis.Client
}

type StoredUser struct {
	Id    int32
	Name  string
	Email string
}

type StoredUsers struct {
	Users []StoredUser
}

func (m *RedisManager) Connect(conf *config.RedisConfig) {
	var err error
	defer func() {
		if err != nil {
			m.client.Close()
		}
		m.Ok = err == nil
	}()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.DB,
	})

	err = retry.Do(
		func() error {
			_, err := client.Ping().Result()
			if err != nil {
				log.Println("Trying to connect to the Redis...")
			}
			return err
		},
		retry.Attempts(5),
		retry.Delay(time.Second),
		retry.DelayType(retry.BackOffDelay),
	)

	m.client = client
}

func (m *RedisManager) GetUsersFromRedis() (*[]StoredUser, bool) {
	fromRedis, err := m.client.Get(consts.RedisUsersKey).Bytes()
	if err == nil {
		buf := bytes.NewBuffer([]byte(fromRedis))
		dec := gob.NewDecoder(buf)
		var s StoredUsers
		err = dec.Decode(&s)
		if err == nil {
			return &s.Users, true
		} else {
			log.Printf("Error: getting users from Redis: %s", err)
		}
	}
	return nil, false
}

func (m *RedisManager) StoreUsersIntoRedis(users *[]StoredUser) {
	forRedis := &StoredUsers{Users: *users}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(forRedis)
	if err == nil {
		err = m.client.Set(consts.RedisUsersKey, buf.Bytes(), time.Minute).Err()
		if err != nil {
			log.Printf("Error: storing users into Redis: %s", err)
		}
	} else {
		log.Printf("Error: storing users into Redis: %s", err)
	}
}

func (m *RedisManager) Close() {
	m.client.Close()
}
