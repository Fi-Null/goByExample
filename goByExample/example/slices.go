package main

import "fmt"

func main() {
	s := make([]string, 3)//建了一个长度为3的 string 类型 slice（初始化为零值）
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2 : 5]//slice[low:high] 语法进行“切片”操作,
	//如这里得到一个包含元素 s[2], s[3],s[4] 的 slice
	fmt.Println("sl1:", l)

	l = s[: 5]//这个 slice 从 s[0] 到（但是包含）s[5]
	fmt.Println("sl2:", l)

	l = s[2 :]//这个 slice 从（包含）s[2] 到 slice 的后一个值
	fmt.Println("sl3:", l)

	t := []string {"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}