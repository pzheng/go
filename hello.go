package main
import (
    "flag"
	"fmt"
)
var skillParam = flag.String("skill", "", "skill to perform")
func swap(x,y string)(string,string){
	return y,x
}
func swapNew(x *int, y *int) {
	var temp int
	temp = *x    /* 保存 x 地址的值 */
	*x = *y      /* 将 y 赋值给 x */
	*y = temp    /* 将 temp 赋值给 y */
 }
 func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }
 // 字符串处理函数，传入字符串切片和处理链
func StringProccess(list []string, chain []func(string) string) {
    // 遍历每一个字符串
    for index, str := range list {
        // 第一个需要处理的字符串
        result := str
        // 遍历每一个处理链
        for _, proc := range chain {
            // 输入一个字符串进行处理，返回数据作为下一个处理链的输入。
            result = proc(result)
        }
        // 将结果放回切片
        list[index] = result
    }
}
func main(){
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
	  )
	  var j int = 4
	  var ptr *int
	  ptr = &j
	  println("a的值为", j);    // 4
	  println("*ptr为", *ptr);  // 4
	  println("ptr为", ptr);    // 824633794744
	fmt.Println(a,b,c,d,e,f,g,h,i)
	fmt.Println("hello world")
	fmt.Println("fas")
	fmt.Println("FFF")
	k,l := swap("tom","jack")
	fmt.Println(k,l)
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	sum := 0
        for i := 0; i <= 10; i++ {
                sum += i
        }
		fmt.Println(sum)
		m := 100
		n := 200
	swapNew(&m,&n)
	fmt.Println(m,n)
	var numbers = make([]int,3,5)
	cc :=[]int{1,2,3}
	fmt.Println(cc[:])
	printSlice(numbers)

	const elementCount = 10
	srcData := make([]int,elementCount)

	 // 将切片赋值
	 for i := 0; i < elementCount; i++ {
        srcData[i] = i
    }
    // 引用切片数据
	refData := srcData
	fmt.Println(refData)
    // 预分配足够多的元素切片
    copyData := make([]int, elementCount)
    // 将数据复制到新的切片空间中
    copy(copyData, srcData)
	var mapLit map[string]int
	var mapAssigned map[string]int
    mapLit = map[string]int{"one": 1, "two": 2}
    mapCreated := make(map[string]float32)
    mapAssigned = mapLit
    mapCreated["key1"] = 4.5
    mapCreated["key2"] = 3.14159
    mapAssigned["two"] = 3
    fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
    fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
    fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	
	// 待处理的字符串列表
    list := []string{
        "go scanner",
        "go parser",
        "go compiler",
        "go printer",
        "go formater",
	}
	chain :=[]func(string)string{
		// removePrefix,
		// strings.TrimSpace,
		// strings.ToUpper,
	}
	StringProccess(list,chain)
	for _,str := range list{
		fmt.Println(str)
	}

	flag.Parse()
    var skill = map[string]func(){
        "fire": func() {
            fmt.Println("chicken fire")
        },
        "run": func() {
            fmt.Println("soldier run")
        },
        "fly": func() {
            fmt.Println("angel fly")
        },
    }
    if f, ok := skill[*skillParam]; ok {
        f()
    } else {
        fmt.Println("skill not found")
	}
	
}