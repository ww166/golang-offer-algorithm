# golang-offer-algorithm

## GOLANG 算法

- ## 关于最小值堆栈算法
  1. 继承方式差异：   
    - 1. 只能通过 struct 嵌入方式实现（自动实现同名变量外层覆盖），内层变量需要手动处理（IntMinStack.Init 函数）
    - 2. 类型强制转换，只能强制转换为明企定义的类型
    - 3. reflect 的局限性：语法啰嗦.
    - 4. 判断是否实现某一接口：强制转换为该接口类型，且判断是否转换成功
    
  2. 关于泛型：
    - 1. 通用算法实现思路1：通过空接口实现通用算法。涉及到数据自身处理，需要以函数接口形式实现。
    - 2. 通用算法实现思路2： 手动实现常用类型的通用算法（int, float64, string 等），参照 [go.sort package](https://github.com/golang/go/blob/master/src/sort/sort.go)
    