package main

import (
	"flag"
	"github.com/zerklabs/auburn"
	"github.com/zerklabs/percheron"
	"log"
	"runtime"
)

var listenIP = flag.String("host", "", "IP to run the webserver on")
var listenOn = flag.Int("listen", 8080, "Port to run the webserver on")
var certificatePath = flag.String("cert", "", "Certificate file for TLS (.pem) (Optional)")
var keyPath = flag.String("key", "", "Private key for certificate (Required if cert given)")
var storePath = flag.String("path", "", "Root of storage path")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// bind the command line flags
	flag.Parse()

	server := &auburn.AuburnHttpServer{
		HttpPort: *listenOn,
		HttpIp:   *listenIP,
	}

	store := &percheron.NewPerchStore(storePath)
	user := store.NewUserInfo("cabrel@zerklabs.com")
	log.Print(user.Marshal())

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

func fetchUsers(perch *percheron.PerchStore) {

}

func fetchBuckets(perch *percheron.PerchStore) {

}

func fetchObjects(perch *percheron.PerchStore) {

}

func lookupUsers(req *auburn.AuburnHttpRequest) {

}

func lookupObjects(req *auburn.AuburnHttpRequest) {

}

func lookupBuckets(req *auburn.AuburnHttpRequest) {

}
