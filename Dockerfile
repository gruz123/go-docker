FROM golang:1.17.9-buster
RUN mkdir /usr/local/go/src/zalupa
ADD . /usr/local/go/src/zalupa
RUN mkdir /usr/local/go/src/zalupa/img
ADD /zalupa/img  /usr/local/go/src/zalupa/img
WORKDIR /usr/local/go/src/zalupa
RUN go build -o zalupa .
CMD ["/usr/local/go/src/zalupa/zalupa"]
