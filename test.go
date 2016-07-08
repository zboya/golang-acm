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
		var ok bool
		for i := 0; i < m-2; i++ {
			ok, _ = isTriangle(ps[i], ps[i+1], ps[i+2])
			if !ok {
				remove(&ps, i+1)
			}
		}

		mTrue := len(ps)
		mFloat := float32(mTrue)
		for i := 0; i < mTrue-2; i++ {
			if i < 1 {
				_, area = isTriangle(ps[i], ps[i+1], ps[i+2])
			} else {
				_, area = isTriangle(ps[0], ps[i+1], ps[i+2])
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

func isTriangle(p0, p1, p2 point) (bool, float32) {
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
