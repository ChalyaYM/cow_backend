# ВАЖНО

На момент 20 фев. 2025 года ДЕВ сервер (genmilk.ru:44300) запускается файлом compose_kafka.yaml ```bash docker compose -f compose_kafka.yaml up -d --build``` ПРОД сервер (genmilk.ru) запускается файлом compose.yaml ```bash docker compose up -d --build```


# ЗАПУСК ПРОЕКТА  

Перед тем как запустить проект, необходимо создать .env файл в папке backend. Шаблон .env файла представлен в template.env. С шаблонным .env будет работать все, кроме почты, т.к. я не стал коммитить пароль для доступа к личной почте.  

Для того, чтобы локально запустить проект необходимо иметь [установленный докер](https://docs.docker.com/engine/install/) и из корня репозитория выполнить команду  

``` bash
docker compose up -d --build
```

После выполнения команды:  

- веб-сервис будет доступен по адресу localhost,  
- напрямую на бэкэнд можно постучаться по адресу localhost:8080,  
- по адресу localhost:5050 можно получить доступ к pgadmin.  

# ПЕРВЫЙ ПОЛЬЗОВАТЕЛЬ  

Для того чтобы создать пользователя нужно зайти в админку, для того чтобы зайти в админку нужно иметь логин и пароль админа, логин и пароль админа не создаются при первом запуске. На данный момент эта проблемма не имеет нормального решения. Можно сделать дамп БД с сайта genmilk.ru:5050 и пользоваться админом оттуда.  

# СТРУКТУРА ПРОЕКТА  

## DOCKER  

Докер контейнеры запускаются с помощью утилиты [docker-compose](https://docs.docker.com/compose/), так как для покрытия необходимых на первом этапе потребностей в "оркестрации" контейнеров функционала docker-compose более чем достаточно.  

С помощью docker-comose запускаются следующие контенеры:  

- контейнер с СУБД postgresql  
- контейнер с pgadmin  
- контейнер с backend
- контейнер с nginx + frontend  

Контейнеры с backend и frontend описаны далее.  
  
Контейнер с СУБД postgresql настроен так, что при инициализации (происходит при самом первом запуске контейнера, тогда, когда база данных еще ни разу не запускалась) выполняются скрипты из папки dockered_db/docker-entrypoint-initdb.d. Если остановить контейнер с postgresql, то данные БД не удалятся, для этого используется volume habrdb-data. База данных не доступна вне контейнера, т.к. на сервере занят порт postgresql.  
  
Pgadmin доступен откуда угодно, порт 5050. Также можно получить доступ к Pgadmin сервера по адресу genmilk.ru:5050

## FRONTEND  

Как работает фронтенд можно спросить у Юлии Чалой.  
  
Контейнер фронтенда сначала собирает vue в статики, затем запускает nginx, для запуска nginx в папке frontend/nginx должны быть два файла genmilk.crt и genmilk.key (сертификат и ключ для https).  
  
Для разработки бэкэнда nginx был не нужен, поэтому бэкэндеры создали два пустых файла с такими названиями, у бэкэндеров упал nginx на этапе запуска, но на сервере всё работало, а руты тестировались по url localhost:8080. Откуда взялись ключи у фронтендера и были ли они вообще можно спросить у Юлии Чалой.

## BACKEND  

Докер-контейнер бэкэнда работает очень просто: собирает бэкэнд в исполняемый файл, этот исполняемый файл копирует в контенер с линуксом, к исполняемому файлу копирует шаблоны и .env файл с переменными окружения.  

### .env  

В .env файле описаны переменные окружения, имеющие следующее значение:  

- DB_HOST - адрес сервиса базы данных
- DB_UESR - пользователь базы данных  
- DB_PASSWORD - пароль для пользователя базы данных  
- DB_NAME - имя бд  
- DB_PORT - порт, по которому база данных отвечает на запросы  
- JWT_KEY - ключ шифрования, использующийся для авторизации  
- PORT - порт, прослушиваемый сервером  
- EMAIL_FROM - адрес электронной почты, с которого отправляются сообщения genmilk
- EMAIL_PASS - пароль использующийся для авторизации пользователя (протокол smtp)
- SMTP_HOST - адрес SMTP сервера, исопльзующегося для отправки сообщений  
- SMTP_PORT - порт SMTP сервера  

### Исользуемые библиотеки  

- [gin](https://gin-gonic.com/) используется для обеспечения работы rest-api  
- [gorm](https://gorm.io/) ORM. Все взаимодействия с базой данных происходят с её помощью.  
- [swaggo](https://github.com/swaggo/swag) библиотека для генерации сваггера  

### Запуск вне докер контейнера  
  
Бэкэнд можно запустить вне докер контейнера, например в дебаггере предпочтительной IDE. Для этого, нужно в .env файле изменить переменные окружения DB_HOST и DB_PORT на хост и порт СУБД postgresql в системе. Так как контейнер с БД не отвечает на запросы извне докер контейнера нужно либо использовать свою postgresql (**сложный вариант**), либо в файле compose.yaml раскомментировать строки, открывающие контейнер с бд для внешенего мира (**простой вариант**: DB_HOST=localhost DB_PORT=5432).  

### Руты  

Код для обработки запросов пользователя находится в server/routes. Для каждой группы рутов создана отдельная папка, так код, обрабатывающий /api/cows/* находится в backend/routes/cows.  

Фильтрация коров нужна как в аналитике, так и во вкладке "животные", поэтому код для формирования запроса к базе данных с пользовательскими фильтрами вынесен отдельно и находится в backend/filter. Для фильтрации используется BaseFilteredModel (backend/filters/filter.go), содержащая в себе запрос к БД, к которому дописываются условия для каждого фильтра и параметры фильтрации. От BaseFilteredModel наследуется CowFilteredModel (backend/filters/cows_filter/cow_filtered_model.go), принимающая в себя данные из post-запроса к фильтрам и запрос к БД см. NewCowFilteredModel (backend/filters/cows_filter/cow_filtered_model.go). Каждое поле фильтра обрабатывается отдельным объектом фильтром (объектом, реализующим интерфейс Filter из файла backend/filters/cows_filter/filter.go), в методе Apply проверяются входные условия фильтрации и дописывается условие в запрос, если необходимо. Функция ApplyFilters применяет все переданные в неё объекты фильтров к запросу, после чего методом GetQuery можно получить готовый запрос к БД.  

### Модели  

Модели GORM описывают таблицы в базе данных, внешние ключи, ограничения и действия выполняемые ORM до/после добавления записи в базу данных.  Описания моделей находятся в backend/models

Как работает GORM можно посмотреть [тут](https://gorm.io/)  

### Шаблоны  

В проекте есть второй фронтенд, реализованный с помощью шаблонов jinja2, находящихся в backend/templates. Фронтенд используется для администрирования системы, поэтому возможности vue для его реализации избыточны. Руты отправляющие фронтенд админа описаны в backend/routes/admin.

### Тонкая логика аналитики  

Изначально аналитика была сделана без фильтров и get рутами, затем с фильтрами и post рутами. Появилось сразу 2 группы рутов, которые делают одно и то же.  
От get рутов не отказались, т.к. на главной странице сайта должна быть ровно та же аналитика, что в этих гет-рутах и эта аналитика должна работать. При этом требовались разные права доступа к аналитике главной страницы и аналитике вкладки "Аналитика".  
Post руты выполняют свою прямую задачу и используются в аналитике, как и должны.

### Даты  

Вечная боль... Отправляются как строки формата "ГГГГ-ММ-ДД", потому что так отправляет библиотека фронтенда и принимает голанг. swaggo не понимает что именно будет отправлено, поэтому в сваггере про даты написана неверная информация.

### Вложенные объекты  

Часть рутов бэкэнда отправляет вложенные объекты. Эти объекты лишний раз не селектятся из БД, поэтому в большинстве рутов не отправляются в ответе. Это не значит, что в описании рутов, которые отправляют вложенные объекты об этом что-то написано, поэтому лучше использовать postman и проверять.  

### Загрузка файлов  

Не описана в сваггере, т.к. не понятно как. Требует авторизации. Как работает можно посмотреть в папке backend/templates (в файлах *Load.tmpl лежат готовые html формы).  

### Авторизация  

/auth/login возвращает JWT токен, который должен быть в заголовке "Authorization" каждого запроса.
