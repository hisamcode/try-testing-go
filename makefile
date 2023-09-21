
rw:
	go run ./webapp/cmd/web

tw:
	go test ./webapp/... -cover -v -count 1 -tags integration

cw:
	go test -coverprofile ./webapp/coverage.out ./webapp/...; go tool cover -html ./webapp/coverage.out

dw:
	cd webapp; docker compose up -d

dwd:
	cd webapp; docker compose down

twdbrepo:
	cd webapp/pkg/repository/dbrepo; go test . -v -tags integration

# primeapp
tp:
	go test ./primeapp/ -cover

cp:
	go test -coverprofile ./primeapp/coverage.out ./primeapp/; go tool cover -html ./primeapp/coverage.out

rp:
	go run ./primeapp/