package uc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type CompName struct {
	name     string
	lastTime time.Time
}

func fileNameAddYear(fileName string) string {
	var name = fileName
	timeNow := time.Now()

	lastDot := strings.LastIndex(name, ".")
	if lastDot != -1 {
		name = name[:lastDot] + fmt.Sprintf("%v", timeNow.Year()) + name[lastDot:]
	}

	return name
}

func fileFindCompIPs(filePath string) (map[string]CompName, error) {
	compIPs := make(map[string]CompName)

	// Открываем файл для чтения
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return compIPs, err
	}
	defer file.Close()

	// Читаем файл построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Разбиваем строку на части
		parts := strings.Fields(line)
		if len(parts) < 5 {
			// Если строка содержит меньше 5 частей (неправильный формат), пропускаем её
			continue
		}

		// Получаем IP и имя компьютера (3-й и 4-й элемент строки)
		dateStr := fmt.Sprintf("%s %s", parts[0], parts[1])
		compIP := parts[2]
		compName := parts[3]

		// Преобразуем строку в time.Time
		dateTimeLayou := "2006-01-02 15:04:05"
		parsedTime, err := time.Parse(dateTimeLayou, dateStr)
		if err != nil {
			log.Println("Ошибка при преобразовании:", err)
		}

		// Сохраняем имя компьютера в карту
		compIPs[compIP] = CompName{
			name:     compName,
			lastTime: parsedTime,
		}
	}

	return compIPs, nil
}
