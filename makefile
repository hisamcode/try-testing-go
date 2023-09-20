
rw:
	go run ./webapp/cmd/web

tw:
	go test ./webapp/... -cover -v

cw:
	go test -coverprofile ./webapp/coverage.out ./webapp/...; go tool cover -html ./webapp/coverage.out

# primeapp
tp:
	go test ./primeapp/ -cover

cp:
	go test -coverprofile ./primeapp/coverage.out ./primeapp/; go tool cover -html ./primeapp/coverage.out

rp:
	go run ./primeapp/