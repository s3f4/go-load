FROM nginx:alpine as dev
COPY nginx-dev.conf /etc/nginx/nginx.conf
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait
ENTRYPOINT  /wait && nginx -g "daemon off;"

FROM nginx:alpine as prod
COPY nginx.conf /etc/nginx/nginx.conf