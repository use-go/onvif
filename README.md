# Goonvif
Библиотека **Goonvif** создана для упрощения взаимодействия с ONVIF устройствами. На данный момент в библиотеке реализована поддержка NVT(Network Video Transmitter) устройств, а именно следующих ONVIF сервисов:
- Core или DeviceManagement
- Media
- Imaging
- PTZ
- Analytics
# Dependencies
[etree](https://github.com/beevik/etree)
# Установка
Для установки библиотеки необходимо воспользоваться утилитой go get:
```
go get github.com/yakovlevdmv/goonvif
```

# Начало работы
Чтобы начать работать с камерой, необходимо создать объект `device`.
Для этого необходимо воспользоваться функцией `func NewDevice(xaddr string) (*device, error)`,
которая принимает адрес ONVIF устройства и возвращает указатель на созданный объект либо ошибку.
Если камера не доступна, указан неверный адрес для ONVIF сервиса камеры (возможно находся по другому порту) или же камера вообще не поддерживает ONVIF
функция вернет не `nil` error, а в качестве указателя на объект вернет `nil`.
### Пример подключения к камере
Пусть камера в сети находится по адресу 192.168.13.42, а ее ONVIF сервисы расположены на 1234 порту. Тогда,
```
dev, err := goonvif.NewDevice("192.168.13.42:1234")
```
сработает успешно, а
```
dev, err := goonvif.NewDevice("192.168.13.42:80")
```
вернет нулевой объект камеры и ошибку:
> camera is not available at 192.168.13.42:80 or it does not support ONVIF services

Модернизируем код, добавив обработку ошибки, и получим:
```
dev, err := goonvif.NewDevice("192.168.13.42:80")
if err != nil {
    panic(err)
}

///Работа с камерой
```
### Поддерживаемые ONVIF сервисы
Теперь, когда камера доступна, можно приступать к работе с ней. Однако стандарт ONVIF имеет множество сервисов, точка доступа (endpoint) к которым не закреплена стандартом (кроме DeviceManagment: http://onvif_host/onvif/device_service).
Поэтому дальше встает вопрос о поддерживаемых камерой сервисах и определении их endpoint'ов.
Для получения поддерживаемых камерой сервисов необходимо вызвать метод GetCapabilities сервиса DeviceManagement.
Однако данная библиотека автоматизирует данный процесс, поэтому при создании объекта device при помощи `func NewDevice(xaddr string) (*device, error)`
библиотека одновременно обрабатывает поддерживаемые камерой сервисы. Таким образом чтобы получить поддерживаемые устройством сервисы, можно воспользоваться двумя путями:
1. Вызвать метод GetCapabilities сервиса DeviceManagement(как это сделать будет рассмотрено дальше) и обработать ответ.
2. Довериться библиотеке и вызвать функцию  `func (dev *device)GetServices() map[string]string`, которая вернет map[string]string, ключом которой является название сервиса, а значением - endpoint данного сервиса
### Работка с камерой
Для работы с различными сервисами камерами необходимо отправить корректный SOAP запрос, в теле которого находится вызываемый метод и принимаемые им функции.
**Goonvif** берет на себя работу по созданию корректного SOAP запроса и его отправке. В **Goonvif** определены структуры, для каждой функции каждого (поддерживаемого данной бибилиотекой) сервиса ONVIF:
- [DeviceManagement Service](Device/types.go)

- [Media Service](Media/types.go)

- [Imaging Service](Imaging/types.go)

- [PTZ Service](PTZ/types.go)

- [Analytics Service](Analytics/types.go)

[Список всех сервисов стандарта (и документация к ним)](https://www.onvif.org/profiles/specifications/)

Рассмторим, как организована отправка запросов в **Goonvif** на нескольких примерах.
1. Метод GetCapabilities сервиса DeviceManagement
Все необходимые типы данных определены в пакете [Device](Device/types.go).
В файле (https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl) можно увидеть:
![GetCapabilities](img/exmp_GetCapabilities.png)

Таким образом, Функция GetCapabilities принимает в качестве аргумента перечисление:
`enum { 'All', 'Analytics', 'Device', 'Events', 'Imaging', 'Media', 'PTZ' }`
Чтобы вызвать данный метод создадим объект `Device.GetCapabilities`:
```
capabilities := Device.GetCapabilities{Category:"All"}
```
Для вызова данной функции воспользуемся методом `func (dev device) CallMethod(method interface{}) (string, error)`:
```
resp, err := dev.CallMethod(capab)
if err != nil {
    log.Println(err)
} else {
    fmt.Println(resp)
}
```
2. Создание пользователя методом CreateUsers сервиса DeviceManagement
Все необходимые типы данных определены в пакете [Device](Device/types.go).
В файле (https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl) можно увидеть структуру запроса:
![GetCapabilities](img/exmp_CreateUsers.png)

Создадим объект `Device.CreateUsers`:
```
createUsers := Device.CreateUsers{User: onvif.User{Username:"korolev", Password:"qwerty", UserLevel:"User"}}
resp, err := dev.CallMethod(createUsers)
if err != nil {
	log.Println(err)
} else {
	fmt.Println(resp)
}
```

В данном примере можно наблюдать использования пакета onvif, в котором определено большинство типов, используемых в различных сервисах.
Поэтому при создании структур запросов необходимо это учитывать.

#####**ВАЖНО**
Некоторые камеры работают специфично. Это означает, что в зависимости от модели камеры можно не получить ошибки при неправильном запросе. Поэтому советую проверять точно ли выполнилась операция. Например, для метода CreateUsers вывести список всех пользователей и сравнить.
