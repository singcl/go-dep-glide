## Go Dep Glide 包管理工具

dep: https://gocn.vip/article/414

glide: https://www.cnblogs.com/CraryPrimitiveMan/p/7828357.html

https://studygolang.com/articles/10523

## Notice

```sh
# 解决glide up时出现的类似的错误
# 在TianChao golang.org是被qiang了的，所以基本都会遇到这样的错误
[ERROR] Error looking for golang.org/x/crypto/acme/autocert: Cannot detect VCS
```

解决方案：https://my.oschina.net/u/553243/blog/1475626?p=2

### 基本思路：

1.  在 glide.yaml 中添加 package`- package: golang.org/x/crypto`
2.  设置镜像：`glide mirror set golang.org/x/crypto github.com/golang/crypto`
3.  `glide up`
