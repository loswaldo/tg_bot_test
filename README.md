# tg_bot_test

telegram bot for weather. 
Using [telegram API](https://go-telegram-bot-api.dev/getting-started/index.html) and [Tomorrow.io weather API](https://www.tomorrow.io/weather-api/)

For start application you need to create two yml config files 
1. api-config.yml file
    ```yaml
    weather_api_key: "" //you're weather api key
    tg_api_key: "" //you're telegram bot api tocken
    ```
2. db-config.yml file
    ```yaml
    host: ""
    port: ""
    user: ""
    password: ""
    db_name: ""
    ssl_mode: ""
    ```
   
And after that you just write 
    ```bash
    sudo docker-compose up  
    ```