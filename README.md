# gin-flipped-api

<div align=center>
<img src="https://img.shields.io/badge/golang-1.18-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

本项目基于 [gin-vue-admin v2.5.0](https://github.com/flipped-aurora/gin-vue-admin/),感谢原框架作者们的付出

首次使用,安装数据库,需要将根目录的main.go文件中的: //initialize.RegisterTables(global.GVA_DB) // 初始化admin库表,这一行取消注释,用来初始化数据库,数据库初始化完成后建议注释掉这一行,不然那每次启动都会遍历数据库表查询

api请求格式:
Content-Type:application/json

管理后台[gin-flip-admin](https://github.com/haoleiqin/gin-flip-admin/)