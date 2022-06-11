# Тестовое задание Континуум
***
Репозиторий содержит реализацию 
[тестового задания](https://docs.google.com/document/d/1T7D0vDoxE2aBE37lxp22vPxlgZNLd3eg/edit)
на летнюю практику (бекенд)

**Запуск**

Для запуска требуется выполнить следующие шаги:

**1) Установить себе язык Golang**
```
$ sudo apt install golang

$ go version
```
**2) Установить себе драйвер для работы Golang с SQL**

```
$ go get -u github.com/go-sql-driver/mysql
```

**3) Установить себе библиотеку [Fyne](https://developer.fyne.io/started/)**

```
$ go get fyne.io/fyne/v2
```
>Может понадобится прописать
>```
>$ go mod tidy
>```

**4) Скачать исходный код и перейти в директорию с проектом**
```
$ git clone https://github.com/OkDenAl/kontinuumTestTask.git
$ cd kontinuumTestTask
```

**5) Создать базу данных `continuum` в MySQL**

**6) Заполнить ее, скопировав содержимое [файла](https://github.com/OkDenAl/kontinuumTestTask/blob/main/sql/continuum.sql)**

**7) Запустить программу**
```
$ go run ./cmd/app/main.go
```
Опционально программу можно запускать через `kontinuum.exe`  
Также, если будет желание самому собрать `.exe` после изменений,
то это можно сделать прописав 
```
fyne package -os linux -icon <путь до картинки>
fyne package -os windows -icon <путь до картинки>
```
***
## Немного о файлах проекта

`main.go` - запускает графический интерфейс для работы с программой

`sqlHandler.go` - вся логика программы, связанная с sql и обработкой данных с sql

В целом, комментарии в основном коде я оставил, думаю, будет понятно, что за что отвечает, надеюсь...

***

Остались вопросы? [Пиши](https://vk.com/d.okutin)  
Буду рад ответить :) И выслушать конструктивную критику
