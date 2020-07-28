# pb

## memberpb
```shell

# 生成会员服务
protoc -I memberpb/ memberpb/member.proto --go_out=plugins=grpc:memberpb

```