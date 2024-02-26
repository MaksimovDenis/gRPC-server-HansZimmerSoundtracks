console.clear();

const loginBtn = document.getElementById('login');
const signupBtn = document.getElementById('signup');

loginBtn.addEventListener('click', (e) => {
	let parent = e.target.parentNode.parentNode;
	Array.from(e.target.parentNode.parentNode.classList).find((element) => {
		if(element !== "slide-up") {
			parent.classList.add('slide-up')
		}else{
			signupBtn.parentNode.classList.add('slide-up')
			parent.classList.remove('slide-up')
		}
	});
});

signupBtn.addEventListener('click', (e) => {
	let parent = e.target.parentNode;
	Array.from(e.target.parentNode.classList).find((element) => {
		if(element !== "slide-up") {
			parent.classList.add('slide-up')
		}else{
			loginBtn.parentNode.parentNode.classList.add('slide-up')
			parent.classList.remove('slide-up')
		}
	});
});

$(document).ready(function() {
    $('.signup').submit(function(event) {
        event.preventDefault(); // Предотвращаем стандартное поведение формы (отправку)

        var email = $('.signup input[type="email"]').val();
        var password = $('.signup input[type="password"]').val();
        
        console.log("Email:", email);
        console.log("Password:", password);
        
        // Отправляем AJAX запрос на сервер
        $.ajax({
            type: "POST", // Метод запроса
            url: "/signup", // URL адрес обработчика на сервере
            data: JSON.stringify({ email: email, password: password }), // Данные для отправки (преобразованные в JSON строку)
            contentType: "application/json", // Тип содержимого запроса
            success: function(response) {
                // Обработка успешного ответа от сервера
                console.log("Response from server:", response);
            },
            error: function(xhr, status, error) {
                // Обработка ошибки
                console.error("Error:", error);
            }
        });
        
        // Обновляем значения элементов на странице (это можно удалить, если не нужно)
        $('#emailValue').text(email);
        $('#passwordValue').text(password);
        
        // В этом месте вы можете выполнить другие операции после отправки данных на сервер
    });
});

$(document).ready(function() {
    $('.login').submit(function(event) {
        event.preventDefault(); // Предотвращаем стандартное поведение формы (отправку)

        var email = $('.login input[type="email"]').val();
        var password = $('.login input[type="password"]').val();
        
        console.log("Email:", email);
        console.log("Password:", password);
        
        // Отправляем AJAX запрос на сервер
        $.ajax({
            type: "POST", // Метод запроса
            url: "/login", // URL адрес обработчика на сервере
            data: JSON.stringify({ email: email, password: password }), // Данные для отправки (преобразованные в JSON строку)
            contentType: "application/json", // Тип содержимого запроса
            success: function(response) {
                // Обработка успешного ответа от сервера
                console.log("Response from server:", response);
            },
            error: function(xhr, status, error) {
                // Обработка ошибки
                console.error("Error:", error);
            }
        });
        
        // Обновляем значения элементов на странице (это можно удалить, если не нужно)
        $('#emailValue').text(email);
        $('#passwordValue').text(password);
        
        // В этом месте вы можете выполнить другие операции после отправки данных на сервер
    });
});