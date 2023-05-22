FROM test

WORKDIR /app

COPY go.mod go.sum  ./
RUN go mod download

COPY *.go ./
COPY *.java ./
COPY *.Dockerfile ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-demo

CMD [ "/go-docker-demo" ]