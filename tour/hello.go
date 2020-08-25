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
