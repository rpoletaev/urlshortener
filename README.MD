
## Запуск
```
docker-compose up --build
```

## примеры curl: 
### создание ссылки:
```
curl 'localhost:3001/link/' -H 'Content-Type: application/json' -X POST --data-raw '{"link": "https://https://www.mynimo.com/cebu/cebu-city-jobs\\?__cf_chl_captcha_tk__\\=dca8bc35201470968c1d90772947c37c80cb8a05-1606209761-0-AQL4Vuw5vXMGTwmt5uuMURwStxuzdx0qIQnm_bltdvNetM_9mEbdr0ctxjdCZHSYjTixtJ21CQIe68RAGnCB-CtiR5lAzkskul--yjUtf6VApw76Vmi2iTu2Ny7XQNHiKFiSS73NA0nDTJgDgFGblsTDrQUchsIfj4euPa5ackhq59tZDvydlK-QEE3SBsic7ubHgIs2Wa8IYN3L4ynB8-Kigaz9Bss6BPsgybUSeQfMaZoPa_cowDBRGapZTviZOepMjFLltfiKVEuXNgypVGWdFVGkw2YXNxOlcV_C-Fv5c7IMAtxoevaKYly5Uq5ibaYNa_9mCOIV_LcZxxEMo-YgWFBq6_7w199bvccuC9pNmOet8s_D43xARW5HRvAH0FNPC-inDCcDMfwX_H306q9L1lKkHYCTxags2lBUEhjB1rDZSnlYQPklm8vXt9Os6LFzCQV6sC7-cTpcS8RBXD-Ky7jF9AxGjrzJ5qTB0RrAyErTrbWeW1--S7xbkfRu_9TqXz5ZpnVrYhjtxjV2kruG2QNW5hv3UXxLKV9tgkK5"}'
```
### возвращает:
    {"ShortLink":"short/kB83"} //имя домена, в данном случае short можно конфигурировать в переменной окружения 'SERVICE_DOMAIN'

### редирект по ссылке:
```
    curl 'localhost:3001/link/kB83'
```

### мерил с помощью:
    https://github.com/americanexpress/baton

### Пути для статистики
/stat/ip
/stat/url