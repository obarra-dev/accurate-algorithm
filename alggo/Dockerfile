FROM golang:latest

COPY --chmod=755 run-test.sh /go/run-test.sh

COPY . /go/alggo/

WORKDIR /go

CMD ./run-test.sh

#docker build -t alggo-golang .
#docker run alggo-golang