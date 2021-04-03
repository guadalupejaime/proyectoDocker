FROM node
WORKDIR /usr/src/app
COPY . ./
RUN npm install -g @angular/cli
RUN npm install
RUN ng build --prod

# FROM nginx:1.13
# WORKDIR /root/
# COPY --from=0 /usr/src/app/dist/containerE/ /usr/share/nginx/html/
# COPY ./nginx-custom.conf /etc/nginx/conf.d/default.conf
