#cli
cliv:
	go run ./webapp/cmd/cli -action=valid

clie:
	go run ./webapp/cmd/cli -action=expired

cli-HS384:
	go run ./webapp/cmd/cli -action=HS384

# api
ra:
	go run ./webapp/cmd/api

ta:
	# go test ./webapp/cmd/api/... -cover -v -count 1 -tags integration
	go test ./webapp/cmd/api/... -cover -v -count 1

ca:
	go test -coverprofile ./webapp/coverage.out ./webapp/cmd/api/...; go tool cover -html ./webapp/coverage.out

# webapp
rw:
	go run ./webapp/cmd/web

tw:
	# go test ./webapp/... -cover -v -count 1 -tags integration
	go test ./webapp/... -cover -v -count 1

cw:
	go test -coverprofile ./webapp/coverage.out ./webapp/...; go tool cover -html ./webapp/coverage.out

twdbrepo:
	cd webapp/pkg/repository/dbrepo; go test . -v -tags integration

# primeapp
tp:
	go test ./primeapp/ -cover

cp:
	go test -coverprofile ./primeapp/coverage.out ./primeapp/; go tool cover -html ./primeapp/coverage.out

rp:
	go run ./primeapp/

# docker
dc:
	cd webapp; docker compose up -d

dd:
	cd webapp; docker compose down