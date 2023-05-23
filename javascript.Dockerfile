FROM node

ARG FILE=app.js

COPY ${FILE} app.js

CMD ["node", "app.js"]