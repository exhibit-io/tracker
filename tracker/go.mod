module tracker

go 1.22.5

replace github.com/exhibit-io/tracker/config => ./config

require (
	github.com/exhibit-io/tracker/config v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.5
	github.com/julienschmidt/httprouter v1.3.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
