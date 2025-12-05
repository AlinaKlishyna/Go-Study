FROM golang:latest

COPY . /arithmetic-operation

WORKDIR /arithmetic-operation

EXPOSE 8000

RUN go run ./arithmetic-operation

CMD [ "go", "run", "./arithmetic-operation" ]