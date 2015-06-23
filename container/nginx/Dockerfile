FROM nginx:latest
MAINTAINER takasing

ADD nginx.conf /etc/nginx/nginx.conf
RUN rm /etc/nginx/conf.d/default.conf
ADD conf.d/golang-app.conf /etc/nginx/conf.d/golang-app.conf
