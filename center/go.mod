module center

go 1.14

replace com => ../com

require (
	com v0.0.0-00010101000000-000000000000
	github.com/davyxu/cellnet v4.1.0+incompatible
	github.com/go-xorm/xorm v0.7.9
	github.com/nats-io/nats.go v1.9.1
	google.golang.org/grpc v1.26.0
)
