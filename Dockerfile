FROM node:16 as builder

WORKDIR /build
COPY ./web .
COPY ./VERSION .
RUN npm install
RUN REACT_APP_VERSION=$(cat VERSION) npm run build

FROM golang AS builder2

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
COPY --from=builder /build/build ./web/build
RUN go mod download
RUN go build -ldflags "-s -w -X 'wechat-server/common.Version=$(cat VERSION)' -extldflags '-static'" -o wechat-server

FROM alpine

ENV PORT=3000
COPY --from=builder2 /build/wechat-server /
EXPOSE 3000
WORKDIR /data
ENTRYPOINT ["/wechat-server"]
