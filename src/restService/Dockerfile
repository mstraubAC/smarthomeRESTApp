# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
## 
FROM golang:1.19-alpine AS build

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o restService -ldflags="-w -s"

##
## STEP 2 - DEPLOY
##
FROM scratch
WORKDIR /
COPY --from=build /app/restService /restService
# COPY --from=build /app/envs/dev /envs/dev

EXPOSE 7777
ENTRYPOINT [ "/restService" ]