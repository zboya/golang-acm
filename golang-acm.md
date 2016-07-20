# golang ACM练习场

* A+B问题 （难度0）

**描述**

此题为练手用题，请大家计算一下a+b的值

**输入**

输入两个数，a,b

**输出**

输出a+b的值

**样例输入**

2 3

**样例输出**

5

```go
package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Println(a + b)
}

```

* Fibonacci数 (难度：1)

**描述**

无穷数列1，1，2，3，5，8，13，21，34，55...称为Fibonacci数列，它可以递归地定义为
F(n)=1 ...........(n=1或n=2)
F(n)=F(n-1)+F(n-2).....(n>2)
现要你来求第n个斐波纳奇数。（第1个、第二个都为1)

**输入**

第一行是一个整数m(m<5)表示共有m组测试数据
每次测试数据只有一行，且只有一个整形数n(n<20)

**输出**

对每组输入n，输出第n个Fibonacci数

**样例输入**

3
1
3
5

**样例输出**

1
2
5
```go

```
* ASCII码排序 （难度2）

**描述**

输入三个字符（可以重复）后，按各字符的ASCII码从小到大的顺序输出这三个字符。

**输入**

第一行输入一个数N,表示有N组测试数据。后面的N行输入多组数据，每组输入数据都是占一行，有三个字符组成，之间无空格。

**输出**

对于每组输入数据，输出一行，字符中间用一个空格分开。

**样例输入**

2
qwe
asd

**样例输出**

e q w
a d s

```go
package main

import "fmt"

func main() {

    var n int //n组数据
    var str string
    var result string

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&str)
        result = result + sort(&str) + "\n"
    }
    fmt.Println(result)
}

func sort(str *string) string {
    strByte := []byte(*str)
    n := len(strByte)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if strByte[j+1] < strByte[j] {
                strByte[j+1], strByte[j] = strByte[j], strByte[j+1]
            }
        }

    }
    result := insertSpace(&strByte)
    return string(result)
}

func insertSpace(src *[]byte) []byte {
    l := len(*src)
    dst := make([]byte, 2*l)
    for i := 0; i < len(dst); i++ {
        if i%2 == 0 {
            dst[i] = (*src)[i/2]
        } else {
            dst[i] = ' '
        }
    }
    return dst
}

```
* 括号配对问题 （难度3）

**描述**

现在，有一行括号序列，请你检查这行括号是否配对。

**输入**

第一行输入一个数N`（0<N<=100）`,表示有N组测试数据。后面的N行输入多组输入数据，每组输入数据都是一个字符串S(S的长度小于10000，且S不是空串），测试数据组数少于5组。数据保证S中只含有"[","]","(",")"四种字符

**输出**

每组输入数据的输出占一行，如果该字符串中所含的括号是配对的，则输出Yes,如果不配对则输出No

**样例输入**

3
[(])
(])
([[]()])

**样例输出**

No
No
Yes

```go
package main

import (
    "fmt"
)

func main() {
    var n int
    var str string
    var result string
    fmt.Scan( &n)
    for i := 0; i < n; i++ {
        fmt.Scan( &str)
        //match
        rst := matchPairBracket(&str)
        result = result + rst
    }
    fmt.Printf("----\n%s", result)
}

func matchPairBracket(str *string) string {
    var stack []byte
    r := []byte("()[]")
    strLen := len(*str) % 2
    strByte := []byte(*str)

    if strLen == 1 || r[0] != strByte[0] && r[2] != strByte[0] {
        return "NO\n"
    } else {
        for i := 0; i < len(*str); i++ {
            if r[0] == strByte[i] || r[2] == strByte[i] {
                push(&stack, strByte[i])
            } else {
                if r[1] == strByte[i] && r[0] == stack[len(stack)-1] {
                    pop(&stack)
                } else if r[3] == strByte[i] && r[2] == stack[len(stack)-1] {
                    pop(&stack)
                } else {
                    push(&stack, strByte[i])
                }
            }
        }
        if len(stack) == 0 {
            return "YES\n"
        } else {
            return "NO\n"
        }
    }

}

func push(stack *[]byte, r byte) *[]byte {
    *stack = append(*stack, r)
    return stack

}

func pop(stack *[]byte) {
    *stack = (*stack)[:len(*stack)-1]
}

```

