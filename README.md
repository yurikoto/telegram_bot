# Yurikoto Telegram机器人

> 本项目fork自[hitokoto-osc/telegram_bot](https://github.com/hitokoto-osc/telegram_bot) ，使用或自部署本代码请严格遵守MIT协议。本仓库代码可供学习使用，但Yurikoto提供的壁纸、台词资源严禁商用

[![Go Report Card](https://goreportcard.com/badge/github.com/yurikoto/yurikoto-telegram-bot)](https://goreportcard.com/report/github.com/yurikoto/yurikoto-telegram-bot) [![Maintainability](https://api.codeclimate.com/v1/badges/293b8754a06684395ea0/maintainability)](https://codeclimate.com/github/yurikoto/yurikoto-telegram-bot/maintainability)

## 须知
如果您需要自部署本代码，请提交issue并提供您的服务器ip地址。审核通过后，我们将把您的服务器ip加入白名单。不在白名单内的ip会受到严格的频次控制（目前是每分钟5次）。

## 简介

Yurikoto Telegram机器人，fork自[hitokoto-osc/telegram_bot](https://github.com/hitokoto-osc/telegram_bot) ，如果您有任何建议或改进想法，欢迎提交issue或pr。 

[Yurikoto主页](https://yurikoto.com) [Yurikoto机器人](https://t.me/yurikoto_bot)

## TODO
telebot-v3更新后封装频次控制为中间件

## 自部署指南

### 环境

以下为Yurikoto官方使用的环境
CentOS 8.2
Go 1.15.7 （服务端可不需要）

### 网络

为了与Telegram服务器保持通讯，请将本项目部署在境外。

### 配置文件

修改`config.template.yml`并重命名为`config.yml`

### 编译

```shell
go mod download
SET GOOS=linux
SET GOARCH=amd64
go build
```

执行上述命令后项目目录会出现可执行文件`yurikoto-telegram-bot`

### 运行

将可执行文件与`config.yml`上传到服务器。在`lib/systemd/system`下创建`yurikoto-telegram-bot.service`，内容如下：

```ini
[Unit]
Description=Yurikoto Telegram Bot Service
After=network.target

[Service]
Type=simple
Restart=on-failure
RestartSec=5s
ExecStart=/path/yurikoto-telegram-bot -c /path/


[Install]
WantedBy=multi-user.target
```

并将`path`替换为可执行文件所在目录（`ExecStart`末尾斜杠需保留）。执行：

```shell
systemctl enable yurikoto-telegram-bot
systemctl start yurikoto-telegram-bot
systemctl status yurikoto-telegram-bot
```

若显示"active"字样，则说明部署成功。