FROM nginx:alpine
LABEL MAINTAINER="liwei1dao@gamil.com"
COPY /nginx.conf.tpl /etc/nginx/nginx.conf.tpl
#copy web
COPY  /dist/ /var/www/
EXPOSE 80