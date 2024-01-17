FROM golang:alpine
WORKDIR /src
COPY go.mod /src/
COPY go.sum /src/
COPY main.go /src/
COPY  *.html /src/
RUN go build -o /src/gobuildapp



FROM golang:alpine
WORKDIR /app
COPY go.mod /app
COPY *.html /app
COPY --from=0 /src/gobuildapp /app
EXPOSE 8181
CMD [ "/app/gobuildapp" ]






