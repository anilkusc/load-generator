FROM golang:1.14 as BUILD
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app .
FROM scratch
WORKDIR /app
COPY --from=BUILD /bin/app .
ENTRYPOINT ["./app"]