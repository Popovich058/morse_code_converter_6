package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"six-sprint/internal/service"
)

// Обрабатываем корневой эндпоинт /.
// Возвращаем содержимое файла index.html.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Failed to open index.html", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write(content)
	if err != nil {
		return
	}
}

// Обрабатываем эндпоинт /upload.
// - Парсим multipart-форму.
// - Читаем загруженный файл.
// - Конвертируем его содержимое через service.DetectAndConvert.
// - Сохраняем результат в локальный файл.
// - Возвращаем результат конвертации в ответе.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "No file uploaded or invalid form field", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusInternalServerError)
		return
	}

	converted, err := service.AnalysisAndConvert(string(data))
	if err != nil {
		http.Error(w, "Conversion failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Формируем имя файла: timestamp + исходное расширение
	timestamp := time.Now().UTC().Format("20060102150405.000") 
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt"
	}
	outputFilename := timestamp + ext

	outFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.Write([]byte(converted))
	if err != nil {
		http.Error(w, "Failed to write to output file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = w.Write([]byte(converted))
	if err != nil {
		return
	}
}
