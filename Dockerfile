# I probably don't recommend this but I'm big lazy and want to keep costs as low as possible.
# good luck :)

FROM golang:1.22 AS go-build

WORKDIR /workspace
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o app main.go

FROM node:22-alpine AS web-build

RUN mkdir -p /app
WORKDIR /app

COPY web/package.json ./
RUN npm install --frozen-lockfile

COPY web/ .
RUN npm run build

FROM node:22-alpine AS runner

RUN apk add --no-cache bash

RUN mkdir -p /app
WORKDIR /app

COPY --from=web-build /app/package*.json .
COPY --from=web-build /app/next.config.mjs ./
COPY --from=web-build /app/.next/standalone ./
COPY --from=web-build /app/.next/static ./.next/static

COPY config.yaml .
COPY hack/mono-start.sh ./mono-start.sh

COPY --from=go-build /workspace/app .

ENV HOSTNAME 0.0.0.0

EXPOSE 3000
#ENTRYPOINT ["node", "server.js"]
CMD ["/bin/bash", "mono-start.sh"]