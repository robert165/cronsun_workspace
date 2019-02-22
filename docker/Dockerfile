FROM golang:1.11.2-alpine
#FROM golang:1.10-alpine as builder

ADD docker/sources.list /etc/apt/sources.list


# Install basic dependencies
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc libc-dev

#如果你熟悉 Go，也可以从源码编译, 要求 go >= 1.11+
ENV GOROOT /usr/local/go
ENV GOPATH /data/go
ENV PATH $GOROOT/bin:$PATH
RUN mkdir -p $GOPATH/src


RUN mkdir -p $GOPATH/src/golang.org/x/
WORKDIR $GOPATH/src/golang.org/x/
RUN git clone https://github.com/golang/sys.git
WORKDIR $GOPATH/src/golang.org/x/sys
RUN go get ./...
#RUN go install sys

# github.com/golang/protobuf
RUN mkdir -p $GOPATH/src/github.com/golang/
WORKDIR $GOPATH/src/github.com/golang
RUN git clone https://github.com/golang/protobuf protobuf
WORKDIR $GOPATH/src
#RUN go get ./...
RUN go install github.com/golang/protobuf/proto
RUN go install github.com/golang/protobuf/protoc-gen-go

#golang.org/x/text
RUN mkdir -p $GOPATH/src/golang.org/x/
WORKDIR $GOPATH/src/golang.org/x/
RUN git clone https://github.com/golang/text.git text
#WORKDIR $GOPATH/src/golang.org/x/text
WORKDIR $GOPATH/src
RUN set GO111MODULE=on
RUN go install $GOPATH/src/golang.org/x/text
#RUN go get $GOPATH/src/golang.org/x/text/...
#RUN go get ./...

#golang.org/x/crypto/ssh/terminal
RUN mkdir -p $GOPATH/src/golang.org/x/
WORKDIR $GOPATH/src/golang.org/x/
RUN git clone https://github.com/golang/crypto.git
WORKDIR $GOPATH/src/golang.org/x/crypto
RUN go get ./...

#golang.org/x/net
RUN mkdir -p $GOPATH/src/golang.org/x/
WORKDIR $GOPATH/src/golang.org/x/
RUN git clone https://github.com/golang/net.git
WORKDIR $GOPATH/src/golang.org/x/net
RUN go get ./...
#RUN go install sys


#google.golang.org/genproto/googleapis/rpc/status
RUN mkdir -p $GOPATH/src/google.golang.org/
WORKDIR $GOPATH/src/google.golang.org
RUN git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
WORKDIR $GOPATH/src/
RUN go install google.golang.org/genproto/googleapis/rpc/status


WORKDIR $GOPATH/src
RUN set GO111MODULE=on


## install etcd
RUN mkdir -p $GOPATH/src/go.etcd.io/
RUN mkdir -p $GOPATH/src/github.com/coreos/etcd
WORKDIR $GOPATH/src/go.etcd.io/
COPY /etcd $GOPATH/src/go.etcd.io/etcd
COPY /etcd $GOPATH/src/github.com/coreos/etcd
COPY /etcd/clientv3 $GOPATH/src/github.com/coreos/etcd/clientv3




RUN mkdir -p $GOPATH/src/google.golang.org/
WORKDIR $GOPATH/src/google.golang.org
RUN git clone https://github.com/grpc/grpc-go $GOPATH/src/google.golang.org/grpc
WORKDIR $GOPATH/src/
RUN go install google.golang.org/grpc


#github.com/google/uuid
WORKDIR $GOPATH/src
RUN set GO111MODULE=on
RUN go get github.com/google/uuid


#github.com/grpc-ecosystem/go-grpc-middleware
WORKDIR $GOPATH/src
RUN set GO111MODULE=on
RUN go get github.com/grpc-ecosystem/go-grpc-middleware

#go.uber.org/atomic
RUN mkdir -p $GOPATH/src/go.uber.org/
WORKDIR $GOPATH/src/go.uber.org
RUN git clone https://github.com/uber-go/atomic $GOPATH/src/go.uber.org/atomic
WORKDIR $GOPATH/src/
RUN go install go.uber.org/atomic

#go.uber.org/multierr
RUN mkdir -p $GOPATH/src/go.uber.org/
WORKDIR $GOPATH/src/go.uber.org
RUN git clone https://github.com/uber-go/multierr $GOPATH/src/go.uber.org/multierr
WORKDIR $GOPATH/src/
RUN go install go.uber.org/multierr


#go.uber.org/zap
RUN mkdir -p $GOPATH/src/go.uber.org/
WORKDIR $GOPATH/src/go.uber.org
RUN git clone https://github.com/uber-go/zap $GOPATH/src/go.uber.org/zap
WORKDIR $GOPATH/src/
RUN go install go.uber.org/zap




WORKDIR $GOPATH/src
RUN set GO111MODULE=on
RUN go get $GOPATH/src/go.etcd.io/etcd
WORKDIR $GOPATH/src/
RUN go get  github.com/coreos/etcd
#RUN go get $GOPATH/src/github.com/coreos/etcd
#RUN go get $GOPATH/src/github.com/coreos/etcd/clientv3


WORKDIR $GOPATH/src
RUN go get -u github.com/shunfei/cronsun


#COPY /cronsun $GOPATH/src/cronsun
WORKDIR $GOPATH/github.com/shunfei/src/cronsun


RUN go mod vendor
# 如果 go mod vendor 下载失败，请尝试 https://goproxy.io


#解决docker终端不能输出中文问题
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
# python命令输出默认为ASCII，需要设定变量PYTHONIOENCODING为UTF-8
ENV PYTHONIOENCODING UTF-8

ADD build.sh /opt/config/
CMD ["/opt/config/build.sh"]

CMD ["cronweb","-conf conf/base.json"]
EXPOSE 7079


