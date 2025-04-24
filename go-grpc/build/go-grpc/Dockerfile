FROM golang:1.24.2-alpine AS builder

ENV HOME=/go/src/app
WORKDIR ${HOME}

RUN apk update && apk add git

ADD .air.toml ${HOME}/.air.toml
ADD go.mod ${HOME}/go.mod
ADD ./api ${HOME}/api
ADD "./cmd" "${HOME}/cmd/"
ADD ./internal ${HOME}/internal

RUN go mod tidy

FROM builder AS development

ENV HOME=/go/src/app
WORKDIR ${HOME}

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