* 二进制字符串匹配 （难度：3）

**描述**

给定两个字符串A、B，这两个字符只包含‘0’和‘1’，你的任务就是计算在B中A出现了多少次？例如：A：‘11’，B：‘1001110110’，结果应该是3.

**输入**

第一行输入的数N，表示输入有N组测试数据，每组测试数据包含两行，第一行为字符A，A的长度<=10，第二行为字符B，B的长度<=1000，并且保证字符B的长度大于字符A的长度。

**输出**

每组测试数据的输出是一个整数，这个整数是A在B中出现的次数。

**样例输入**

3
11
1001110110
101
110010010010001
1010
110100010101011

**样例输出**

3
0
3
```go

```

* 喷水装置（一）（难度:3）

**描述**

现有一块草坪，长为20米，宽为2米，要在横中心线上放置半径为Ri的喷水装置，每个喷水装置的效果都会让以它为中心的半径为实数Ri(0<Ri<15)的圆被湿润，这有充足的喷水装置`i（1<i<600)`个，并且一定能把草坪全部湿润，你要做的是：选择尽量少的喷水装置，把整个草坪的全部湿润。

**输入**

第一行m表示有m组测试数据
每一组测试数据的第一行有一个整数数n，n表示共有n个喷水装置，随后的一行，有n个实数ri，ri表示该喷水装置能覆盖的圆的半径。

**输出**

输出所用装置的个数

**样例输入**

2
5
2 3.2 4 4.5 6 
10
1 2 3 1 2 1.2 3 1.1 1 2

**样例输出**

2
5


* 街区最短路径问题 （难度：4）

**描述**

一个街区有很多住户，街区的街道只能为东西、南北两种方向。
住户只可以沿着街道行走。
各个街道之间的间隔相等。
用(x,y)来表示住户坐在的街区。
例如（4,20），表示用户在东西方向第4个街道，南北方向第20个街道。
现在要建一个邮局，使得各个住户到邮局的距离之和最少。
求现在这个邮局应该建在那个地方使得所有住户距离之和最小；

**输入**

第一行一个整数n<20，表示有n组测试数据，下面是n组数据;
每组第一行一个整数m<20,表示本组有m个住户，下面的m行每行有两个整数`0<x,y<100`，表示某个用户所在街区的坐标。
m行后是新一组的数据；

**输出**

每组数据输出到邮局最小的距离和，回车结束；

**样例输入**

2
3
1 1
2 1
1 2
5
2 9 
5 20
11 9
1 1
1 20

**样例输出**

2
44

```go

```
* 多边形重心问题 （难度5）

**描述**

在某个多边形上，取n个点，这n个点顺序给出，按照给出顺序将相邻的点用直线连接， （第一个和最后一个连接），所有线段不和其他线段相交，但是可以重合，可得到一个多边形或一条线段或一个多边形和一个线段的连接后的图形； 
如果是一条线段,我们定义面积为0，重心坐标为（0,0）.现在求给出的点集组成的图形的面积和重心横纵坐标的和；

**输入**

第一行有一个整数`0<n<11`,表示有n组数据；
每组数据第一行有一个整数m<10000,表示有这个多边形有m个顶点；

**输出**

输出每个多边形的面积、重心横纵坐标的和，小数点后保留三位；

**样例输入**

3
3
0 1
0 2
0 3
3
1 1
0 0
0 1
4
1 1
0 0
0 0.5
0 1

**样例输出**

0.000 0.000
0.500 1.000
0.500 1.000

