package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"os/exec"
	"strings"
	"time"

	goi2coled "github.com/waxdred/go-i2c-oled"
	"github.com/waxdred/go-i2c-oled/ssd1306"
)

func main() {
	// 初始化OLED显示屏，参数说明：
	// - 电源模式：ssd1306.SSD1306_SWITCHCAPVCC
	// - 高度：32
	// - 宽度：128
	// - I2C地址：0x3C（大多数OLED模块的默认地址）
	// - I2C总线：3（用户指定的I2C总线）
	oled, err := goi2coled.NewI2c(ssd1306.SSD1306_SWITCHCAPVCC, 32, 128, 0x3C, 3)
	if err != nil {
		log.Fatalf("无法初始化OLED显示屏: %v", err)
	}
	defer oled.Close()

	// 定义颜色
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	var ori_timeStr string
	// 持续更新显示时间
	for {
		// 清除OLED图像（设为全黑）
		draw.Draw(oled.Img, oled.Img.Bounds(), &image.Uniform{black}, image.Point{}, draw.Src)

		// 使用 date 命令获取时间
		dateCmd := exec.Command("date", "+%H:%M")
		output, err := dateCmd.Output()
		if err != nil {
			log.Printf("获取时间错误: %v", err)
			continue
		}
		timeStr := strings.TrimSpace(string(output))
		log.Printf("当前时间: %s", timeStr)

		if ori_timeStr != timeStr {
			ori_timeStr = timeStr

			// 为了让时间显示铺满整个128*32的显示屏，我们需要使用更大的字体效果
			// 由于basicfont包中没有更大的内置字体，我们可以通过多次绘制相同字符并调整位置来模拟更大的字体
			drawLargeTime(oled.Img.(*image.RGBA), timeStr, white)
			// _ = drawLargeTime // 防止编译器警告

			// 清除OLED缓冲区并更新
			oled.Clear()
			oled.Draw()

			// 在OLED屏幕上显示内容
			err = oled.Display()
			if err != nil {
				log.Printf("显示错误: %v", err)
			}
		}
		// 每十秒更新一次显示
		time.Sleep(time.Second * 10)
	}
}

