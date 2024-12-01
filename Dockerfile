FROM golang:latest

ARG APP_PORT
ARG WORKER_PORT
ARG APP_NAME

ENV APP_PORT=${APP_PORT}
ENV WORKER_PORT=${WORKER_PORT}
ENV APP_NAME=${APP_NAME}

WORKDIR /app

COPY ${APP_NAME} .

RUN ls -l

EXPOSE ${APP_PORT}
# EXPOSE ${WORKER_PORT}

# ENTRYPOINT ./${APP_NAME}

RUN echo ${APP_PORT}

RUN pwd

RUN chmod +x ${APP_NAME}
ENTRYPOINT [ "./go-autoscale" ]

# ENTRYPOINT ["/bin/bash", "-c", "./$APP_NAME \"$@\"", "--"]
CMD ["start-app"]