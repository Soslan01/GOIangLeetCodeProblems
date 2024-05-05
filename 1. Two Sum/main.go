package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//Map для хранения индексов
	indexMap := make(map[int]int)

	//Итерация по массиву
	for i, num := range nums {
		//Вычисляем разницу между целевым и текущим числами
		complement := target - num
		//Проверяем есть ли число complement в Map
		if value, found := indexMap[complement]; found {
			//Если это число найдено, то возвращаем индекс это числа в Map
			return []int{value, i}
		}
		//Добавляем найденное число в Map
		indexMap[num] = i
	}
	return nil
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3))
}
