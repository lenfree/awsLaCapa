FROM golang:1.7.1

ENV GOPATH=/go

COPY . /go/src/github.com/lenfree/awsLaCapa/

WORKDIR /go/src/github.com/lenfree/awsLaCapa

# This will install all required packages and run test
RUN make script

# This starts the app and web server
CMD ["make", "appstart"]
