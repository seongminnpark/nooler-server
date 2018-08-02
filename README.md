# nooler-server
Remote Alarm System

## Setup
Run mysql server:
```
sudo service mysql start
```

## Running 
Build command (at poject root directory):
```
go build -o ./nooler ./cmd/nooler/main.go
```
Run compiled binary:
```
./nooler
```

## Testing
```
go test -v ./test
```