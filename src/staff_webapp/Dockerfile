FROM node:20-alpine AS build_angular

WORKDIR /app

COPY package*.json ./

RUN npm install -g @angular/cli
RUN npm install

COPY . .

RUN npm run build --prod

FROM nginx:alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build_angular /app/dist/staff_webapp/browser /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]