# elasticproxy
elasticproxy 代理elasticsearch 提供给kibana 调用， 可以在该项目中实现传输数据的修改，例如修改请求报文以及响应报文等

## 已实现功能
- 修改请求报文支持kibana基于filebeat的offset字段做二级排序

## 使用介绍
> 前提条件 已经有elastcsearch和kibana服务， 并且日志收集器为filebeat且日志数据中有offset字段
> - 在release中下载elasticproxy可执行文件 
> - 启动elasticproxy 通过` -elastic_host` 启动参数指定elasticsearch的地址, 默认elasticproxy的端口为8899
> - 启动kibana修改kibana的配置文件中elasticsearch的地址改为elasticproxy的地址

