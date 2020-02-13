FROM centos:7 
LABEL maintainer=caoyingjun name=kube-toolbox

RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py \
    && python get-pip.py \
    && rm get-pip.py \
    pip install docker kubernetes

COPY kube_start /kube_start
RUN chmod +x /kube_start

CMD ["/kube_start"]
