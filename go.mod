module github.com/ubunifupay/visa

go 1.12

require (
	github.com/golang/protobuf v1.3.1
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-grpc v1.0.1
	github.com/micro/go-micro v1.1.0
	go.etcd.io/bbolt v1.3.2 // indirect
	google.golang.org/grpc v1.20.1
)

replace github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3 => github.com/golang/lint v0.0.0-20190409202823-5614ed5bae6fb75893070bdc0996a68765fdd275

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20190108154635-47c0da630f72
