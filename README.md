awsLaCapa
==========

Build Status: [![Build Status](https://travis-ci.org/lenfree/awsLaCapa.svg?branch=master)](https://travis-ci.org/lenfree/awsLaCapa)

As a AWS user, it takes a large amount of time to grab all AWS resources with
logging into AWS console. Moreover, there are cli but they are use by resources
such as get a list of s3 buckets, we run s3 command, with ec2, we use ec2. This
serves as a quick way to get all information via API.

This is just another layer on top of it and abstracts all resources and make 
it simple.

Status: Work In Progress

[Binary Releases](https://github.com/lenfree/awsLaCapa/releases)

### Getting started:
----------------

### Install required packages:
```
$ make install
```

### Run test:
```
$ make script
```

### Start server:
```
$ cat>>~/.aws/credentials<<EOF
[default]
aws_access_key_id = <ACCESS_KEY_ID>
aws_secret_access_key = <SECRET_ACCESS_KEY>
EOF

$ cd /go/src/github.com/lenfree/awsLaCapa
$ make appstart
```

### Run Binary:
```
1. Download binary release
2. Create directory conf in same path as where you would run the binary file (mkdir conf)
3. cat >>conf/app.conf<<EOF
appname = awsLaCapa
httpport = 8080
runmode = prod
autorender = true
copyrequestbody = true
EnableDocs = true
EOF
4. Copy API doco static directory(https://github.com/lenfree/awsLaCapa/tree/master/swagger) to same project directory
```

[More info on running Binary plus static files](https://github.com/chu888chu888/Go-beego/blob/master/docs/en/Quickstart.md#deployment)

### Rest API Documentation available in below URL:
```
http://localhost:8081/swagger/#!/s3/S3Controller_GetAll
```


### Retrieve all S3 buckets
```
$ curl localhost:8081/v1/s3
[
  {
    "bucket_name": "mybucket-1",
    "bucket_creation_date": "2016-09-08T06:49:21Z"
  },
  {
    "bucket_name": "mybucket-2",
    "bucket_creation_date": "2016-09-08T06:49:20Z"
  }
]
```

### Run in docker container:
```
$ git clone https://github.com/lenfree/awsLaCapa.git
$ cd awsLaCapa
$ make build
$ cat>>credentials<<EOF
[default]
aws_access_key_id = <ACCESS_KEY_ID>
aws_secret_access_key = <SECRET_ACCESS_KEY>
EOF
$ make run HOST_PORT=8081
```

### Contributing:
```
1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
```
