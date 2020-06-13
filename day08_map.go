package main

import "fmt"

func main() {
	/**
	Map
		hash 实现，也是引用类型
	1.map 是无序的，每次打印出来的 map 都回不一样，必须通过 key 获取
	2.长度不固定，和 slice 一样，是引用类型
	3.内置的 len 函数同样适用于 map，返回 map 拥有的 key 的数量
	4.key 可以是所有可比较的类型，如布尔型、整数型、浮点型、复杂型、字符串型..也可以是键
	*/

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Println(rating)

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	//插入 key-value 对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	//使用 key 输出 map 值
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["United States"]
	// 如果 ok 是 true，则存在
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of united States is not present")
	}

	//delete(map, key)函数用于删除集合的元素，参数为 map 和其对应的 key，删除函数不返回任何值
	fmt.Println("原始 map")

	//打印 map
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	//删除元素
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("删除元素后 map")

	//打印 map
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)

	personSalary := map[string]int{"steve": 12000, "jamie": 15000}
	personSalary["mike"] = 9000
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
}
