## primeapp
```go
go test .\primeapp\
go test .\primeapp\ -cover
go test -coverprofile .\primeapp\coverage.out .\primeapp\
go tool cover -html .\primeapp\coverage.out
go test -coverprofile .\primeapp\coverage.out .\primeapp\; go tool cover -html .\primeapp\coverage.out
go test -run Test_isPrime .\primeapp\ -v
// grouped
go test -run Test_alpha .\primeapp\ -v
```