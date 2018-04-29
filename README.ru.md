# Goonvif
Простое управление IP-устройствами, включая камерами. Goonvif - это реализация протокола ONVIF для управления IP-устройствами. Целью создания данной библиотеки является удобное и легкое управление IP-камерами и другими устройствами, поддерживающими стандарт ONVIF.

## Установка
Для установки библиотеки необходимо воспользоваться утилитой go get:
```
go get github.com/yakovlevdmv/goonvif
```
## Поддерживаемые сервисы
Следующие сервисы полностью реализованы:
- Device
- Media
- PTZ
- Imaging

## Использование

### Общая концепция
1) Подключение к устройству
2) Аутентификация (если необходима)
3) Определение типов данных
4) Выполнение необходимого метода

#### Подключение к устройству
Если в сети находится устройство по адресу *192.168.13.42*, а ее ONVIF сервисы используют порт *1234*, тогда подключиться к устройству можно следующим способом:
```
dev, err := goonvif.NewDevice("192.168.13.42:1234")
```

*ONVIF порт может отличаться в зависимости от устройства и, чтобы узнать, какой порт использовать, можно зайти в веб-интерфейс устройства. **Обычно это 80 порт.***

#### Аутентификация
Если какая-либо функция одного из сервисов ONVIF требует аутентификацию, необходимо использовать метод `Authenticate`.
```
device := onvif.NewDevice("192.168.13.42:1234")
device.Authenticate("username", "password")
```

#### Определение типов данных
Каждому сервису ONVIF в этой библиотеке соответствует свой пакет, в котором определены все типы данных этого сервиса, причем название пакета идентично названию сервиса и начинается с заглавной буквы.
В Goonvif определены структуры для каждой функции каждого поддерживаемого этой библиотекой сервиса ONVIF.
Определим тип данных функции `GetCapabilities` сервиса `Device`. Это делается следующим образом:
```
capabilities := Device.GetCapabilities{Category:"All"}
```
Почему у структуры GetCapabilities поле Category и почему значение этого поля All?

На рисунке ниже показана документация функции [GetCapabilities](https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl). Видно, что функция принимает один пареметр Category и его значение должно быть одно из следующих:  `'All', 'Analytics', 'Device', 'Events', 'Imaging', 'Media' или 'PTZ'`. 

![Device GetCapabilities](img/exmp_GetCapabilities.png)

Пример определения типа данных функции GetServiceCapabilities сервиса [PTZ](https://www.onvif.org/ver20/ptz/wsdl/ptz.wsdl):
```
ptzCapabilities := PTZ.GetServiceCapabilities{}
```
На рисунке ниже видно, что GetServiceCapabilities не принимает никаких аргументов. 

![PTZ GetServiceCapabilities](img/GetServiceCapabilities.png)

*Общие типы данных находятся в пакете xsd/onvif. Типы данных (структуры), которые могут быть общими для всех сервисов определены в пакете onvif.*

Пример оределения типа данных функции CreateUsers сервиса [Device](https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl):
```
createUsers := Device.CreateUsers{User: onvif.User{Username:"admin", Password:"qwerty", UserLevel:"User"}}
```

По рисунку ниже видно, что в данном примере полем структуры CreateUsers должен быть User, типом данных которого является структура User, содержащая поля Username, Password, UserLevel и необязательный Extension. Структура User находится в пакете onvif.

![Device CreateUsers](img/exmp_CreateUsers.png)

#### Выполнение необходимого метода
Для выполнения какой-либо функции одного из сервисов ONVIF, структура которой была определена, необходимо использовать `CallMethod` объекта device.
```
createUsers := Device.CreateUsers{User: onvif.User{Username:"admin", Password:"qwerty", UserLevel:"User"}}
device := onvif.NewDevice("192.168.13.42:1234")
device.Authenticate("username", "password")
resp, err := dev.CallMethod(createUsers)
```
