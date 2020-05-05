FROM node:lts-alpine

RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN rm -rf /app/node_modules

EXPOSE 8080
CMD [ "http-server", "dist" ]
