SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o ../zdoc/b0pass-mac/b0pass .

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ../zdoc/b0pass-linux/b0pass .

SET CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o ../zdoc/b0pass-win/b0pass.exe .