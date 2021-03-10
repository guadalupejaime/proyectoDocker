FROM golang:1.15

COPY go.mod /proyecto/
COPY go.sum /proyecto/

COPY /cmd /proyecto/cmd
COPY /cmd /proyecto/cmd
COPY /pkg/ /proyecto/pkg/

# build a go binary
WORKDIR /proyecto/cmd/scraper/
RUN go build -o /proyecto/bin/scraper .

# build a go binary
WORKDIR /proyecto/cmd/api/
RUN go build -o /proyecto/bin/api .

# move to binaries path
WORKDIR /proyecto/bin/
