#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
将输入的96个16进制数据，按照每3个字节一组，合成为32组
"""

import sys

def group_hex_data(hex_data):
    """
    将输入的96个16进制数据按每3个字节一组进行分组
    :param hex_data: 包含96个16进制数据的列表
    :return: 包含32组数据的列表，每组3个字节
    """
    # 验证输入数据数量
    if len(hex_data) != 96:
        raise ValueError(f"输入数据数量错误，需要96个数据，实际收到{len(hex_data)}个数据")
    
    # 按每3个字节一组进行分组
    result = []
    for i in range(0, 96, 3):
        # 取出3个字节并合并
        group = hex_data[i:i+3]
        # 将每组数据转换为十六进制字符串格式输出，大写并添加0x前缀
        group_hex = '0x00' + ''.join([f'{byte:02X}' for byte in group])
        result.append(group_hex)
    
    return result

def main():
    # 示例用法
    print("本脚本用于将96个16进制数据按每3个字节一组合成为32组")
    print("请输入96个十六进制数据，以空格分隔（例如：01 02 03 ... ff）")
    
    try:
        # 从用户输入获取数据
        user_input = input().strip()
        # 分割输入并转换为整数
        hex_values = []
        
        # 尝试处理用户输入
        if user_input:
            # 先处理逗号分隔的情况
            if ',' in user_input:
                parts = [p.strip() for p in user_input.split(',')]
            else:
                # 再处理空格分隔的情况
                parts = user_input.split()
            
            # 转换为整数，处理可能的0x前缀
            hex_values = []
            for part in parts:
                # 移除可能的0x前缀
                if part.startswith('0x'):
                    part = part[2:]
                # 转换为整数
                hex_values.append(int(part, 16))
        
        # 如果用户没有输入数据，使用示例数据
        if not hex_values:
            print("未输入数据，使用示例数据进行演示...")
            # 创建示例数据 (0x00 到 0x5F，共96个数据)
            hex_values = list(range(96))
        
        # 处理数据
        groups = group_hex_data(hex_values)
        
        # 输出结果
        print(f"\n处理结果（共{len(groups)}组）：")
        for i, group in enumerate(groups):
            print(f"{group}")
            
    except ValueError as e:
        print(f"错误: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"发生未知错误: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()