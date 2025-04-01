# Till bot

Till bot — телеграм-бот, сообщающий сколько времени осталось до того или иного события.

## Установка и запуск
Установите токен бота из [@botfather](https://t.me/botfather) в переменную среды окружения `BOT_TOKEN`.
```sh
export BOT_TOKEN=ваш_токен
```
Склонируйте репозиторий и соберите бота.
```sh
git clone https://github.com/Akionka/tillbot.git
cd tillbot
go build .
./tillbot
```
