EXEC_APP := docker-compose exec -T app
INTERFACES = $(shell $(EXEC_APP) find -name "interface.go" )
GO_VER = 1.20

gotest: lint
	@$(EXEC_APP) go test -coverprofile=cover.out ./...
	@$(EXEC_APP) go tool cover -html=cover.out -o cover.html

lint:
	@$(EXEC_APP) go mod tidy -go=${GO_VER}
	@$(EXEC_APP) golangci-lint run --fix --timeout 3m

up:
	@docker-compose up -d
	@$(EXEC_APP) go mod tidy -go=${GO_VER}

down:
	@docker-compose down

gen-mock:
	@$(foreach src,$(INTERFACES), $(EXEC_APP) sh tools/mockgen/codegen.sh ${src} || exit;)
