from golang
MAINTAINER Senghoo Kim "shkdmb@gmail.com"
RUN go get github.com/tools/godep
ADD . /go/src/github.com/senghoo/web2pic
RUN cd /go/src/github.com/senghoo/web2pic && godep go install
EXPOSE 80
CMD ["/go/bin/web2pic", "--address", "0.0.0.0:80"]

