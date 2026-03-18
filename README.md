## 项目说明

本音乐网站的客户端和管理端使用 **Vue** 框架来实现，服务端使用 **Spring Boot + MyBatis** 来实现，数据库使用了 **MySQL**。

## 项目功能

- 音乐播放
- 用户登录注册
- 用户信息编辑、头像修改
- 歌曲、歌单搜索
- 歌单打分
- 歌单、歌曲评论
- 歌单列表、歌手列表分页显示
- 歌词同步显示
- 音乐收藏、下载、拖动控制、音量控制
- 后台对用户、歌曲、歌手、歌单信息的管理

<br/>

## 技术栈

### 后端

**SpringBoot + MyBatis + Redis** **+ minio**

### 部署

**docker**

### 前端

**Vue3.0 + TypeScript + Vue-Router + Vuex + Axios + ElementPlus + Echarts**

<br/>

## 开发环境

JDK： jdk-8u141

mysql：mysql-5.7.21-1-macos10.13-x86_64（或者更高版本）

redis：5.0.8 或 [docker启动redis | 想飞跃的鱼 (nanshaws.github.io)](https://nanshaws.github.io/docker/docker启动redis(完美过程).html)

node：14.17.3

IDE：IntelliJ IDEA 2018、VSCode

minio: 下载本地最新 或者 [docker完美启动minio | 想飞跃的鱼 (nanshaws.github.io)](https://nanshaws.github.io/docker/docker完美启动minio(完美过程).html)

<br/>

## 下载运行

### 1、下载项目到本地

```bash
git clone https://github.com/tiaodan/study-music-server-java.git

```



### 3、修改配置文件

1）创建数据库
将 `music-website/music-server/sql` 文件夹中的 `tp_music.sql` 文件导入数据库。

2）修改用户名密码
修改 `music-website/music-server/src/main/resources/application.properties` 文件里的 `spring.datasource.username` 和 `spring.datasource.password`；

### 4、启动项目

- **启动管理端**：进入 music-server 文件夹，运行下面命令启动服务器

```js
// 方法一
./mvnw spring-boot:run

// 方法二
mvn spring-boot:run // 前提装了 maven
```

- **启动 redis**：直接在终端输入下面命令

```
redis-server
```

> 下载地址：https://redis.io/
>


### 5、常见问题

1、图片、音乐加载失败
把 music-website/music-server 目录下的 img、song 目录移动到 music-website 目录。

2、音乐播放不了
可能是音乐损毁了，重新更换一下音乐资源。

<br/>

### 6、部署在linux上，用 docker【本地运行可以忽略】

将以下东西存储到Linux服务器上：

![image-20240108131746139](./img/image-20240108131746139.png)

还有编译好的jar包，都放到同一目录里面，如下：

![image-20240108131824788](./img/image-20240108131824788.png)

```
docker compose up --build
```

运行结果：

![image-20240108131927175](./img/image-20240108131927175.png)

