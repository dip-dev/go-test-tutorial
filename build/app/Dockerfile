FROM golang:1.19

ENV REPOSITORY github.com/dip-dev/go-test-tutorial
ARG WORKDIR=${GOPATH}/src/${REPOSITORY}

# golangci-lint install
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.0 && \
    mv bin/golangci-lint /usr/local/bin/golangci-lint && \
    rmdir bin

# mock-gen install
# RUN go install github.com/golang/mock/mockgen@v1.6.0
