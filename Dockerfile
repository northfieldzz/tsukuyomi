FROM node:20-bullseye as base

RUN apt-get update && apt-get upgrade -y
WORKDIR /usr/src/app


FROM base as development


FROM base as builder

COPY . .
RUN yarn install
RUN yarn compile
ENTRYPOINT yarn start


FROM base as production

ENV fly_launch_runtime="nodejs"
ENV NODE_ENV production

WORKDIR /usr/src/app
COPY --from=builder /usr/src/app /usr/src/app
ENTRYPOINT yarn start

CMD yarn start