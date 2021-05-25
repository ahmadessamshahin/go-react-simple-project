FROM node 


WORKDIR /code
COPY package.json .
COPY yarn.lock .
RUN yarn
COPY . .

CMD [ "yarn", "start" ]

