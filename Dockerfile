FROM golang:1.22 as builder

WORKDIR /workspace
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o app main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /

COPY config.yaml .
COPY --from=builder /workspace/app .
USER nonroot:nonroot

ENTRYPOINT ["/app"]