```go
package main

import (
    "fmt"
    "math"
)

type point struct {
    x float32
    y float32
}

type result struct {
    a float32
    c float32
}

func main() {

    var n int   //n组数据
    var m int   //m个顶点
    var p point //坐标
    var ps []point
    var rlt []result

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&m)
        for j := 0; j < m; j++ {
            fmt.Scan(&p.x, &p.y)
            ps = append(ps, p)
        }
        r := handle(ps)
        ps = []point{}
        rlt = append(rlt, *r)
    }
    fmt.Println("--------")
    for _, v := range rlt {
        fmt.Println(v.a, v.c)
    }
}

func handle(ps []point) *result {
    var x, y, areaSum float32 = 0, 0, 0
    m := len(ps)
    if m < 3 {
        return &result{0, 0}
    } else {
        var area float32
        for i := 0; i < m-2; i++ {
            tri, _ := areaTriangle(ps[i], ps[i+1], ps[i+2])
            if !tri {
                remove(&ps, i+1)
            }
        }

        mTrue := len(ps)
        mFloat := float32(mTrue)
        for i := 0; i < mTrue-2; i++ {
            if i < 1 {
                _, area = areaTriangle(ps[i], ps[i+1], ps[i+2])
            } else {
                _, area = areaTriangle(ps[0], ps[i+1], ps[i+2])
            }
            areaSum = areaSum + area

        }
        for _, p := range ps {
            x = x + p.x
            y = y + p.y
        }
        if mTrue < 3 {
            return &result{0, 0}
        } else {
            return &result{areaSum, (x + y) / mFloat}
        }

    }

}

//三点的面积
func areaTriangle(p0, p1, p2 point) (bool, float32) {
    t1 := p0.x*p1.y + p1.x*p2.y + p2.x*p0.y
    t2 := p1.x*p0.y + p2.x*p1.y + p0.x*p2.y
    if t1 == t2 {
        return false, 0
    } else {
        area := math.Abs(float64(t2-t1) / 2)
        return true, float32(area)
    }
}

func remove(slice *[]point, i int) {
    *slice = append((*slice)[:i], (*slice)[i+1:]...)
}

```



* 一种排序

**描述**

现在有很多长方形，每一个长方形都有一个编号，这个编号可以重复；还知道这个长方形的宽和长，编号、长、宽都是整数；现在要求按照一下方式排序（默认排序规则都是从小到大）；
1.按照编号从小到大排序
2.对于编号相等的长方形，按照长方形的长排序；
3.如果编号和长都相同，按照长方形的宽排序；
4.如果编号、长、宽都相同，就只保留一个长方形用于排序,删除多余的长方形；最后排好序按照指定格式显示所有的长方形；

**输入**

第一行有一个整数 `0<n<10000`,表示接下来有n组测试数据；
每一组第一行有一个整数 `0<m<1000`，表示有m个长方形；
接下来的m行，每一行有三个数 ，第一个数表示长方形的编号，
第二个和第三个数值大的表示长，数值小的表示宽，相等
说明这是一个正方形（数据约定长宽与编号都小于10000）；

**输出**

顺序输出每组数据的所有符合条件的长方形的 编号 长 宽

**样例输入**

1
8
1 1 1
1 1 1
1 1 2
1 2 1
1 2 2
2 1 1
2 1 2
2 2 1

**样例输出**

1 1 1
1 2 1
1 2 2
2 1 1
2 2 1

```go

```

* posters

**描述**

The citizens of Bytetown, AB, could not stand that the candidates in the mayoral election campaign have been placing their electoral posters at all places at their whim. The city council has finally decided to build an electoral wall for placing the posters and introduce the following rules: 
•   Every candidate can place exactly one poster on the wall. 
•   All posters are of the same height equal to the height of the wall; the width of a poster can be any integer number of bytes (byte is the unit of length in Bytetown). 
•   The wall is divided into segments and the width of each segment is one byte. 
•   Each poster must completely cover a contiguous number of wall segments.
They have built a wall 10000000 bytes long (such that there is enough place for all candidates). When the electoral campaign was restarted, the candidates were placing their posters on the wall and their posters differed widely in width. Moreover, the candidates started placing their posters on wall segments already occupied by other posters. Everyone in Bytetown was curious whose posters will be visible (entirely or in part) on the last day before elections. 
Your task is to find the number of visible posters when all the posters are placed given the information about posters' size, their place and order of placement on the electoral wall. 

