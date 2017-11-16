## 概述

SDN 作为 Kubernetes 的容器网络解决方案，需提供网络连通、租户隔离、防火墙、负载均衡等功能，并保障网络高可用、可扩展和安全性。

### 连通性

1, 同一 Kubernetes namespace 下的容器可以不经过 NAT 实现网络互通  
2, 物理网络（宿主机所在的网络）与任意 namespace 下的容器可以不经过 NAT 实现网络互通  
3, 所有容器可以访问外网  
4, 所有容器可以通过 kube-dns 查询集群内部的 service 域名  

### 租户隔离

不同 namespace 下的容器网络不互通

### 防火墙

1, 具有外网IP的物理机，仅对外开放有限端口（80/443）  
2, 所有物理机仅对容器开放有限端口，例如 kube-dns TCP/UDP 53 端口  

### 负载均衡

支持 type=LoadBalancer 的 Kubernetes service，参考 https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer

### 高可用

1, 某个物理机宕机，不会导致整个集群的网络故障，部分受影响的网络流量能够在较短时间内恢复  
2, 网络配置出现异常时（例如防火墙规则、容器网卡、路由表丢失），系统能够自动检测并在较短时间内修复  
3, 网络综合可用性 99.9% 以上  

### 可扩展

集群内的网络流量过大时，可以通过水平扩展服务器数量来缓解压力  

### 性能

容器之间的流量带宽可以达到物理网卡的 90% 以上，延迟在 1ms 以内。

### 安全性

集群内部能够应对 ARP Spoofing 和 Syn Flood 等攻击
