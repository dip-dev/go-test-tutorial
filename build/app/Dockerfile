FROM golang:1.20

ENV REPOSITORY github.com/dip-dev/go-test-tutorial
ARG WORKDIR=${GOPATH}/src/${REPOSITORY}

# golangci-lint install
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.52.2 && \
    mv bin/golangci-lint /usr/local/bin/golangci-lint && \
    rmdir bin

# mock-gen install
RUN go install go.uber.org/mock/mockgen@latest
