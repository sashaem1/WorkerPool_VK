package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("data/test.txt")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	//задаёт количество строк в выходном файле
	N := 100

	for i := 1; i <= N; i++ {
		writer.WriteString(fmt.Sprintf("Строка %d\n", i)) // запись строки
	}
	writer.Flush() // сбрасываем данные из буфера в файл

}
