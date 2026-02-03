default: build

build: build-web build-linux build-windows 

build-windows:
	@echo 'Build window binary' 
    CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o bin/ai-alert.exe cmd/ai-model/main.go

build-linux:
	@echo 'Build linux binary'
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/ai-alert.linux cmd/ai-model/main.go


build-web:
	cd web && npm install && npm run build
	cd ../
	rm -rf internal/static/dist && mkdir -p internal/static/dist
	cp -rf web/build/* internal/static/dist
	

lint:
	env GOGC=25 golangci-lint run --fix -j 8 -v ./... --timeout=5m
