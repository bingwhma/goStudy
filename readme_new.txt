golang的new和make主要区别如下：

1、make只能用来分配及初始化类型为slice，map，chan的数据；new可以分配任意类型的数据
2、new分配返回的是指针，即类型*T；make返回引用，即T；
3、new分配的空间被清零，make分配后，会进行初始化。effective go举了一个例子，见：http://golang.org/doc/effective_go.html#allocation_make

表达式 new(T) 分配了一个零初始化的 T 值，并返回指向它的指针。
var t *T = new(T)
或
t := new(T)