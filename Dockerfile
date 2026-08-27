FROM golang:1.26 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/cbdb ./cmd/cli/main.go

# Bake the read-only CBDB SQLite file into the image (no PVC needed).
COPY data/cbdb.sqlite3 /data/cbdb.sqlite3

WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/bin/cbdb", "server", "--db=/data/cbdb.sqlite3"]