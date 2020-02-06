FROM alpine:edge
LABEL maintainer=caoyingjun name=helm-toolbox

COPY helm /usr/bin/helm 
COPY cmd_start /cmd_start

RUN chmod +x /cmd_start

CMD ["/cmd_start"]
