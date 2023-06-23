FROM golang:1.20

ENV PATH=${PATH}:${GOPATH}/bin

ENV REPOSITORY github.com/dip-dev/go-tutorial
ARG WORKDIR=${GOPATH}/src/${REPOSITORY}
WORKDIR ${WORKDIR}

# golangci-lint install
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.52.2 && \
    mv bin/golangci-lint /usr/local/bin/golangci-lint && \
    rmdir bin

RUN go install github.com/cosmtrek/air@v1.41.0

ENTRYPOINT [ "air", "-c", ".air.conf" ]