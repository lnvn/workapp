FROM golang:1.22.2 AS build

ARG BINARY_NAME=main
WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ${BINARY_NAME} ./main.go

FROM gcr.io/distroless/static:nonroot
# FROM golang:1.22.2
ARG BINARY_NAME

COPY --from=build /app/${BINARY_NAME} /bin/${BINARY_NAME}

CMD [ "/bin/main" ]