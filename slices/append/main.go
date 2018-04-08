package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	fmt.Println(len(greeting))

	fmt.Println(cap(greeting))

	greeting[0] = "Good morning!"
	greeting[1] = "Bounjour!"
	greeting[2] = "Buenos dias!"

	greeting = append(greeting, "Suprobadham!")
	greeting = append(greeting, "Cao allala!")
	greeting = append(greeting, "Babi lon!")
	greeting = append(greeting, "Cari caro!")

	fmt.Println(greeting[6])

	fmt.Println(len(greeting))

	fmt.Println(cap(greeting))

	greeting2 := []string{
		"Abra cadabra",
		"Valeticesamo",
		"Curururu",
		"Guingoloco!",
		"Uala!",
		"Eitia!",
	}

	newGretting := append(greeting, greeting2...)

	fmt.Println(newGretting)

	fmt.Println(len(newGretting))

	fmt.Println(cap(newGretting))

	oneMoreGretting := append(greeting[:2], greeting2[3:5]...)

	fmt.Println(oneMoreGretting)

	fmt.Println(len(oneMoreGretting))

	fmt.Println(cap(oneMoreGretting))

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
}
