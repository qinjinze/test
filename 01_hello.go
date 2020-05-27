//1）go语言以包做为管理单位
//2）每个文件必须先声明包
//3）程序必须有一个main包（重要）

package main

import "fmt"

type Human struct {
	Name string
}

var people = Human{Name: "qinjin"}

//入口行数
func main() { //左括号必须和函数同行
	//打印
	//打印“hello.go”打印到屏幕，Println()会自动换行
	//调用函数大部分需要导入包
	/*
		这也是注释，这是块注释
	*/

	fmt.Println("hello！ 秦晋，我是你开发的得力助手！\n\n\n") //go语言语句结尾没有分号
	//1、各种打印函数的比较
	fmt.Println("1、各种打印函数的比较")
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
	// ab1 2 3cd        //非字符串参数之间会添加空格
	// a b 1 2 3 c d   // 所有参数之间会添加空格
	// ab 1 2 3 cd    // Printf 将参数列表 a 填写到格式字符串 format 的占位符中，填写后的结果写入到标准输出中。

	//2、普通占位符占位符的使用
	c := 30
	fmt.Println("2、普通占位符占位符的使用\n")
	fmt.Printf("%v\n", people)  //%v  值和值对应类型一起打印出来,其中\n表示换行
	fmt.Printf("%+v\n", people) //%+v 打印结构体时，会添加字段名
	fmt.Printf("%#v\n", people) //%#v 自动匹配格式输出
	fmt.Printf("%T\n", people)  //%T  自动匹配变量输出：对应类型
	fmt.Printf("c type is %T\n", c)
	fmt.Printf("%%\n\n") //字面上的百分号，并非值的占位符

	//3、布尔占位符
	fmt.Println("3、布尔占位符\n")
	fmt.Printf("%t,\t%t\n", true, false) //%t true 或 false。\t输出tab的一个空格
	//4、整数占位符
	fmt.Println("4、整数占位符\n")
	fmt.Printf("%b %b\n", 5, c)        //%b  二进制表示
	fmt.Printf("%c\n\n", 0x4E2D)       //%c  相应Unicode码点所表示的字符
	fmt.Printf("%d\n", 0x12)           //%d  十进制表示
	fmt.Printf("%f\n", 10.123456789)   //%f  有小数点而无指数，例如 123.456
	fmt.Printf("%s\n", []byte("Go语言")) //%s 输出字符串表示（string类型或[]byte)
	fmt.Printf("%q\n\n", "Go语言")       //%q 双引号围绕的字符串，由Go语法安全地转义

	//5、格式化输入
	//1) 格式化输入：从输入端读取字符串（以空白分隔的值的序列），
	//2) 并解析为具体的值存入相应的 arg 中，arg 必须是变量地址。
	//3) 字符串中的连续空白视为单个空白，换行符根据不同情况处理。
	//4) \r\n 被当做 \n 处理。

	// 对于 Scan 而言，回车视为空白
	fmt.Println("5、格式化输入\n")
	a1, b1, c1 := "", 0, false
	fmt.Scan(&a1, &b1, &c1)
	fmt.Println(a1, b1, c1)
	// 在终端执行后，输入 abc 1 回车 t 回车
	// 结果 abc 1 true
	fmt.Println("----------------")
	// 整型变量也可以解析八进制和十六进制
	fmt.Scanln(&a1, &b1, &c1)
	fmt.Println(a1, b1, c1)
	// 在终端执行后，输入 def 0x20 T 回车
	// 结果 def 32 true

	// 格式字符串可以指定宽度
	fmt.Println("----------------")
	fmt.Scanf("%4s%d%t", &a1, &b1, &c1)
	fmt.Println(a1, b1, c1)
	// 在终端执行后，输入 1234567true 回车
	// 结果 1234 567 true

}
