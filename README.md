# ⚡ flashdb

**FlashDB** — это высокопроизводительное in-memory key-value хранилище, написанное на Go.  
Проект вдохновлён Redis и создан для изучения конкурентности, сетевого взаимодействия и разработки высоконагруженных систем.

---

## 🚀 Возможности

- ⚡ Хранение данных в памяти (in-memory)
- 🔑 Key-value модель
- 🌐 TCP сервер с поддержкой клиентских подключений
- 📦 Поддержка базового протокола (RESP)
- 🧵 Конкурентная обработка запросов (goroutines)
- 🔒 Потокобезопасное хранилище

---

## 🛠 Поддерживаемые команды (MVP)

- `PING` — проверка соединения
- `SET key value` — сохранить значение
- `GET key` — получить значение
- `DEL key` — удалить ключ

---

## ⚙️ Запуск

```bash
go run ./cmd/server
```

После запуска появится интерактивный CLI с приглашением `flashdb>`.

---

## 📖 Использование

### Доступные команды CLI:

- `ping` — проверить соединение с сервером
- `set <key> <value>` — сохранить значение по ключу
- `get <key>` — получить значение по ключу
- `del <key>` — удалить ключ
- `exit` или `quit` — выйти из приложения

### Примеры:

```bash
flashdb> ping
PONG

flashdb> set mykey hello world
OK

flashdb> get mykey
hello world

flashdb> del mykey
(integer) 1

flashdb> get mykey
(nil)
```

---

## 🏗 Архитектура

- **`cmd/server/main.go`** — точка входа сервера и CLI
- **`internal/server/server.go`** — TCP сервер
- **`internal/handler/connection.go`** — обработка клиентских подключений и команд
- **`internal/storage/storage.go`** — потокобезопасное in-memory хранилище
- **`cmd/cli/app.go`** — интерактивный CLI
