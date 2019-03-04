最近在折腾工单标题预测，模型离线训练使用tensorflow，在线预测使用tensorflow-serving，官方推荐使用docker。

# 为什么要使用docker

tensorflow-serving使用C++，有较复杂的编译依赖和运行时依赖。
Docker封装Linux Container，并将用程序与该程序的依赖，打包在一个文件里面。


# 如何使用

## image
Docker 把应用程序及其依赖打包在 image 文件里。
``` shell

# 列出本机的所有 image 文件。
$ docker image ls
REPOSITORY           TAG                 IMAGE ID            CREATED             SIZE
nginx                latest              8c9ca4d17702        3 days ago          109MB
tensorflow/serving   latest              38bee21b2ca0        5 days ago          229MB
hello-world          latest              fce289e99eb9        2 months ago        1.84kB

# 删除 image 文件
$ docker image rm [imageName]


# 从 docker hub 拉取image
$ docker image pull tensorflow/serving

```


## 容器
image 文件生成的容器实例，本身也是一个文件，称为容器。容器即可理解为Linux Container，包含运行时环境。

```
# 列出本机正在运行的容器
$ docker container ls
CONTAINER ID        IMAGE                COMMAND                  CREATED             STATUS              PORTS                              NAMES
9d52670ee76a        tensorflow/serving   "/usr/bin/tf_serving…"   About an hour ago   Up About an hour    8500/tcp, 0.0.0.0:8501->8501/tcp   musing_turing

# 删除容器
$ docker container rm [containerID]

```

## 运行

docker container run命令会从image文件，生成一个正在运行的容器实例。docker container run命令具有自动抓取image文件的功能。


```
# 运行tensorflow serving

$ docker run -t --rm -p 8501:8501  -v  "$TESTDATA/saved_model_half_plus_two_cpu:/models/half_plus_two"  -e MODEL_NAME=half_plus_two tensorflow/serving
```

