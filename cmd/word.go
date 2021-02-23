package cmd

import (

	"github.com/20100204/gotravel/internal/word"
	cobra "github.com/spf13/cobra"
	"log"
	"strings"
)
const(
MODE_UPPER  =  iota +1   // 单词转为大写
MODE_LOWER
MODE_UNDERSCORE_TO_UPPER_CAMELCASE
MODE_UNDERSCORE_TO_LOWER_CAMELCASE
MODE_CAMELCASE_TO_UNDERSCORE   //驼峰转下划线
)
var desc = strings.Join([]string{
	"该命令支持各种单词格式转换,模式如下：",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 写划线转为大驼峰",
	"4: 下划线转为小驼峰",
	"5: 驼峰转为下划线",
},"\n")
var str string
var mode int8
var wordCmd = &cobra.Command{
	Use:                        "word",
	Short:                      " 单词格式转换",
	Long:                       "支持多种单词格式转换", 
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)

		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UpderscoreToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCameCase(str)
		default:
			log.Fatalf("暂不支持改格式转换")
		}
		log.Printf("输出结果：%s",content)
	},
 
	 
}



func init()  {
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入单词转换的格式")
}
