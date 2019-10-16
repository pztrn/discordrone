FROM golang:1.13.1-alpine AS build

WORKDIR /discordrone
COPY . .

RUN CGO_ENABLED=0 go build -tags netgo -ldflags '-w -extldflags "-static"'

FROM alpine:3.10
LABEL maintainer "Stanislav N. <pztrn@pztrn.name>"

COPY --from=build /discordrone/discordrone /app/discordrone
RUN apk add --no-cache ca-certificates

ENTRYPOINT [ "/app/discordrone" ]
