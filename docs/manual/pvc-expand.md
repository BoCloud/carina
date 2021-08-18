#### 磁盘扩容

对于存储来说随着使用总会有扩容的需求，carina支持在线扩容，如下所示

```shell
$ kubectl get pvc -n carina
NAMESPACE  NAME        STATUS  VOLUME                                    Capacity  STORAGECLASS  AGE
carina     carina-pvc  Bound   pvc-80ede42a-90c3-4488-b3ca-85dbb8cd6c22  7G        carina-sc     20d

  
```

进行在线扩容

```shell
$ kubectl patch pvc/carina-pvc \
  --namespace "carina" \
  --patch '{"spec": {"resources": {"requests": {"storage": "15Gi"}}}}'
  
```

进入容器查看容量

```shell
$ kubectl exec -it web-server -n carina bash
$ df -h
Filesystem                                 Size  Used Avail Use% Mounted on
overlay                                    199G   17G  183G   9% /
tmpfs                                      64M     0   64M   0% /dev
/dev/vda2                                  199G   17G  183G   9% /conf
/dev/carina-vg-hdd/volume....              15G     0   64M   0% /www/nginx/work
tmpfs                                      3.9G     0  3.9G   0% /tmp/k8s-webhook-server/serving-certs
```

#### 注意事项

* 如果创建的磁盘使用了缓存盘即bcache，由于受到bcache底层技术限制设备扩容后需要容器重新启动新的设备容量才会生效