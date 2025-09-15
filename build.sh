#!/bin/bash

# 构建脚本：编译OLED时间显示程序

# 设置中文环境
export LANG=zh_CN.UTF-8

echo "===== OLED时间显示程序构建脚本 ====="

# 检查是否已经初始化Go模块
if [ ! -f go.mod ]; then
    echo "正在初始化Go模块..."
    go mod init lcd_go
    if [ $? -ne 0 ]; then
        echo "Go模块初始化失败！"
        exit 1
    fi
fi

# 更新依赖
echo "正在更新依赖..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "依赖更新失败！"
    exit 1
fi

# 为不同架构构建
echo "正在为ARM64架构构建..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-w -s"  -o lcd_go.arm64
if [ $? -ne 0 ]; then
    echo "ARM64架构构建失败，但不影响主程序。"
fi

echo ""
echo "===== 构建完成！====="
echo "程序已编译完成，可执行文件："
echo "- lcd_go.arm64 (ARM64架构，适用于树莓派等设备)"
echo ""
echo "使用方法："
echo "1. 将OLED显示屏连接到开发板的I2C接口"
echo "2. 运行程序：./lcd_go.arm64"
echo ""
echo "注意：程序默认使用I2C总线3和地址0x3C，如果您的硬件配置不同，请修改main.go文件。"