package utils

import (
	"attendance_uniapp/repositories"
	"fmt"
	"strings"
)

// SummaryToComponents 将 Summary 切片转化为 wxml 表格组件的字符串
func SummaryToComponents(summaries []repositories.Summary) string {
	var builder strings.Builder

	// 表头
	builder.WriteString("<view class=\"table\">\n")
	builder.WriteString("<view class=\"table-header\">\n")
	builder.WriteString("<text>学号</text>\n")
	builder.WriteString("<text>姓名</text>\n")
	builder.WriteString("<text>统计情况</text>\n")
	builder.WriteString("</view>\n")

	// 遍历每个 Summary，生成表格行
	for _, summary := range summaries {
		builder.WriteString("<view class=\"table-row\">\n")
		builder.WriteString(fmt.Sprintf("<text>%s</text>\n", summary.StudentId))
		builder.WriteString(fmt.Sprintf("<text>%s</text>\n", summary.StudentName))
		builder.WriteString(fmt.Sprintf("<text>%s</text>\n", summary.Statistics))
		builder.WriteString("</view>\n")
	}

	// 关闭表格容器
	builder.WriteString("</view>\n")

	// 返回最终的 wxml 字符串
	return builder.String()
}