**输入**

The first line of input contains a number c giving the number of cases that follow. The first line of data for a single case contains number 1 <= n <= 10000. The subsequent n lines describe the posters in the order in which they were placed. The i-th line among the n lines contains two integer numbers li and ri which are the number of the wall segment occupied by the left end and the right end of the i-th poster, respectively. We know that for each 1 <= i <= n, 1 <= li <= ri <= 10000000. After the i-th poster is placed, it entirely covers all wall segments numbered li, li+1 ,... , ri.

**输出**

For each input data set print the number of visible posters after all the posters are placed. 
The picture below illustrates the case of the sample input. 
![pic1](http://acm.pku.edu.cn/JudgeOnline/images/2528_1.jpg)

**样例输入**

1
5
1 4
2 6
8 10
3 4
7 10

**样例输出**

4

```go

```

* skiing

**描述**

Michael喜欢滑雪百这并不奇怪， 因为滑雪的确很刺激。可是为了获得速度，滑的区域必须向下倾斜，而且当你滑到坡底，你不得不再次走上坡或者等待升降机来载你。Michael想知道载一个区域中最长底滑坡。区域由一个二维数组给出。数组的每个数字代表点的高度。下面是一个例子 
1 2 3 4 5
16 17 18 19 6
15 24 25 20 7
14 23 22 21 8
13 12 11 10 9
一个人可以从某个点滑向上下左右相邻四个点之一，当且仅当高度减小。在上面的例子中，一条可滑行的滑坡为24-17-16-1。当然25-24-23-...-3-2-1更长。事实上，这是最长的一条。

**输入**

第一行表示有几组测试数据，输入的第二行表示区域的行数R和列数`C(1 <= R,C <= 100)`。下面是R行，每行有C个整数，代表高度h，0<=h<=10000。
后面是下一组数据；

**输出**

输出最长区域的长度。

**样例输入**

1
5 5
1 2 3 4 5
16 17 18 19 6
15 24 25 20 7
14 23 22 21 8
13 12 11 10 9

**样例输出**

25
```go

```

* 奇偶数分离 （难度：1）

**描述**

有一个整型偶数n(2<= n <=10000),你要做的是：先把1到n中的所有奇数从小到大输出，再把所有的偶数从小到大输出。

**输入**

第一行有一个整数i（2<=i<30)表示有 i 组测试数据；
每组有一个整型偶数n。

**输出**

第一行输出所有的奇数
第二行输出所有的偶数

**样例输入**

2
10
14

**样例输出**

1 3 5 7 9 
2 4 6 8 10 

1 3 5 7 9 11 13 
2 4 6 8 10 12 14
```go

```

* 喷水装置（二）(难度：4)

**描述**

有一块草坪，横向长w,纵向长为h,在它的橫向中心线上不同位置处装有n(n<=10000)个点状的喷水装置，每个喷水装置i喷水的效果是让以它为中心半径为Ri的圆都被润湿。请在给出的喷水装置中选择尽量少的喷水装置，把整个草坪全部润湿。

**输入**

第一行输入一个正整数N表示共有n次测试数据。
每一组测试数据的第一行有三个整数n,w,h，n表示共有n个喷水装置，w表示草坪的横向长度，h表示草坪的纵向长度。
随后的n行，都有两个整数xi和ri,xi表示第i个喷水装置的的横坐标（最左边为0），ri表示该喷水装置能覆盖的圆的半径。

**输出**

每组测试数据输出一个正整数，表示共需要多少个喷水装置，每个输出单独占一行。
如果不存在一种能够把整个草坪湿润的方案，请输出0。

**样例输入**

2
2 8 6
1 1
4 5
2 10 6
4 5
6 5

**样例输出**

1
2
```go

```

