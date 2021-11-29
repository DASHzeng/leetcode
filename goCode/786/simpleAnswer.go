package main

import (
	"fmt"
	"sort"
)

type pair struct {
	x int
	y int
}

//暴力做法
func kthSmallestPrimeFraction(arr []int, k int) []int {

	var res []pair
	for i, v := range arr {
		for j := 0; j < i; j++ {
			res = append(res, pair{arr[j], v})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		a, b := res[i], res[j]
		return a.x*b.y < a.y*b.x
	})
	//fmt.Println(res)
	return []int{res[k-1].x, res[k-1].y}
}

//使用优先队列寻找有序矩阵的第K小值

//使用双指针+二分查找 找到刚好大于K-1个元素的目标元素x

func main() {
	fmt.Println("ok")
}
