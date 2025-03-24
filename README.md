# go-stomp-websocket

Golang имплементация STOMP-клиента поверх websocket

#### Поддержаны операции: 
* Установление STOMP соединения
* Подписка на события

#### Использование:
 
 Для того, чтобы начать использование STOMP-клиента необходимо:
 1. Задать путь для установления подключения
 2. Создать Stomp-client используя токен или используя пользовательский Dial
 3. Задать канал получения фреймов
 
#### Пример

Для подключения к Watch API Tenant-Manager'а, работающего по протоколу STOMP имеющий путь:
```
ws://tenant-manager:8080/api/v3/tenant-manager/watch
```

#### Создаем STOMP-клиент
##### C использованием токена
```
token, _ := tenantWatchClient.Сredential.GetAuthToken()
url, _ := url.Parse("ws://localhost:8080/api/v3/tenant-manager/watch")
stompClient, _ := go_stomp_websocket.ConnectWithToken(*url, token)
```

##### C пользовательским Dial
```
type ConnectionDialer interface {
    Dial(webSocketURL url.URL, dialer websocket.Dialer, requestHeaders http.Header) (*websocket.Conn, *http.Response, error)
}
```
```
url, _ := url.Parse("ws://localhost:8080/api/v3/tenant-manager/watch")
dialer := websocket.Dialer{}
//конфигурация dialer
requestHeaders http.Header{}
//добавление хейдеров
connDial ConnectionDialerImpl{} //имлементирует метод Dial интерфейса ConnectionDialer 
stompClient, _ := go_stomp_websocket.Connect(*url, dialer, requestHeaders, connDial)
```

Подписываемся на события
```
subscr, _ := stompClient.Subscribe("/tenant-changed") 
```

Реагируем на получаемые фреймы
```
go func() {
    for {
        var tenant = new(tenant.Tenant)
        frame := <-subscr.FrameCh // Получаем фрейм
        if len(frame.Body) > 0 {
            err := json.Unmarshal([]byte(frame.Body), tenant) // Преобразуем тело фрейма в структуру Tenant
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Printf("Received tenant with id:%s\n", tenant.ObjectId)
            }
        }
    }
}()
```