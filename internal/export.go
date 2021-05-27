/*
Copyright © 2021 Roberto Esteves <robertoesteves@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package internal

import (
	"fmt"
	"log"
	"strconv"
)

func ParseInput(args []string) []float64{
	nums := make([]float64, 10)
	for i, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			panic(nil)
		}

		nums[i] = num
	}
	return nums
}

func PrintOutput(num float64, isFloat bool) {
		if isFloat {
			fmt.Println(num)
		} else {
			fmt.Println(int(num))
		}
}

func EvenOrOdd(nums []float64, isEven, isOdd, isFloat bool) ([]float64) {
	if isFloat {
		fmt.Printf("Floating precision will be lost when using even or odd flags.")
	}

	if (isEven && isOdd) {
		log.Fatal("Can't use even and odd at the same time.")
	} else if (isEven) {
		for i, num := range nums {
			if (int(num) % 2) != 0 {
				nums[i] = 0;
			}
		}
	} else if (isOdd) {
		for i, num := range nums {
			if (int(num) % 2) == 0 {
				nums[i] = 0;
			}
		}
	}

	return nums
}