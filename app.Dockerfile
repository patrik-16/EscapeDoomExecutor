FROM test

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY *.java ./
COPY .Dockerfile ./.Dockerfile

RUN go build -o /go-docker-demo

CMD [ "/go-docker-demo" ]