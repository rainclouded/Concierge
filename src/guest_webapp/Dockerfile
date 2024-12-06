FROM node:20-alpine AS build_react

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build_react /app/dist /usr/share/nginx/html
RUN chown -R appuser:appgroup /usr/share/nginx/html /etc/nginx/conf.d

EXPOSE 80
USER appuser

CMD ["nginx", "-g", "daemon off;"]