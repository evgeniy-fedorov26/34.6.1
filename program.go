package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Проверяем, что передано два аргумента: входной и выходной файлы
	if len(os.Args) < 3 {
		fmt.Println("Использование: go run main.go <входной_файл> <выходной_файл>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Читаем содержимое входного файла
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Ошибка чтения входного файла:", err)
		return
	}

	// Разбиваем файл на строки
	lines := strings.Split(string(data), "\n")
	// Регулярное выражение для поиска математических выражений
	regex := regexp.MustCompile(`^(\d+)([+-])(\d+)=\?$`)

	// Создаем (или очищаем) выходной файл
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка создания выходного файла:", err)
		return
	}
	defer outFile.Close()

	// Создаем буферизированный писатель
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Обрабатываем каждую строку
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if matches != nil {
			// Преобразуем числа из строки в целые числа
			num1, _ := strconv.Atoi(matches[1])
			num2, _ := strconv.Atoi(matches[3])
			var result int

			// Определяем операцию и выполняем вычисление
			switch matches[2] {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			}

			// Формируем строку с результатом и записываем в файл
			output := fmt.Sprintf("%s%s%s=%d\n", matches[1], matches[2], matches[3], result)
			writer.WriteString(output)
		}
	}
}
