package percheron

import (
	// "github.com/garyburd/redigo/redis"
	"log"
	"net"
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

// Grant-ACL supported headers
// See: http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUTacl.html
var allowed_auth_headers = []string{
	"x-perch-grant-read",
	"x-perch-grant-write",
	"x-perch-grant-read-acp",
	"x-perch-grant-write-acp",
	"x-perch-grant-full-control",
}

type PerchStore struct {
	Path      string
	RedisHost string
	RedisPort int
	Peers     []net.IPAddr
}

func NewPerchStore(folderPath string) *PerchStore {
	yes, err := DoesDirExist(folderPath)

	if err != nil {
		log.Fatal(err)
	}

	if yes {
		// no error means the stat returned successfully
		store := new(PerchStore)
		store.Path = folderPath

		return store
	}

	log.Fatalf("%s does not exist or cannot be accessed", folderPath)

	return new(PerchStore)
}
