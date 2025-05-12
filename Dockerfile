ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM scratch

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder --chmod=777 /usr/src/app/bangs.json /bangs.json
CMD ["run-app"]
