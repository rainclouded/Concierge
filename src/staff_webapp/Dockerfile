FROM node:20-alpine AS build_angular


WORKDIR /app

COPY package*.json ./

RUN npm install -g @angular/cli
RUN npm install

COPY . .

RUN npm run build --prod

FROM nginx:alpine

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build_angular /app/dist/staff_webapp/browser /usr/share/nginx/html
RUN chown -R appuser:appgroup /usr/share/nginx/html


EXPOSE 80
USER appuser

CMD ["nginx", "-g", "daemon off;"]
