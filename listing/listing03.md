Чем отличаются RWMutex от Mutex?

Ответ:
```
RWMutex (Reader-Writer Mutex) и Mutex (Mutual Exclusion Lock) 
- это два типа блокировок в Go, которые обеспечивают безопасность 
доступа к разделяемым данным в многопоточной среде.Основные отличия 
между ними заключаются в их поведении при доступе к ресурсу:

Mutex:
Обеспечивает эксклюзивный доступ к ресурсу только одной горутине 
в любой момент времени.
Другие горутины, желающие получить доступ к ресурсу, должны ждать, 
пока текущая горутина не освободит блокировку Mutex.
Используется, когда необходимо гарантировать, что только одна горутина 
имеет доступ к критической секции кода в определенный момент времени.

RWMutex:
Позволяет множественным горутинам получать доступ к ресурсу для 
чтения одновременно.
Однако, при записи (изменении) ресурса блокируется доступ как 
для чтения, так и для записи другим горутинам.
Используется в ситуациях, когда часто происходят операции чтения, 
но редко происходят операции записи.
```