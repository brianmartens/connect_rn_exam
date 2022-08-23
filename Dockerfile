FROM golang:1.19

RUN apt-get update
RUN apt install imagemagick -y

RUN mkdir /build
RUN cd /build
WORKDIR /build

COPY answer/ .

RUN go build

ENTRYPOINT /build/answer