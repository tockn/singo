FROM node:20.15.1-alpine3.20

RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN rm -rf /app/node_modules

EXPOSE 8080
CMD [ "http-server", "dist" ]
