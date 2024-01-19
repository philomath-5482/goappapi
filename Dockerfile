FROM golang:alpine As build
WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /src/gobuildapp



FROM scratch
WORKDIR /app
COPY *.html ./
COPY --from=0 /src/gobuildapp /app
EXPOSE 8181
ENTRYPOINT [ "/app/gobuildapp" ]






