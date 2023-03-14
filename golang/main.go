package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"isuct.ru/informatics2022/internal/module1"
	"isuct.ru/informatics2022/internal/module2"
)

func main() {
	fmt.Println("Hello world")
	module1.Summ()
	module2.BubbleSort()
	module2.PairSort()
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr := strings.Split(line, " ")
	arr := make([]int, n)
	for idx, val := range str_arr {
		arr[idx], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	arr = module2.MergeSort(arr, 0, n)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	module2.DifferentCount()
	module2.InversionCount()
	module2.Stock()

}
