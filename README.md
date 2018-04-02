# Парсер комментариев с classic-online.ru
## Usage
1. Скопируйте шаблоны файлов с настройками
```
cd config
cp database.yml.example database.yml
cp headers.yml.example headers.yml
cd ..
```
Положите свои cookie в config/headers.yml, послед авторизации в браузере
2.
```
go get -v
go run main.go
```
