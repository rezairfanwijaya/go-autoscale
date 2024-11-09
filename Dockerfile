FROM golang:latest

WORKDIR /APP

COPY . .

RUN go mod tidy

ARG PORT
ENV PORT=$PORT

RUN GOOS=linux go build -o go-autoscale

EXPOSE 6767

ENTRYPOINT [ "./go-autoscale" ]