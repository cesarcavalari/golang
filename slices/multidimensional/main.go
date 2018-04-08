package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//stutent1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	//store the record
	records = append(records, student1)

	//student2
	student2 := make([]string, 4)
	student2[0] = "GOmez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	//store the record
	records = append(records, student2)

	fmt.Println(records)

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)
}
