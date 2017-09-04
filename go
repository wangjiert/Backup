1.源文件中非注释的第一行指明这个文件属于那个包
2.package main表示一个可以独立执行的程序
3.main函数是每一个可执行程序所必须包含的，一般都是在启动后第一个执行的函数（如果有init函数则先执行这个函数）
4.标识符（常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，那么这个对象就可以被外部包的代码所使用，如果是小写字母开头，则对包外是不可见
5.一行代表一个语句的结束
6.变量的声明必须用空格隔开
7.声明变量使用var关键字 var identifier type
8.声明变量时可以指定变量类型，也可以根据值自行判断变量类型,或则省略var使用:=，如果左边的变量已经声明过，编译会报错
9.:=必须出现在函数中,不可以用于全局变量的声明和赋值
10.是否区分大小写
11.取址符合&,*指针变量
12.var(
	vname1 v_type1
	vname2 v_type2
)
	用于声明全局变量
13.局部变量只定义不使用会编译出错
14._空白标识符用于抛弃值
15.常量定义 const identifier [type] = value
16.常量表达式中，函数必须是内置函数否则编译不过
17.const(
	vname1 = v_value1
	vname2 = v_value2
)
	用作枚举
18.iota 在每一个const关键字出现时，被重置为0,然后在下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
19.函数定义 func function_name( [parameter list] ) [return_type] {
				函数体
			} 
20.函数可以返回多个值
21.数组声明 var variable_name [SIZE] variable_type
22.声明指针 var var_name *var-type
23.空指针为nil
24.结构体定义 type struct_variable_type struct {
	member definition
	member definition
	...
}

