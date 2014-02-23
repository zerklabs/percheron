package main

import (
	"flag"
	"github.com/zerklabs/auburn"
	"github.com/zerklabs/percheron"
	"log"
	"runtime"
)

var (
	listenIP        = flag.String("host", "", "IP to run the webserver on")
	listenOn        = flag.Int("listen", 8080, "Port to run the webserver on")
	certificatePath = flag.String("cert", "", "Certificate file for TLS (.pem) (Optional)")
	keyPath         = flag.String("key", "", "Private key for certificate (Required if cert given)")
	storePath       = flag.String("path", "", "Root of storage path")
	redisHost       = flag.String("redis", "127.0.0.1:6379", "Redis Server")

	store   *percheron.PerchStore
	users   []percheron.User
	objects []percheron.Object
	buckets []percheron.Bucket
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// bind the command line flags
	flag.Parse()

	server := &auburn.AuburnHttpServer{
		HttpPort: *listenOn,
		HttpIp:   *listenIP,
	}

	store = percheron.NewPerchStore(*storePath, *redisHost)

	go fetchUsers()
	go fetchBuckets()
	go fetchObjects()

	server.Handle("/lookup/users", lookupUsers)
	server.Handle("/lookup/objects", lookupObjects)
	server.Handle("/lookup/buckets", lookupBuckets)

	if len(*certificatePath) > 0 {
		// check if key given
		if len(*keyPath) > 0 {
			server.StartTLS(*certificatePath, *keyPath)
		}

		log.Fatal("Private key required to enable TLS mode")

	} else {
		server.Start()
	}
}

func fetchUsers() {
	conn := store.Pool.Get()
	defer conn.Close()
}

func fetchBuckets() {
	conn := store.Pool.Get()
	defer conn.Close()
}

func fetchObjects() {
	conn := store.Pool.Get()
	defer conn.Close()
}

func lookupUsers(req *auburn.AuburnHttpRequest) {

}

func lookupObjects(req *auburn.AuburnHttpRequest) {

}

func lookupBuckets(req *auburn.AuburnHttpRequest) {

}
