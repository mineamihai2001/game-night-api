ARG BASE_IMAGE=golang:latest
ARG DEPENDENCIES_IMAGE=node:latest


FROM $DEPENDENCIES_IMAGE as dependencies
RUN npm i -g nodemon


FROM $BASE_IMAGE as base
COPY --from=dependencies / /
WORKDIR /app
COPY . .
CMD [ "nodemon", "--signal", "SIGTERM" ]
