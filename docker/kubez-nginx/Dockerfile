FROM centos:7
LABEL maintainer=caoyingjun name=kubez-nginx

COPY kubez-nginx /kube-nginx
COPY kubez_start /kube_start

RUN chmod +x /kubez_start

CMD ["/kubez_start"]
