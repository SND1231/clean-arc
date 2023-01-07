# stage for development, which contains tools for code generation and debugging.
FROM golang:1.19-bullseye as dev

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        wget \
        make \
        unzip \
        git \
        clang-format \
        vim \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# install mockgen
RUN go install github.com/golang/mock/mockgen@v1.6.0

# install sqlboiler
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

WORKDIR /work/src

WORKDIR /work

COPY . ./