FROM golang:1.17.9-buster
RUN mkdir /usr/local/go/src/zalupa
ADD . /usr/local/go/src/zalupa
ADD ../storage/zalupa_img  /usr/local/go/src/zalupa/img
WORKDIR /usr/local/go/src/zalupa
RUN go build -o zalupa .
CMD ["/usr/local/go/src/zalupa/zalupa"]
