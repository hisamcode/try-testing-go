tp:
	go test ./primeapp/ -cover

cp:
	go test -coverprofile ./primeapp/coverage.out ./primeapp/; go tool cover -html ./primeapp/coverage.out

rp:
	go run ./primeapp/ -v