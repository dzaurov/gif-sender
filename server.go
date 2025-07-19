package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Этот обработчик теперь использует http.ServeFile,
	// чтобы отдать конкретный файл с диска.
	http.HandleFunc("/20352709.gif", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Получен запрос к /20352709.gif, отдаю файл...")
		// Устанавливаем заголовок CORS на случай, если он понадобится для тестов
		w.Header().Set("Access-control-allow-origin", "*")
		// Функция ServeFile автоматически определит Content-Type и другие заголовки
		http.ServeFile(w, r, "20352709.gif")
	})

	// --- Запуск HTTP сервера ---
	addr := ":32018"
	fmt.Printf("Запускаю HTTP сервер на http://127.0.0.1%s\n", addr)
	fmt.Println("Чтобы остановить, нажмите Ctrl+C")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
