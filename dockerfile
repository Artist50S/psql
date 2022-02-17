FROM golang:1.17
RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go ./
RUN go build -o /app
EXPOSE 8000
CMD [ "/app" ]