FROM golang:1.21 as monolith

WORKDIR /project

COPY go.mod .
RUN go mod download

COPY . /project
RUN go build -o /bin/monolith -v ./cmd/service

RUN rm -rf /project

CMD ["/bin/monolith"]