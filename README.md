В данном разделе описана только установка проекта.

С условием задачи и комментариями к решению можно ознакомиться [здесь](https://github.com/plutonio00/pay_task_go/tree/master/docs)

## Установка

Для установки необходимо наличие docker и docker-compose.

Для развёртывания проекта в UNIX-подобных ОС необходимо выполнить следующие действия:

1. `cd path-to-project`
2. `cp .env.example .env`
3. `cd docker`
4. `cp .env.example .env`
5. `docker-compose build --no-cache`
6. `docker-compose up -d`
7. `docker exec go-pay make`
8. `docker exec go-pay make db_init`
9. `docker exec go-pay make exec`

В результате:
1. Будут созданы файлы env для docker и проекта
2. Будут собраны и запущены docker-контейнеры
3. Будет произведена сборка go кода
4. Накатятся миграции
5. Будет запущен сервер

Кроме того, при сборке docker-контейнера для postgres будут автоматически созданы база данных для проекта и пользователь с правами на нее.

После выполнения данных действий проект будет доступен по url `http://127.0.0.1:8080/`

Swagger документация будет доступна по адресу `http://127.0.0.1:8080/swagger/index.html`