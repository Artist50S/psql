FROM golang:1.17
RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go ./
RUN go build -o /app
EXPOSE 80 5432  
HEALTHCHECK --interval=5s --timeout=10s --retries=3 CMD curl -sS 127.0.0.1:80 || exit 1
CMD [ "/app" ]
