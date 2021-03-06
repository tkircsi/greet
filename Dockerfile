FROM golang:1.16 as build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/greetsrv ./greet_server

FROM scratch as bin
COPY --from=build /out/greetsrv /

EXPOSE 50051
ENTRYPOINT [ "/greetsrv" ]