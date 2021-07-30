package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 100
	var num2 float64 = 23.456
	var b bool = true
	var charts byte = 'a'
	var str string
	//法1：fmt.Sprintf()
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str typp=%T values=%v\n",str,str)
	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str typp=%T values=%v\n",str,str)
	str = fmt.Sprintf("%t",b)
	fmt.Printf("str typp=%T values=%q\n",str,str)
	str = fmt.Sprintf("%c",charts)
	fmt.Printf("str typp=%T values=%q\n",str,str)
	//法2：strconv包的函数
	str = strconv.FormatInt(int64(num1),2)
	//FormatInt()函数返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	fmt.Printf("str typp=%T values=%q\n",str,str)

	str = strconv.FormatBool(b)
	fmt.Printf("str typp=%T values=%q\n",str,str)
}
