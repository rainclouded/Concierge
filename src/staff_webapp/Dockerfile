# choose node
FROM node:20-alpine AS build_angular

# set working directory inside the container
WORKDIR /staff_webapp

# copy package.json & package-lock.json
COPY package*.json ./

# install dependencies/packages
RUN npm install -g @angular/cli
RUN npm install

# copy all the project files
COPY . .

# react default expose port
EXPOSE 4200

# build our webapp
CMD ["ng", "serve", "--host", "0.0.0.0"]