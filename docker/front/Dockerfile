ARG NODE_VERSION=18
FROM node:${NODE_VERSION}

WORKDIR /usr/src/chat-example

COPY ./front /usr/src/chat-example

RUN npm install -g bun
RUN bun install

CMD ["bun", "run", "dev"]
