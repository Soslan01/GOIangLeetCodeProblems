package main

import "fmt"

// Создание структуры узла связного списка, куда мы будем записывать числа
type ListNode struct {
	Value int
	Next  *ListNode
}

// Функция для сложения двух чисел, представленных в виде связных списков
func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	//Создаем узел, в котором будем хранить результат
	result := &ListNode{}
	//Переменная, в которой будет храниться указатель на текущий узел результата
	current := result
	//Переменная для хранения значения, которое мы будем переносить
	carry := 0

	//Итерация по связным спискам
	for list1 != nil || list2 != nil {
		//В эти переменные мы будем записывать значения узлов
		val1, val2 := 0, 0
		if list1 != nil {
			val1 = list1.Value
			list1 = list1.Next
		}
		if list2 != nil {
			val2 = list2.Value
			list2 = list2.Next
		}

		//Вычисляем сумму значени узлов
		summ := val1 + val2 + carry
		//Вычисляем значение узла результата
		current.Next = &ListNode{Value: summ % 10}
		//Обновляем переносимое значение
		carry = summ / 10
		//Перемещаем указатель на узел переменной, которая хранит в себе результат
		current = current.Next
	}
	//Проверяем, есть ли остаточное переносимое значение
	if carry > 0 {
		current.Next = &ListNode{Value: carry}
	}
	//Возвращаем узел, в котором храним результат
	return result.Next
}

// Функция для создания связного списка из массива значений
func createLinkedList(nums []int) *ListNode {
	result := &ListNode{}
	current := result
	for _, num := range nums {
		current.Next = &ListNode{Value: num}
		current = current.Next
	}
	return result.Next
}

// Функция для преобразования связного списка в массив
func linkedListToArray(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func main() {
	list1 := createLinkedList([]int{2, 4, 3})
	list2 := createLinkedList([]int{5, 6, 4})

	result12 := addTwoNumbers(list1, list2)
	fmt.Println(linkedListToArray(result12))

	list3 := createLinkedList([]int{0})
	list4 := createLinkedList([]int{0})

	result34 := addTwoNumbers(list3, list4)
	fmt.Println(linkedListToArray(result34))

	list5 := createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	list6 := createLinkedList([]int{9, 9, 9, 9})

	result56 := addTwoNumbers(list5, list6)
	fmt.Println(linkedListToArray(result56))
}
