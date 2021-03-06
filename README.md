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
2.  设置镜像：`glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git`
3.  `glide up`

在~/.glide/mirrors.yaml 中保存了我们设置的所以镜像资源。可以直接查看或者修改。

## Install

#### 方法一：

```sh
# 如果失败请自己去github下载镜像：https://blog.csdn.net/AlexWoo0501/article/details/73409917
go get -u github.com/singcl/go-dep-glide
```

#### 方法二：

```sh
# 第一步：
git clone https://github.com/singcl/go-dep-glide.git

# 第二步安装依赖:
glide up
# 或者
dep ensure

# 第三步成功：
OK
```
