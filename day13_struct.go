package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

type BookLang struct {
}

//结构体实例化
func (s BookLang) String() string {
	return "data"
}

type Human struct {
	name   string
	age    int
	weight int
}

/**
匿名字段实际上就是字段的继承
若存在匿名字段中的字段与非匿名字段名字相同，则最外层的优先访问，就近原则

所有的内置类型和自定义类型都是可以作为匿名字段的
*/
type Student struct {
	Human      //匿名字段，那么默认 Student 就包含了 Human 的所有字段
	speciality string
}

/**
嵌套结构体
	一个结构体可能包含一个字段，而这个字段反过来就是一个结构体
*/
type Address struct {
	city, state string
}

//如果结构体类型以大写字母开头，那么它是一个导出类型，可以从其他包访问。类似的，如果结构体的字段以大写开头，则可以从其他包访问它们
type Person struct {
	name string
	age  int
	//address Address
	Address // 在结构体中属于匿名结构体的字段成为提升字段，因为可以被访问，就好像它们属于拥有匿名结构字段的结构一样
}

//结构体是值类型，如果每个字段具有可比性，则是可比较的，如果它们对应的字段相等，则认为两个结构体变量是相等的

func main() {
	/**
	make 用于内建类型（map、slice 和 channel）的内存分配，返回初始化后的（非零）值
	new 用于各种类型的内存分配，分配零值填充的 T 类型的内存空间，并且返回其地址，即一个 *T 类型的值
	 */

	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.bookId)

	printBook(&Book1)

	printBook(&Book2)

	fmt.Printf("%v\n", BookLang{})

	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)

	var p Person
	p.name = "Naveen"
	p.age = 50
	//p.address = Address{
	//	city:  "Chicago",
	//	state: "Illinois",
	//}
	p.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	//fmt.Println("City: ", p.address.city)
	//fmt.Println("State: ", p.address.state)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}

func printBook(book *Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.bookId)
}
