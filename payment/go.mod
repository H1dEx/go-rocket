module github.com/H1dEx/go-rocket/payment

go 1.23.1

require (
	github.com/H1dEx/go-rocket/shared v0.0.0-20250713153242-c73c4eed3ccc
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.73.0
)

require (
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/H1dEx/go-rocket/shared => ../shared
