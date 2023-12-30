FROM node:lts-alpine as ui-build-stage

WORKDIR /app

# COPY /ui/package*.json ./
COPY ./ui ./
RUN yarn install --frozen-lockfile

RUN yarn build

FROM scratch

ENV GID 1000
ENV UID 1000

COPY --from=ui-build-stage --chown=${UID}:${GID} /app/dist /www/
COPY --from=ui-build-stage --chown=${UID}:${GID} /app/src/assets/images /www/assets/images
COPY golink /

ENTRYPOINT ["/golink"]