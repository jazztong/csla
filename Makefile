default: test
TEST?=./...
BUILD_PATH = build

binary:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-extldflags "-static"' -a -tags netgo -installsuffix netgo -o bin/csla

test:
	mkdir -p $(BUILD_PATH)
	go test -v -cover -coverprofile=$(BUILD_PATH)/coverage.out -json -race $(TEST) > $(BUILD_PATH)/report.json

release:
	standard-version
	git push --follow-tags origin HEAD:master

template:
	go-bindata -o provider/fileloader/asset.go template/...
