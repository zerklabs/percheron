package percheron

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"net"
	"time"
)

// ACL constants for Objects
//
// Default: ACL_PRIVATE
const (
	ACL_PRIVATE = iota
	ACL_PUBLIC_READ
	ACL_PUBLIC_READ_WRITE
	ACL_AUTH_READ
)

var (
	// Grant-ACL supported headers
	// See: http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUTacl.html
	allowed_auth_headers = []string{
		"x-p-grant-read",
		"x-p-grant-write",
		"x-p-grant-read-acp",
		"x-p-grant-write-acp",
		"x-p-grant-full-control",
	}
)

type PerchStore struct {
	Path      string
	RedisHost string
	Pool      *redis.Pool
	Peers     []net.IPAddr
}

func NewPerchStore(folderPath string, redisHost string) *PerchStore {
	yes, err := DoesDirExist(folderPath)

	if err != nil {
		log.Fatal(err)
	}

	if yes {
		// no error means the stat returned successfully
		store := new(PerchStore)
		store.Path = folderPath
		store.RedisHost = redisHost

		return store
	}

	log.Fatalf("%s does not exist or cannot be accessed", folderPath)

	return new(PerchStore)
}

func (self *PerchStore) EstablishPool() {
	self.Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", self.RedisHost)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Fatal(err)
			}

			return err
		},
	}
}
