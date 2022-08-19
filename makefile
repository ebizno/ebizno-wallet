gen:
	protoc --proto_path=adapter/product/grpc/protofile adapter/product/grpc/protofile/*.proto --go_out=adapter/product/grpc/pb --go-grpc_out=adapter/product/grpc/pb
	protoc --proto_path=adapter/cart/grpc/protofile adapter/cart/grpc/protofile/*.proto --go_out=adapter/cart/grpc/pb --go-grpc_out=adapter/cart/grpc/pb

mockgen:
	/home/paulo/go/bin/mockgen -source=domain/interfaces/repository/categorie_interface_repository.go -destination=domain/interfaces/mocks/categorie_interface_mock.go -package=mocks

test unit:
	go test -v ./test/unit

test all:
	go test -v ./... -coverprofile=coverage.out