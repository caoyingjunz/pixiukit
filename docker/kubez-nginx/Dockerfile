FROM centos:7
LABEL maintainer=caoyingjun name=kubez-nginx

COPY kubez-nginx /kubez-nginx
COPY kubez_start /kubez_start

RUN chmod +x /kubez_start

USER root:root

CMD ["/kubez_start"]
