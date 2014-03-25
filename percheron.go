package percheron

import (
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
