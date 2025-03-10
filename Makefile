-include .env
export $(shell sed 's/=.*//' .env)

GOMOCK_VERSION:=v0.4.0
GINKGO_VERSION:=v2.13.2

run:
	@go run src/cmd/p-n16-k8/main.go

test:
	@go run github.com/onsi/ginkgo/v2/ginkgo --cover -v ./src/...