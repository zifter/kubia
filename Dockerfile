FROM golang:1.15-alpine

WORKDIR /go/src/
COPY . .

RUN go install .

CMD ["k8s"]
