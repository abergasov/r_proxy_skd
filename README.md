## Description
Дублирующий, локальный прокси сервер для СКД с контроллерами IronLogic Z5R Web

Общение через протокол WebJSON

## Installation
```shell script
sudo nano /lib/systemd/system/skd.service
```
Пихаем сордержимое
```shell script
[Unit]
Description=skd_proxy

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/path_to_go_folder/r_proxy_skd/main

[Install]
WantedBy=multi-user.target
```
