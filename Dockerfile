FROM golang:1.16.4-alpine3.12 as base
RUN apk update && \
    apk add --update make && \
    apk add --update git && \
    apk add build-base
WORKDIR /src
COPY ./go.mod ./go.sum ./
#COPY ./go.mod ./
RUN go mod download
COPY ./ ./
RUN make build

FROM base as dev
RUN go get github.com/codegangsta/gin
CMD gin -p=3202

FROM golang:1.16.2-alpine3.12 as prod
COPY --from=base /src/build .
EXPOSE 3202
CMD ["./build"]
