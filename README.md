# WorkerPool для стажировки VK

Данная программа является реализацией примитивного worker-pool с возможностью динамически добавлять и удалять работников в pool в процессе обработки данных

## Установка и запуск

1. Необходимо клонировать репозиторий:
```
git clone https://github.com/sashaem1/WorkerPool_VK.git
```
2. Скомпилировать и запустить программу:
```
make run
```

## Описание

Данная программа создаёт WorkerPool на 6 работников, запускает обработку данных из файла (работники выводят строки из файла и ожидают пол секунды, имитируя работу). Далее через 2 секунды три работника завершают свою работу, а ещё через 2 секунды создаются два новых работника и подключаются к рабочему процессу

## Изменение обрабатываемых данных

Данные генерируются с помощью программы, находящейся в cmd/makeData, если необходимо изменить количество обрабатываемых строк, можно изменить переменную N на необходимое значение и запустить программу:
```
go run cmd/makeData/main.go
```