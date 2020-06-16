# load-generator
Http load generator with concurrency support.<br>
For more information:<br>
LoadGenerator --help<br>
# BUILD FOR WINDOWS
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./LoadGenerator.exe .<br>
# BUILD FOR LINUX
CGO_ENABLED=0 go build -o ./LoadGenerator .
