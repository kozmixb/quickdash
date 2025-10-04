########## BUILDER ##########
FROM golang:1.25.1-trixie AS builder

COPY . /src
WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and 
# only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

# Build statically linked file and strip debug information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-extldflags=-static -s -w" -v -o app

########## RESULT ##########
FROM alpine:3.22

ARG USER_UID=65532
ARG USER_GID=65532
ARG USER_NAME=golang
ARG GROUP_NAME=golang

RUN addgroup -g ${USER_GID} ${GROUP_NAME} && \
    adduser -u ${USER_UID} -G ${GROUP_NAME} -D ${USER_NAME}

WORKDIR /app

COPY --from=builder --chown=${USER_NAME}:${GROUP_NAME} /src/app /app
COPY --chown=${USER_NAME}:${GROUP_NAME} ./views /app/views

USER ${USER_NAME}

EXPOSE 3000

CMD ["/app/app"]