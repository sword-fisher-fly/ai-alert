default: build

run:
	GIN_MODE=release go run main.go

build: build-linux build-windows build-web

build-windows:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o ai-alert.exe cmd/ai-model/main.go

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ai-alert ./...

build-linux-arm:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o ai-alert ./...

build-web:
	pushd
	cd web && npm ci && npm run build
	popd
	rm -rf internal/static/dist && mkdir -p internal/static/dist
	cp -rf web/build/* internal/static/dist

lint:
	env GOGC=25 golangci-lint run --fix -j 8 -v ./... --timeout=5m 