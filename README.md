# nooler-server
Remote Alarm System

## Setup
Run mysql server:
```
sudo service mysql start
```

## Running 
Build and run (at project root):
```
go build -o ./nooler ./cmd/nooler/main.go
./nooler
```

## Testing
```
go test -v ./test
```