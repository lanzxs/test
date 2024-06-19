package main

import (
	"fmt"
)

// 快速排序
func quikSort(arr []int, left, right int) {
	if left < right {
		p := pation(arr, left, right)
		quikSort(arr, left, p-1)
		quikSort(arr, p+1, right)
	}
}

func pation(arr []int, left, right int) int {

	provi := arr[left]
	for left < right {
		for left < right && arr[right] >= provi {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= provi {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = provi
	return left
}

// 选择排序
func xuan(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		minindex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}
		if minindex != i {
			arr[minindex], arr[i] = arr[i], arr[minindex]
		}
	}
}

// 冒泡排序
func mao(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 插入排序
func cha(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		min := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > min {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = min
	}
}

type BitMap []byte

const byteSize = 8 //定义的bitmap为byte的数组，byte为8bit

func NewBitMap(n uint) BitMap {
	return make([]byte, n/byteSize+1)
}
func (bt BitMap) set(n uint) {
	if n/byteSize > uint(len(bt)) {
		fmt.Println("大小超出bitmap范围")
		return
	}
	byteIndex := n / byteSize //第x个字节（0,1,2...）
	fmt.Println(byteIndex)
	offsetIndex := n % byteSize //偏移量(0<偏移量<byteSize)
	fmt.Println(offsetIndex)
	//bt[byteIndex] = bt[byteIndex] | 1<<offsetIndex //异或1（置位）
	//第x个字节偏移量为offsetIndex的位 置位1
	fmt.Println(bt[byteIndex])
	fmt.Println(1 << offsetIndex)
	fmt.Println(bt[byteIndex] | 1<<offsetIndex)
	bt[byteIndex] |= 1 << offsetIndex //异或1（置位）

}

func (bt BitMap) isExist(n uint) bool {
	if n/byteSize > uint(len(bt)) {
		fmt.Println("大小超出bitmap范围")
		return false
	}
	byteIndex := n / byteSize
	offsetIndex := n % byteSize
	//fmt.Println(bt[byteIndex] & (1 << offsetIndex))
	return bt[byteIndex]&(1<<offsetIndex) != 0 //TODO：注意：条件是 ！=0，有可能是：16,32等
}

func main() {
	//arr:= []int{2,4,6,8,3,1,8,10,5,9,3}
	//quikSort(arr,0,len(arr)-1)
	//fmt.Println(arr)
	bitmap := NewBitMap(20)
	fmt.Println(bitmap)
	bitmap.set(20)
	fmt.Println(bitmap)
	bitmap.set(19)
	fmt.Println(bitmap)
	bitmap.set(18)
	fmt.Println(bitmap)
}
