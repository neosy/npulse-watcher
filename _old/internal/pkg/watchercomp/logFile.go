package watchercomp

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func fileNameAddYear(fileName string) string {
	var name = fileName
	timeNow := time.Now()

	lastDot := strings.LastIndex(name, ".")
	if lastDot != -1 {
		name = name[:lastDot] + fmt.Sprintf("_%v", timeNow.Year()) + name[lastDot:]
	}

	return name
}

func homeFolderSignUpdate(path string) (string, error) {
	// Если путь начинается с ~, заменяем его на домашнюю директорию
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Println("Error getting the home directory:", err)
			return path, err
		}
		// Заменяем ~ на домашнюю директорию
		path = filepath.Join(homeDir, path[1:])
	}

	return path, nil
}

func findOrCreateLogFile(filePath string) error {
	logFolderPath, _ := filepath.Split(filePath)

	// Проверяем, существует ли директория
	if _, err := os.Stat(logFolderPath); os.IsNotExist(err) {
		// Если директория не существует, создаем её (включая все промежуточные)
		err := os.MkdirAll(logFolderPath, os.ModePerm)
		if err != nil {
			log.Println("Error creating directories:", err)
			return err
		}
	}

	// Если файл не найден, то создаем его
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Println("Error creating file:", err)
			return err
		}
		defer file.Close()
	}

	return nil
}

func getCompIPsFromFile(filePath string) (map[string]compState, error) {
	compIPs := make(map[string]compState)

	err := findOrCreateLogFile(filePath)

	if err != nil {
		return compIPs, err
	}

	// Открываем файл для чтения
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, err
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
		cIP := parts[2]
		cName := parts[3]
		cStatus, _ := strconv.Atoi(parts[4])

		// Преобразуем строку в time.Time
		dateTimeLayou := "2006-01-02 15:04:05"
		parsedTime, err := time.Parse(dateTimeLayou, dateStr)
		if err != nil {
			log.Println("Ошибка при преобразовании:", err)
		}

		// Сохраняем имя компьютера в map
		compIPs[cIP] = compState{
			ip:       cIP,
			name:     cName,
			lastTime: parsedTime,
			status:   CompWatchStatus(cStatus),
		}
	}

	return compIPs, nil
}

func writeToFile(filePath string, ip string, name string, time time.Time, status CompWatchStatus) (err error) {
	var file *os.File

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return err
		}
	} else {

		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Error opening the file:", err)
			return err
		}
	}
	defer file.Close()

	// Форматируем строку
	line := fmt.Sprintf("%s %s %s %s %d\n", time.Format("2006-01-02"), time.Format("15:04:05"), ip, name, status)

	// Записываем строку в файл
	if _, err := file.WriteString(line); err != nil {
		log.Println("Error writing to the file:", err)
	}

	return err
}
