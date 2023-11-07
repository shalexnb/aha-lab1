package main

import (
	"fmt"
	"net/http"
)


var usersDB = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Сервер запущен на http://localhost:3000")

	for username, password := range usersDB {
		fmt.Printf("Пользователь в БД: Имя пользователя - %s, Пароль - %s\n", username, password)
		break
	}

	http.ListenAndServe(":3000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	message := ""

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if storedPassword, ok := usersDB[username]; ok && storedPassword == password {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		} else {
			message = "Неверное имя пользователя или пароль. Попробуйте снова."
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Авторизация</title>
			<style>

				body {
					display: flex;
					justify-content: center;
					align-items: flex-start;
					height: 100vh;
					margin: 0;
				}

				.container {
					text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Авторизация</h1>
				<p>%s</p>
				<form method="post" action="/">
					<label for="username">Имя пользователя:</label>
					<input type="text" id="username" name="username" required><br>
					<br>
					<label for="password">Пароль:</label>
					<input type="password" id="password" name="password" required><br>
					<br>
					<button type="submit">Войти</button>
				</form>
			</div>
		</body>
		</html>
	`, message)
}


func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Панель управления</title>
			<style>

				body {
					display: flex;
					justify-content: center;
					align-items: flex-start;
					height: 100vh;
					margin: 0;
				}

				.container {
					text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Нет времени объяснять</h1>
				<p>Текущее время: <span id="current-time"></span></p>

				<img src="https://cs.pikabu.ru/post_img/2013/05/05/7/1367748024_1928166780.jpg" alt="картинка с шуткой">

				<script>
					function updateCurrentTime() {
						var currentTimeElement = document.getElementById("current-time");
						var currentTime = new Date().toLocaleTimeString();
						currentTimeElement.textContent = currentTime;
					}

					updateCurrentTime();
					setInterval(updateCurrentTime, 1000);
				</script>
			</div>
		</body>
		</html>
	`)
}