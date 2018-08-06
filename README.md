# nooler-server
Remote alarm system (punch a button in your room and someone gets notified. Also works with smartwatches).

## Setup
Run mysql server (or have it run at startup):
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