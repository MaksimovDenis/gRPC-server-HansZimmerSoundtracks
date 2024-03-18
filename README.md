# World of Hans Zimmer 
gRPC сервер, с поддержкой авторизации и аутентификации пользователя. Данное веб-приложение позоволяет прослушивать композиции к таким фильмам как: Интерсттеллар, Дюна, Бетмен, Начало, Пираты карибского моря, написанные Хансом Циммером. 

## Интсрукция по установке и использованию
- Для запуска приложения, необходимо из корневой директории проекта выполнить команду:    
`docker run -d -p 8080:8080 grpc-hans-zimmer`
- Команда запустит контейнер описанный в файле `Dockerfile`
- После запуска, приложение будет доступно по ссылке:  
[http://localhost:8000](URL)  
![image](https://github.com/MaksimovDenis/gRPC-server-HansZimmerSoundtracks/assets/44647373/b7fa6e27-6963-4600-8331-0d29038b4ca2)
- Для авторизации вы можете использовать следующие данные или зарегестрироваться в системе:
  login: admin@74.ru  
  password: admin
- После этого вам будет доступен функционал приложения:
![image](https://github.com/MaksimovDenis/gRPC-server-HansZimmerSoundtracks/assets/44647373/3e170ccb-2397-4be3-a83f-8e82016b76fd)
![image](https://github.com/MaksimovDenis/gRPC-server-HansZimmerSoundtracks/assets/44647373/a9e28ee4-a742-4f88-90e1-0a9c149a7483)
  
## Технологии и зависимости
- Язык программирование: Golang 1.21.6  
- SQLite
- yandexStorage
- Docker

Основные зависимости:  
github.com/aws/aws-sdk-go-v2: Библиотека для работы с AWS SDK в языке Go, предоставляющая доступ к различным сервисам AWS.  
github.com/brianvoe/gofakeit/v6: Библиотека для генерации случайных данных, полезная для тестирования и заполнения фейковыми данными.  
github.com/fatih/color: Пакет для работы с цветовой выдачей текста в терминале.  
github.com/golang-jwt/jwt: Реализация JWT (JSON Web Tokens) в языке Go.  
github.com/golang-migrate/migrate/v4: Инструмент для управления миграциями баз данных.  
github.com/grpc-ecosystem/go-grpc-middleware/v2: Библиотека для добавления промежуточного ПО (middleware) в gRPC-серверы и клиенты.  
github.com/ilyakaznacheev/cleanenv: Пакет для загрузки конфигураций из переменных окружения, файла .env или других источников.  
github.com/mattn/go-sqlite3: SQLite драйвер для языка Go.  
github.com/stretchr/testify: Библиотека для тестирования на языке Go.  