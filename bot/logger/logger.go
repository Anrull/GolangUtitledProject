package logger

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

var Logger *slog.Logger

func New(level slog.Level, filename string) {
    // Создаем файл, если его нет
    dir := filepath.Dir(filename)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        if err := os.MkdirAll(dir, 0755); err != nil {
            log.Printf("Не удалось создать директорию для логирования: %v\n", err)
            panic(err)
        }
    }
    // Открываем файл для записи логов. Если файл не существует, он будет создан.
    // Флаг os.O_APPEND добавляет новые записи в конец файла.
    logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Printf("Не удалось открыть файл для логирования: %v\n", err)
        panic(err)
    }

    // Создаем новый обработчик, который пишет в файл.
    handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
        Level: level,
    })

    logger := slog.New(handler)

	Logger = logger
}


func Info(message string, args ...any) {
	Logger.Info(message)
}

func Error(msg string, args ...any) {
	Logger.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

func Debug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}