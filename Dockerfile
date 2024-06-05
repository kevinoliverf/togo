
FROM arm64v8/golang:1.20

WORKDIR /app

COPY go.* ./
RUN go mod download


COPY . ./
RUN go mod tidy
RUN go build -v -o bin/ all 


EXPOSE 8000

CMD [ "bin/proto_handler" ]
