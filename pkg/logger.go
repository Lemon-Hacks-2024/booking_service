package pkg

import (
	"io"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Ошибка открытия файла логов: %v", err)
	}
	defer logFile.Close()
	// Создаем новый writer, который пишет и в консоль, и в файл
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "", log.LstdFlags)
	logger.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	return logger
}
