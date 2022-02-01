# TestTask
Выгрузка товаров в json файл (ozon.json) из магазина OZON

### 1. Параметры(Flag) приложения
- `--apiKey` (string) Данные Header Api-Key
- `--clientId` (string) Данные Header Client-Id
- `--method` (string) Метод запроса (default: "POST")
- `--url` (string) URL API (default: "https://api-seller.ozon.ru/v2/category/tree")
- `--expFile` (bool) Выгружает json файл товаров (default: false)

### 1. App Docker build and run
- Build: `sudo docker build -t ozon .`
- Run: `sudo docker run ozon ./app --apiKey=... --clientId=...`

### 2. App run 
Run: `go run cmd/api/main.go --apiKey=... --clientId=...`
