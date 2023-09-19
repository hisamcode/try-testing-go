```go
go test .\primeapp\
go test .\primeapp\ -cover
go test -coverprofile .\primeapp\coverage.out .\primeapp\
go tool cover -html .\primeapp\coverage.out
go test -coverprofile .\primeapp\coverage.out .\primeapp\; go tool cover -html .\primeapp\coverage.out
```