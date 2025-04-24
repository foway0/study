FROM golang:1.24.2-alpine AS builder

ENV HOME=/go/src/app
WORKDIR ${HOME}

RUN apk update && apk add git

ADD go.mod ${HOME}/go.mod
ADD ./api ${HOME}/api
ADD "./cmd" "${HOME}/cmd/"
ADD ./internal ${HOME}/internal

RUN go mod tidy

FROM builder AS development

ENV HOME=/go/src/app
WORKDIR ${HOME}

COPY --from=builder ${HOME} ${HOME}

CMD go run ${HOME}/cmd/go-grpc/main.go
