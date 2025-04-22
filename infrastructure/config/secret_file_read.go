package iconfig

import (
	"fmt"
	"log"
	"os"
)

// Загрузка данных из файла secret
func secretFileRead(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Panic(
			fmt.Sprintf("Can't read secret file %v", name),
			err,
		)
		return ""
	}

	return string(data)
}
