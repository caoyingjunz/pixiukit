FROM centos:7
LABEL maintainer=caoyingjun name=helm

COPY helm /usr/bin/helm 
COPY cmd_start /cmd_start

RUN chmod +x /cmd_start

CMD ["/cmd_start"]
