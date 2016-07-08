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
	// var core point //重心
	// var area float32 // 面积

	var ps []point
	var rlt []result

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&p.x, &p.y)
			ps = append(ps, p)
		}
		// fmt.Printf("%#v\n", ps)
		r := handle(ps)
		ps = []point{}
		rlt = append(rlt, *r)
		//deal
	}
	fmt.Println("--------")
	for _, v := range rlt {
		fmt.Println(v.a, v.c)
	}
}

func handle(ps []point) *result {
	var x, y, areaSum, mTrue float32 = 0, 0, 0, 1
	m := len(ps)
	if m < 3 {
		return &result{0, 0}
	} else {
		for i := 0; i < m-2; i++ {
			if i < 4 {
				ok, area := isTriangle(ps[i], ps[i+1], ps[i+2])
				if !ok {
					ps = append(ps[:i+1], ps[i+2:]...)
					// fmt.Println("is a Line")
				}
			} else {
				ok, area := isTriangle(ps[0], ps[i+1], ps[i+2])
				if !ok {
					ps = append(ps[:i+1], ps[i+2:]...)
					// fmt.Println("is a Line")
				}
			}

			fmt.Println("area", area)

			areaSum = areaSum + area

		}

		for _, p := range ps {
			x = x + p.x
			y = y + p.y
		}
		mTrue = float32(len(ps))
		if mTrue < 3 {
			return &result{0, 0}
		} else {
			fmt.Println(x, y, (x+y)/mTrue)
			return &result{areaSum, (x + y) / mTrue}
		}

	}

}

func isTriangle(p0, p1, p2 point) (bool, float32) {
	//判断三点的面积是否为零
	// t1 := (p1[0] - p3[0]) * (p2[1] - p3[1])
	// t2 := (p2[0] - p3[0]) * (p1[1] - p3[1])
	t1 := p0.x*p1.y + p1.x*p2.y + p2.x*p0.y
	t2 := p1.x*p0.y + p2.x*p1.y + p0.x*p2.y
	if t1 == t2 {
		return false, 0
	} else {
		area := math.Abs(float64(t2-t1) / 2)
		return true, float32(area)
	}
}
