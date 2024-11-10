FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

ARG PORT
ARG APP_NAME

ENV PORT=${PORT}
ENV APP_NAME=${APP_NAME}

RUN echo "PORT: \"$PORT\"" > application.yml
RUN echo "APP_NAME: \"$APP_NAME\"" >> application.yml

RUN GOOS=linux go build -o go-autoscale

EXPOSE ${PORT}

ENTRYPOINT [ "./go-autoscale" ]