
# Kubez-nginx

- 创建 `kubez nginx` 系统目录
```
mkdir -p /var/lib/kubez-nginx
```

- 创建 `kubez nginx` 配置 `/var/lib/kubez-nginx/kubez-nginx.conf`

```
worker_processes 1;

events {
    worker_connections  1024;
}

stream {
    upstream backend {
        hash $remote_addr consistent;
        # 配置后端 ip:port
        server 10.10.33.31:30006  max_fails=3 fail_timeout=30s;
    }

    server {
        # vip 配置
        listen 10.10.33.32:80;
        proxy_connect_timeout 1s;
        proxy_pass backend;
    }
}
```

- 启动代理服务
```
docker run -d --name <container_name> --privileged=true -v /var/lib/kubez-nginx/:/etc/kubez-nginx/ --net host jacky06/kubez-nginx:v1.0.0
```

- 检查服务已经正常启动
```
# docker ps
CONTAINER ID        IMAGE                        COMMAND             CREATED             STATUS              PORTS               NAMES
349c7df1a78f        jacky06/kubez-nginx:v1.0.0   "/kubez_start"      20 minutes ago      Up 20 minutes                           container_name
```

- 更新 kubernetes 的 service 为 LoadBalancer 模式
```
go run main.go --externalip <external_ip> --kubeconfig /path/to/config --name <service_name> --namesapce <service_namespaces>
```

- 执行后检查
```
# kubectl get svc ingress-nginx -n kube-system
NAME            TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx   LoadBalancer   10.254.158.47   10.10.33.32   80:30006/TCP,443:30008/TCP   18d
```
