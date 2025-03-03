# 2025-q1-practice


Warehouse Management System

Описание

Этот проект представляет собой систему управления складом, написанную на Golang с использованием PostgreSQL в качестве базы данных. API предоставляет возможности управления складами, продуктами, запасами и аналитикой продаж.

Стек технологий

Backend: Golang

Database: PostgreSQL

Router: Gorilla Mux

ORM: pgx (PostgreSQL driver)

Containerization: Docker

Запуск проекта

1. Клонирование репозитория

git clone https://github.com/RUST-GOLANG/2025-q1-practice.git
cd warehouse-api

2. Настройка базы данных

Создайте базу данных PostgreSQL и настройте подключение в config.env:

DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=warehouse_db

3. Запуск с Docker

docker compose up -d

4. Локальный запуск без Docker

Установка зависимостей

go mod download

Запуск приложения

go run cmd/main.go

API Эндпоинты

Управление складами

POST /api/warehouses – Создать склад

GET /api/warehouses – Получить список складов

Управление продуктами

POST /api/products – Добавить продукт

GET /api/products – Получить список продуктов

Управление запасами

POST /api/inventory – Добавить товар на склад

PUT /api/inventory/update – Обновить количество товаров

Аналитика

GET /api/analytics/warehouse – Получить аналитику по складам

GET /api/analytics/top-warehouses – Топ-склады по продажам
