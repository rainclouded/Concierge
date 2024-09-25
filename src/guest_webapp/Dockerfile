# choose node
FROM node:20-alpine AS build_react

# set working directory inside the container
WORKDIR /guest_webapp

# copy package.json & package-lock.json
COPY package*.json ./
# install dependencies/packages
RUN npm install

# copy all the project files
COPY . .

# react default expose port
EXPOSE 5173

# build our webapp
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]