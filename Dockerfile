FROM node:lts-alpine as ui-build-stage

WORKDIR /app

COPY ./ui ./
RUN yarn install --frozen-lockfile

RUN yarn build

FROM alpine

ENV GID 1000
ENV UID 1000

# Curl is used for healthcheck
RUN apk add curl

COPY --from=ui-build-stage --chown=${UID}:${GID} /app/dist /www/
COPY --from=ui-build-stage --chown=${UID}:${GID} /app/src/assets/images /www/assets/images
COPY golink /usr/bin

ENTRYPOINT ["/usr/bin/golink"]