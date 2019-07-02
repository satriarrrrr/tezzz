# Tezzz

Playground for exploring unit testing in Golang

Install gomock & mockgen to create mock an interface:
```bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

Create a folder to store generated code by mockgen
```bash
mkdir -p mock/repayment
mockgen github.com/satriarrrrr/tezzz/repayment Repository > mock/repayment/mock_repository.go
```

How to run:
```bash
go run main.go --help
go run main.go -principal=500000 -apr=8.15 -mature=2020-04-10 -sold=2019-07-01 -taxrate=15
```

Run test coverage:
```bash
go test -v ./... -coverprofile=coverage.out
```