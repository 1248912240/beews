# This is the first stage, for building things that will be required by the
# final stage (notably the binary)
FROM airdb/beego
MAINTAINER Dean dean@airdb.com

ENV GO111MODULE "on"
WORKDIR  /go/src/github.com/airdb/beews
# COPY go.mod $WORKDIR
COPY . $WORKDIR
RUN export GO111MODULE=on && \
	export GOPROXY=https://goproxy.io && \
    go mod download

RUN bee pack
#ADD . $WORKDIR
#RUN godep get -v

# The second and final stage
RUN echo ${WORKDIR}/${APPNAME}

#FROM scratch
FROM alpine

# Copy the binary from the builder stage
COPY --from=0 beews.tar.gz /srv

EXPOSE 8080

CMD ["/bin/bash"]
