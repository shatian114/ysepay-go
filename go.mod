module github.com/shatian114/ysepay-go

go 1.27.2

require (
	github.com/emmansun/gmsm v0.45.1
	github.com/shatian114/putils-go v0.0.1
	github.com/shatian114/verifysign v0.0.1
)

require golang.org/x/crypto v0.58.0 // indirect

replace (
	github.com/shatian114/putils-go => ../../../putils-go
	github.com/shatian114/verifysign => ../verifysign
)
