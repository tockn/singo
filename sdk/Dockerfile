FROM node:10

WORKDIR /app
ADD ./package.json /app
ADD ./tsconfig.json /app
ADD ./webpack.config.js /app
ADD ./src/ /app/src/
ADD ./index.html /app

RUN yarn
CMD yarn start
