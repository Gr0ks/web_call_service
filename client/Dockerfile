# боевая среда
FROM nginx:stable-alpine
COPY ./dist /usr/share/nginx/html
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./cert.pem /etc/ssl/cert.pem
COPY ./key.pem /etc/ssl/key.pem
EXPOSE 443
CMD ["nginx", "-g", "daemon off;"]