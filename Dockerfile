# FROM golang:1.22 as builder

# WORKDIR /workspace
# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o app main.go

# FROM gcr.io/distroless/static:nonroot
# WORKDIR /

# COPY config.yaml .
# COPY --from=builder /workspace/app .
# USER nonroot:nonroot

# ENTRYPOINT ["/app"]

FROM node:22-alpine as web-build

RUN mkdir -p /app
WORKDIR /app

COPY web/package.json ./
RUN npm install --frozen-lockfile

COPY web/ .
RUN npm run build

FROM node:22-alpine as web-runner
RUN mkdir -p /app
WORKDIR /app
COPY --from=web-build /app/package*.json .
COPY --from=web-build /app/next.config.mjs ./
COPY --from=web-build /app/.next/standalone ./
COPY --from=web-build /app/.next/static ./.next/static

ENV HOSTNAME 0.0.0.0

EXPOSE 3000
ENTRYPOINT ["node", "server.js"]
