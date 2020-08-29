package main

import (
	"fmt"
)

var c, python, java = false, true, false

func main() {
	/* 	// 関数内でのvar宣言+初期化省略表現
	   	name := "kotodu"
	   	number := 43
	   	const vue string = "vue"
	   	fmt.Sprintln(number)
	   	// n := string(number)
	   	fmt.Println("Hello, 世界")
	   	fmt.Println("時刻", time.Now())
	   	fmt.Println("My favorite number is", rand.Intn(10))
	   	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	   	fmt.Println(math.Pi)
	   	fmt.Println(add(45, 56))
	   	fmt.Println(c, python, java, name)
	   	// fmt.Println(n)
	       fmt.Println(vue) */
	forfor()
	bo(true)
	point()
	stackDefer()
	createBusRoutes()
	ary()
	rangerange()
	mapm()
	createFibonacci()
}

func add(a, b int) int {
	return a + b
}

// namedReturn
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func forfor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func bo(b bool) {
	if b {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func point() {
	i := 42
	j := 2701

	// ポインタ参照
	// iのメモリにある値を読み取る
	p := &i

	// ポインタpを通してiから値を読みだす
	fmt.Println(*p)

	// ポインタpを通してiへ値を代入する
	*p = 21

	fmt.Println(i)

	// ポインタ参照
	// jのメモリにある値を読み取る
	p = &j

	// j/37
	*p = *p / 37
	fmt.Println(j)
}

func stackDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//---------------------------------------------

// BusRoute : バス系統構造体
type BusRoute struct {
	RouteName  string
	RouteID    string
	TripCount  int
	IsMidnight bool
}

func createBusRoutes() {
	route1 := BusRoute{"都01", "Toei.To-01", 100, false}
	route2 := BusRoute{"東98", "Tokyu.To-98", 40, false}
	route3 := &route1
	route3.RouteName = "都01折返"
	fmt.Println(route1.RouteName)
	fmt.Println(route2.RouteID)
}

var arr = [2]string{"はろー", "ワールド！"}

// スライス : 可変長配列のようなもの
// 内部的には、元の配列を参照している
var sli = []string{
	"あ",
	"い",
	"う",
	"え",
	"お"}

func ary() {
	fmt.Println(arr)
	fmt.Println(sli)
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// 省略できる
	// 最大値は含まない
	var sli2 = sli[:4]
	fmt.Println(sli2[0:2])

	// len……実際に含まれる要素数
	fmt.Println("len", len(sli2))

	// cap……sliceが参照する配列の容量
	fmt.Println("cap", cap(sli2))

	// スライスのゼロ値はnil
	var sn []int
	fmt.Println(sn, len(sn), cap(sn))
	if sn == nil {
		fmt.Println("nil!")
	}

	// make……動的サイズの配列を作る関数
	a := make([]int, 5)
	printSlice("a", a)

	// make(配列,len,cap)
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// append……要素を追加する
	s = append(s, 2)
	printSlice("s", s)
	// 11が2に置き換わっているのが分かる
	fmt.Println(primes)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func rangerange() {
	pow := []int{
		1,
		2,
		4,
		8,
		16,
		32,
		64,
		128,
	}
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

// LatLong : 緯度経度
type LatLong struct {
	Lat, Long float64
}

func mapm() {
	// latlong
	m := make(map[string]LatLong)
	m["東京駅"] = LatLong{
		35.6809591,
		139.7673068,
	}
	fmt.Println(m["東京駅"])
	// mapの要素が存在するか、2つ目にboolが戻る
	v, ok := m["新宿駅"]
	fmt.Println("新宿駅の値 : ", v, "  存在有無 : ", ok)

}

func fibonacci() func() int {
	var t1 int = 0
	var t2 int = 1
	var i int = 0
	return func() int {
		next := 0
		switch {
		case i == 1:
			next = 1
		case i == 0:
			next = 0
		default:
			next = t1 + t2
			t1 = t2
			t2 = next
		}
		i++
		return next
	}
}

func createFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
