ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM scratch

COPY --from=builder /run-app /usr/local/bin/
COPY  --chmod=777 ./public/bangs.json /bangs.json
ENV BANG_CONFIG_PATH=/bangs.json
CMD ["run-app"]
EXPOSE 8080
