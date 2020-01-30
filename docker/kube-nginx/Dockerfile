FROM centos:7
LABEL maintainer=caoyingjun name=kube-nginx

COPY kube-nginx /usr/kube-nginx
COPY kube_start /kube_start 

RUN chmod +x /kube_start

CMD ["/kube_start"]
