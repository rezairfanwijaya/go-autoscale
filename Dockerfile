FROM golang:latest

WORKDIR /app

COPY go-autoscale .

RUN ls -l

# RUN go mod tidy

ARG PORT
ARG APP_NAME

ENV PORT=${PORT}
ENV APP_NAME=${APP_NAME}

EXPOSE ${PORT}

ENTRYPOINT [ "./go-autoscale" ]