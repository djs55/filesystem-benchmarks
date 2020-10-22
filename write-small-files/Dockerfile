FROM golang AS build

WORKDIR /main
COPY main.go /main
RUN go build main.go

FROM alpine:3.12

VOLUME /volume
COPY --from=build /main/main /main
CMD /main