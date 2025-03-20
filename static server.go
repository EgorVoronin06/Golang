package main

import (
	"log"
	"net/http"
)

func main() {
	// Обслуживаем статические файлы из папки "static"
	fs := http.FileServer(http.Dir("static"))// создает виртуальную файловую систему с корнем в папке "static", создается обработчик HTTP запросов который будет обсуживать файл
	http.Handle("/", fs)//регистрирует созданный файловый сервер для обработки всех запросов по пути   

	// Запускаем сервер на порту 8080
	log.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}
}
