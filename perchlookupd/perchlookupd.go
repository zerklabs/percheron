package percherond

import (
	"flag"
	"github.com/cabrel/auburn"
	"github.com/cabrel/percheron"
	"log"
	"runtime"
)

// var redisServer = flag.String("host", "127.0.0.1", "Redis Server")
// var redisServerPort = flag.Int("port", 6379, "Redis Server Port")
// var responseUrl = flag.String("url", "https://passwords.cobhamna.com", "Server Response URL")
// var redisKeyPrefix = flag.String("prefix", "masq-prod", "Key prefix in Redis for Dictionary")
var listenIP = flag.String("host", "", "IP to run the webserver on")
var listenOn = flag.Int("listen", 8080, "Port to run the webserver on")
var certificatePath = flag.String("cert", "", "Certificate file for TLS (.pem) (Optional)")
var keyPath = flag.String("key", "", "Private key for certificate (Required if cert given)")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// bind the command line flags
	flag.Parse()

	server := &auburn.AuburnHttpServer{
		HttpPort: *listenOn,
		HttpIp:   *listenIP,
	}

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
