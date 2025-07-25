FROM golang:1.22.2 as build

WORKDIR /home/crawler
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o main ./main.go

FROM gcr.io/distroless/static:nonroot
# FROM golang:1.22.2

COPY --from=build /home/crawler/main /bin/main

CMD [ "/bin/main" ]