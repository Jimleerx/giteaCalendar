# giteaCalendar

## 简介

> 由于公司私仓从 GitLab 迁移到了 Gitea ，于是就有了这个项目
>
> GitLab Ver. : [https://github.com/luckykeeper/gitlabCalendar](https://github.com/luckykeeper/gitlabCalendar)

Gitea 私仓提交日历抓取和可视化，向大家展示你在公司的~~摸鱼~~工作情况吧ヾ(≧▽≦*)o

demo: [https://gitea-calendar.luckykeeper.site:44443/](https://gitea-calendar.luckykeeper.site:44443/)

鼠标放上去可以看到当天的 Contribution，数据在程序启动后每 5 分钟抓取一次

![demo](./images/demo.png)

## 使用

### 配置文件

下载 Release 二进制文件，在同级目录新建一个 `config.yaml` 文件，其模板如下

```ini
# DebugMode 调试模式: false 关闭，true 开启，默认：false
DebugMode: false
# LogLevel 日志输出等级: Debug: 5 | Info: 4 | Warn: 3 | Error: 2 | Fatal: 1 , 默认： 4
LogLevel: 4
# ApiPort 服务监听端口: 默认： 44443
ApiPort: 44443
# DataBaseType 数据库类型: sqlite | mysql | postgres ，默认：sqlite3。一般不需要使用非 sqlite 数据库
DataBaseType: "sqlite"
# DatabaseDSN 数据库连接字符串，各数据库示例
## sqlite3: "./giteaCalendar.db"
## mysql: "user:password@tcp(192.168.1.1:3306)/gitea?charset=utf8mb4"
## postgres: "postgres://user:password@192.168.1.1:5432/postgres?sslmode=disable"
DatabaseDSN: "./giteaCalendar.db"

# Server gitea 服务器地址: 你的 gitea 私仓的地址，带协议，结尾不带 “/”
Server: "https://git.example.com:9000"
# UserName 用户名：你在 Gitea 内的用户名
UserName: ""
# ApiKey 访问密钥：在【设置】->【用户设置】->【应用】->【管理 AccessToken】下生成的访问令牌
ApiKey: ""
```

### 启动

```bash
# cd 到程序目录
# 由于 go 打包出来的是一个二进制程序，不需要运行环境，直接运行即可
# ./ 加二进制文件名，linux x64 示例如下:
./giteaCalendar_linux_amd64
```

## 项目路由

- `/` Web 页面，提交日历可视化
- `/metrics` 服务器指标监控
- `/checkPoint/ready` 就绪探针，没有特殊逻辑直接返回 200
- `/checkPoint/liveness` 存活探针，没有特殊逻辑直接返回 200
- `/api/Calendar` 前端获取提交的 API 接口

# 从作者的 [gitlabCalendar]((https://github.com/luckykeeper/gitlabCalendar)) 迁移

如果你也和作者一样，从 GitLab 迁移到了 Gitea，可以保留原先的数据，无缝迁移过来，方法如下：

1. 找到 gitlabCalendar 的数据库，即其目录下的【GitLabCalendar.db】文件
2. 将该文件复制到 giteaCalendar 的运行目录下，并重命名为【giteaCalendar.db】
3. 运行项目，可以看到之前的数据保留

> ⚠注意：如果 gitea 在相同日期存在数据，将会覆盖原 gitlabCalendar 写入的数据

# 姊妹项目

[https://github.com/luckykeeper/gitlabCalendar](https://github.com/luckykeeper/gitlabCalendar) ASP.NET Core Web API 练手项目：GitLab私仓提交日历抓取和可视化，向大家展示你在公司的摸鱼（划掉）工作情况吧ヾ(≧▽≦*)o

# Special  Thanks  To

[scriptex/gitlab-calendar: Embed your Gitlab calendar everywhere. (github.com)](https://github.com/scriptex/gitlab-calendar) 前端文件基于此微有改动
