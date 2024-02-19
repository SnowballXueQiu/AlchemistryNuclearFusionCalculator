package main

import (
	"bufio"
	"fmt"
	"nuclear_fusion_calculator/elements_map"
	"os"
	"strconv"
	"strings"
)

func listCombinations(n int) [][]int {
	var combinations [][]int
	for i := 1; i <= n/2; i++ {
		combinations = append(combinations, []int{i, n - i})
	}
	return combinations
}

func main() {
	for {
		var input int
		fmt.Print("请输入一个数字 (范围为2到118): ")
		inputStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		inputStr = strings.TrimSpace(inputStr)
		num, err := strconv.Atoi(inputStr)
		if err != nil || num < 2 || num > 118 {
			fmt.Println("输入错误: 请输入一个范围在2到118之间的整数")
			continue
		}
		input = num

		combinations := listCombinations(input)

		fmt.Printf("数字 %d 的所有组合为: \n", input)
		for _, combination := range combinations {
			fmt.Printf("(%d: %s, %d: %s)\n", combination[0], elements_map.Elements[combination[0]], combination[1], elements_map.Elements[combination[1]])
		}

		fmt.Println("按下回车继续，或者按下 Ctrl+C 退出")
		reader := bufio.NewReader(os.Stdin)
		if _, err := reader.ReadString('\n'); err != nil {
			fmt.Println("无法读取输入: ", err)
			return
		}
		if _, err := reader.ReadString('\n'); err != nil {
			fmt.Println("无法读取输入: ", err)
			return
		}
	}
}
