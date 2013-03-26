export GOPATH=$PWD

echo ======= RUNNING BASIC ============================
go run llcount.go

echo ======= RUNNING CONCURRENT WITH MAXPROCS=1 =======
export GOMAXPROCS=1
go run pllcount.go

echo ======= RUNNING CONCURRENT WITH MAXPROCS=4 =======
export GOMAXPROCS=4
go run pllcount.go

echo ======= RUNNING CONCURRENT WITH MAXPROCS=8 =======
export GOMAXPROCS=8
go run pllcount.go
