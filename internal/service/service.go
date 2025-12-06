package service

import (
	"strings"
	
	"six-sprint/pkg/morse"
)

// Анализируем входную строку и выполняем конвертацию
// Возвращаем результат конвертации и ошибку, если она возникла
func AnalysisAndConvert(input string) (string, error) {
	// Проверяем, что входная строка не пустая
	if input == "" {
		return "", nil
	}

	// Очищаем строку от лишних пробелов в начале и конце
	input = strings.TrimSpace(input)

	// Определяем, является ли строка кодом Морзе
	// Код Морзе состоит из: точек (.), тире (-), пробелов между символами и кодами
	isMorse := isMorseCode(input)

	if isMorse {
		// Конвертируем код Морзе в текст
		result := morse.ToText(input)
		if result == "" {
			return "", &ConversionError{"Morse code could not be decoded: possibly an incorrect format"}
		}
		return result, nil
	} else {
		// Конвертируем текст в код Морзе
		result := morse.ToMorse(input)
		if result == "" {
			return "", &ConversionError{"failed to encode text into Morse code: possibly invalid characters"}
		}
		return result, nil
	}
}

// Проверяем, является ли строка кодом Морзе.
func isMorseCode(s string) bool {
	for _, char := range s {
		switch char {
		case '.', '-', ' ':
			continue
		default:
			return false
		}
	}
	return true
}

// Представляем ошибку конвертации
type ConversionError struct {
	Message string
}

// Реализуем интерфейс error для ConversionError
func (e *ConversionError) Error() string {
	return e.Message
}
