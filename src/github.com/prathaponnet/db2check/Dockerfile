# A hello world example with Go
#FROM golang:1.13-alpine
#RUN apk add  gcc
#RUN apk update && apk add git 


FROM golang:1.13
RUN apt-get update -y && apt install git && apt install libxml2 -y
RUN apt install gcc
MAINTAINER s.prathap@gmail.com


WORKDIR $GOPATH/src/github.com/prathaponnet/db2check

COPY . $GOPATH/src/github.com/prathaponnet/db2check

RUN mkdir -p /opt/ibm/db2/ && \ 
    mv $GOPATH/src/github.com/prathaponnet/db2check/linuxx64_odbc_cli.tar.gz /opt/ibm/db2/ && \
    cd /opt/ibm/db2/ && \
    tar -xzf linuxx64_odbc_cli.tar.gz && \
    rm -rf linuxx64_odbc_cli.tar.gz

ENV IBM_DB_DIR="/opt/ibm/db2/clidriver"
ENV IBM_DB_HOME="${IBM_DB_DIR}"
ENV IBM_DB_LIB="${IBM_DB_DIR}/lib"
ENV IBM_DB_INCLUDE="${IBM_DB_DIR}/include"
ENV DB2_HOME="${IBM_DB_DIR}/include"
ENV INCLUDE="${IBM_DB_DIR}/include"
ENV DB2LIB="${IBM_DB_DIR}/lib"
ENV PATH="${IBM_DB_DIR}/bin:${PATH}"
#ENV LD_LIBRARY_PATH="${IBM_DB_DIR}/lib:${LD_LIBRARY_PATH}"
ENV LD_LIBRARY_PATH="${IBM_DB_DIR}/lib"

ENV CGO_ENABLED=1
ENV CGO_LDFLAGS="-L${IBM_DB_DIR}/lib"
ENV CGO_CFLAGS="-I${IBM_DB_DIR}/include"

RUN CGO_ENABLED=1 go build -o app  

ENTRYPOINT ./app
EXPOSE 80