FROM golang:1.17.9-buster
RUN mkdir /usr/local/go/src/img-input/img
ADD . /usr/local/go/src/img-input
WORKDIR /usr/local/go/src/img-input
RUN go build -o ii .
CMD ["/usr/local/go/src/img-input/ii"]
