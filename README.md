Здравствуйте! Меня зовут Евгений Щербаков, я студент группы ИВТАСбд-32(42) УЛГТУ. Это моё тестовое задание GO: вариант с автомобилями.
Я выполнил это задание, стараясь соблюдать чистую архитектуру.
Использовал GoLand c версией Go 1.22. Для запуска проекта потребуется установить Docker, там у меня разворачивается база postgres.
В терминал PowerShell нужно будет ввести следующие команды:

1)docker run --name=CarAPI -e POSTGRES_PASSWORD='1234' -p 5436:5432 -d --rm postgres

2)migrate -source ./migr -database 'postgres://postgres:1234@localhost:5436/postgres?sslmode=disable' up

3)go run cmd/main.go

Результаты работы программы:
Добавление:

<img width="541" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/eaca8909-e760-4818-a1ff-733188335161">
<img width="146" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/fdae953d-d055-48a4-aebd-cfd7606b2ab6">

Результат и в свою очередь вывод всего списка(GET):

<img width="263" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/0d41a6cd-cff0-4fea-b95d-c273bee45709">

Редактирование:

<img width="263" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/2122ad88-cdda-4831-af36-0aefa8a87c93">
<img width="261" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/ac1db384-db5f-43cf-922b-bd80b552ddde">

Удаление:

<img width="372" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/1e50a177-f779-411c-8566-0f2f36bae0b7">
<img width="190" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/5dac5834-d74f-4a2c-aeeb-43a38c1e7fdd">
<img width="217" alt="image" src="https://github.com/fgmpa/GolangAPI/assets/90002774/77e23f9b-30bb-42bd-9424-e46a5b2b7c9b">







