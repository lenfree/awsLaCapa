FROM golang

COPY . /

RUN make -f /Makefile script

WORKDIR /go/src/github.com/lenfree/awsLaCapa

CMD ["bee", "run"]
