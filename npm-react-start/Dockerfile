FROM golang AS build

WORKDIR /main
COPY main.go /main
RUN go build main.go

FROM node:lts

VOLUME /volume
COPY --from=build /main/main /main
CMD /main