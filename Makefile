EXEC_APP := docker-compose exec -T app
INTERFACES = $(shell $(EXEC_APP) find -name "interface.go" )
PKGS = $(shell $(EXEC_APP) go list ./... | grep -v vendor )
GO_VER = 1.19

gotest:
	@$(EXEC_APP) go mod tidy -go=${GO_VER}
	@$(EXEC_APP) go test -coverprofile=cover.out ./...
	@$(EXEC_APP) go tool cover -html=cover.out -o cover.html

lint:
	@$(EXEC_APP) go mod tidy -go=${GO_VER}
	@$(EXEC_APP) golangci-lint run --fix --timeout 3m

# コンテナ起動
up:
	@docker-compose up -d
	@$(EXEC_APP) go mod tidy -go=${GO_VER}

down:
	@docker-compose down

gen-mock:
	@$(foreach src,$(INTERFACES), $(EXEC_APP) sh tools/mockgen/codegen.sh ${src} || exit;)

create-gitconfig:
	cat ./build/app/.gitconfig.template | sed s/\<GITHUB_TOKEN\>/${GITHUB_TOKEN}/g > ./build/app/.gitconfig
