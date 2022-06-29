FROM golang:1.17-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

WORKDIR /go/src
ADD ../ccm/cmd ./chaincode
RUN cd ./chaincode \
&& go mod tidy \
&& go build  -o /go/bin/chaincode main.go

FROM alpine AS ccm
WORKDIR chaincode
COPY --from=builder /go/bin/chaincode chaincode
ENTRYPOINT [ "./chaincode"]
