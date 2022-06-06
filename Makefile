test:
	go clean -testcache
	go test ./... -race -coverprofile cp.out
	go tool cover -html=./cp.out -o cover.html

integration-test:
	go clean -testcache
	export EXECUTE_INTEGRATION_TEST=TRUE && go test ./... -race -coverprofile cp.out

run:
	go run main.go < test/input-performance.txt