// drawChar 绘制一个24x32像素的字符
// 每个字符由96字节的数组定义（32行，每行3字节，共32*3=96字节）
// 为了方便在Go中处理，我们使用32个uint32值来表示，每个uint32的低24位有效
func drawChar(img *image.RGBA, char rune, x, y int, color color.RGBA) {
	const width = 24
	const height = 32
	var pixelMap [32]uint32

	// 根据字符选择对应的像素映射
	switch char {
	case '0':
		// 数字0的像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x0000FF00,
			0x0001C180,
			0x000700E0,
			0x000F00E0,
			0x000E0070,
			0x001E0078,
			0x001C0038,
			0x001C0038,
			0x003C0038,
			0x003C003C,
			0x003C003C,
			0x003C003C,
			0x003C0038,
			0x003C0038,
			0x001C0038,
			0x001C0078,
			0x001E0070,
			0x000E0070,
			0x000700E0,
			0x000381C0,
			0x0001C380,
			0x00003C00,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '1':
		// 数字1的像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000C00,
			0x00003C00,
			0x0007FC00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00001C00,
			0x00007E00,
			0x0007FFE0,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '2':
		// 数字2的像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x0000FF80,
			0x000701E0,
			0x000C0070,
			0x001C0070,
			0x001C0078,
			0x001E0078,
			0x001E0070,
			0x00000070,
			0x000000E0,
			0x000001C0,
			0x00000380,
			0x00000700,
			0x00000C00,
			0x00003800,
			0x00006000,
			0x0001C000,
			0x00030008,
			0x00060018,
			0x000C0018,
			0x00180070,
			0x003FFFF0,
			0x003FFFF0,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '3':
		// 数字3的像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x0001FF00,
			0x000603C0,
			0x000C00E0,
			0x001C00E0,
			0x001E00F0,
			0x000E00F0,
			0x000000E0,
			0x000000E0,
			0x000001C0,
			0x00001E00,
			0x00007F00,
			0x000001C0,
			0x00000070,
			0x00000070,
			0x00000038,
			0x00000038,
			0x001E0038,
			0x001E0078,
			0x001E0070,
			0x000C00E0,
			0x00070380,
			0x0000FC00,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '4':
		// 数字4的简化版像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000180,
			0x00000380,
			0x00000780,
			0x00000F80,
			0x00001B80,
			0x00003380,
			0x00006380,
			0x0000C380,
			0x00018380,
			0x00030380,
			0x00060380,
			0x000C0380,
			0x00180380,
			0x00100380,
			0x003FFFFC,
			0x00000380,
			0x00000380,
			0x00000380,
			0x00000380,
			0x00000380,
			0x000007C0,
			0x00007FFC,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '5':
		// 数字5的简化版像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x0007FFF0,
			0x0007FFF0,
			0x00040000,
			0x00040000,
			0x00040000,
			0x00040000,
			0x00040000,
			0x000C0000,
			0x000CFF80,
			0x000D81E0,
			0x000E00F0,
			0x000C0070,
			0x00000038,
			0x00000038,
			0x00000038,
			0x000C0038,
			0x001E0038,
			0x001E0070,
			0x001C0070,
			0x000C00E0,
			0x000303C0,
			0x0000FE00,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '6':
		// 数字6的简化版像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00003FC0,
			0x0001C0E0,
			0x000300F0,
			0x000600F0,
			0x000E0000,
			0x000C0000,
			0x001C0000,
			0x001C0000,
			0x001C7F80,
			0x003DC1E0,
			0x003F0070,
			0x003E0038,
			0x003C0038,
			0x003C003C,
			0x003C003C,
			0x001C0038,
			0x001C0038,
			0x000E0038,
			0x000F0030,
			0x00078060,
			0x0001C1C0,
			0x00003E00,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '7':
		// 数字7的简化版像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x000FFFF8,
			0x000FFFF8,
			0x000E0030,
			0x00180060,
			0x00180040,
			0x00000080,
			0x00000180,
			0x00000300,
			0x00000600,
			0x00000600,
			0x00000C00,
			0x00001C00,
			0x00001800,
			0x00003800,
			0x00003800,
			0x00007800,
			0x00007800,
			0x00007800,
			0x00007800,
			0x0000F800,
			0x00007800,
			0x00007000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '8':
		// 数字8的简化版像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x0001FF00,
			0x000700E0,
			0x000C0070,
			0x001C0030,
			0x00180038,
			0x001C0038,
			0x001C0030,
			0x000F0070,
			0x0007C0E0,
			0x0001FB80,
			0x0001FF00,
			0x00070FC0,
			0x000C01E0,
			0x001C00F0,
			0x00380078,
			0x00380038,
			0x00380038,
			0x00380038,
			0x00180030,
			0x000C0060,
			0x000381C0,
			0x00007E00,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case '9':
		// 数字9的简化版像素图案
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x0001FF00,
			0x000701C0,
			0x000E00E0,
			0x001C0070,
			0x001C0070,
			0x00380038,
			0x00380038,
			0x00380038,
			0x00380038,
			0x003C0078,
			0x001C00B8,
			0x000F0338,
			0x0007FE38,
			0x00000078,
			0x00000078,
			0x00000070,
			0x000000F0,
			0x000000E0,
			0x000F01C0,
			0x000F0380,
			0x00070E00,
			0x0001F800,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	case ':':
		// 冒号的像素图案（两个点）
		pixelMap = [32]uint32{
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00007C00,
			0x00007C00,
			0x00003C00,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00001800,
			0x00007C00,
			0x00007C00,
			0x00003800,
			0x00000000,
			0x00000000,
			0x00000000,
			0x00000000,
		}
	default:
		// 未知字符，不绘制
		return
	}

	// 绘制字符的像素点
	for row := 0; row < height && row < 32; row++ {
		for col := 0; col < width && col < 24; col++ {
			// 检查当前像素点是否应该绘制
			if (pixelMap[row] & (1 << (23 - col))) != 0 {
				// 确保坐标在图像范围内
				if x+col >= 0 && x+col < 128 && y+row >= 0 && y+row < 32 {
					img.Set(x+col, y+row, color)
				}
			}
		}
	}
}

// drawLargeTime 绘制大尺寸的时间字符串，使其铺满128*32的显示屏
func drawLargeTime(img *image.RGBA, timeStr string, color color.RGBA) {
	const charWidth = 24
	const charHeight = 32
	const charSpacing = 0 // 字符间距
	const totalChars = 5  // HH:MM 共5个字符
	const totalWidth = totalChars*charWidth + (totalChars-1)*charSpacing

	// 计算起始位置，使时间居中显示
	startX := (128 - totalWidth) / 2
	startY := 0 // 顶部对齐，因为字符高度是32，正好等于显示屏高度

	// 逐个绘制字符
	for i, char := range timeStr {
		x := startX + i*(charWidth+charSpacing)
		y := startY
		drawChar(img, char, x, y, color)
	}
}
