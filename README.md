# API Parser OZON
Выгрузка товаров в плоскую структуру в json файл из магазина OZON.
Документация для получения API-ключ: https://docs.ozon.ru/api/seller/#tag/Auth

### 1. Параметры(Flag) приложения
- `--apiKey` (string) Данные Header Api-Key
- `--clientId` (string) Данные Header Client-Id
- `--method` (string) Метод запроса (default: "POST")
- `--url` (string) URL API (default: "https://api-seller.ozon.ru/v2/category/tree")
- `--expFile` (bool) Выгружает json файл товаров (default: false)

### 2. App Docker build and run
- Build: `sudo docker build -t ozon .`
- Run: `sudo docker run ozon ./app --apiKey=... --clientId=...`

### 3. App run 
Run: `go run cmd/api/main.go --apiKey=... --clientId=...`
