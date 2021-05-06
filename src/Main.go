package main

import (
	"fmt"
	"module"
)

//func function_name( [parameter list] ) [return_types] {
//	函数体
//}
func testConst()(int,int,string){
	a , b , c := 1 , 2 , "str"
	return a,b,c
}
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Google" + "Play")
	fmt.Println(module.Add(1, 2))

	const age, index int = 12, 13
	var flag bool
	var flagAuto = "auto_type_value"
	intVal := 99
	name1, name2, name3 := 22, 23, 24
	fmt.Println(flag, flagAuto, age, index, intVal)
	fmt.Println(name1, name2, name3)
	fmt.Println(testConst())

	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
		n=1<<iota
		m=3<<iota
		k
		l
	)
	//0 1 2 ha ha 100 100 7 8 512 3072 6144 12288
	fmt.Println(a,b,c,d,e,f,g,h,i,n,m,k,l)
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n" )
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n" )
	}

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10{
		sum += sum
	}
	fmt.Println(sum)

	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}


	//声明数组
	//var variable_name [SIZE] variable_type
	//  将索引为 1 和 3 的元素初始化
	balance := [5]float32{1:2.57,3:7.13}
	fmt.Printf("balance=%f\n", balance)

	//指针
	aa := 20       /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &aa       /* 指针变量的存储地址 */
	if ip != nil { /* ptr 不是空指针 */

	}
	if ip == nil { /* ptr 是空指针 */

	}
	fmt.Printf("a 变量的地址是: %x\n", &aa  )
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	type Books struct {
		title   string
		author  string
		subject string
		bookId  int
	}
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "tomagoyaky", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "tomagoyaky", subject: "Go 语言教程", bookId: 649540})

	//https://www.runoob.com/go/go-slice.html
	//sa := arr[startIndex:endIndex]
	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
	var sa = make([]int,3,5)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sa),cap(sa),sa)


	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}


	/* 创建map */
	ccm := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	/* 打印地图 */
	for country := range ccm {
		fmt.Println(country, "首都是", ccm [ country ])
	}
	/*删除元素*/ delete(ccm, "France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除元素后地图")
	/*打印地图*/
	for country := range ccm {
		fmt.Println(country, "首都是", ccm [ country ])
	}
}
