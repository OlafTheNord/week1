## Неделя 1

#### <span style="color:darkorange;">Описание задачи:</span>

* Обучение JSON
* Работа с OS
* Базовый синтаксис
* Map, Struct




#### <span style="color:darkorange;">Модель данных:  </span>

| Поле | Тип | Описание | Обязательность |
| --- | --- | --- | --- |
| Name | string | ФИО сотрудника | false |
| Department | string | Название департамента | true |
| Salary | float64 | Зарплата | \> = 0 |
| HireDate | string | Дата найма | false |

#### Пример данных:

```
[
  {
    "name": "Иван Иванов",
    "department": "IT",
    "salary": 120000,
    "hire_date": "2022-01-15"
  }
]
```

#### 
<span style="color:darkorange;">Запуск приложения:</span>

go run .

**Файл с сотрудниками должен быть в формате JSON с наименованием employees.json в корневом каталоге**



#### <span style="color:darkorange;">Результат выполнения и возможные ошибки:</span>

```
Положительный результат

Avg of salaries of dept: 
IT: 115000.00
HR: 92500.00
Finance: 105000.00
```



```
Ошибки:
error in handlin - невозможно запустить сервис, проверьте корректность приложения
cannot read file - невозможно открыть и прочитать json, проверьте что файл назван без ошибок и не "битый"
cannot unmarshal file - ошибки данных в json`е
employee salary is negative - отрицательная зарплата сотрудника
employee department is empty - не заполнен департамент у сотрудника
```

(С) OLAFTHENORD (FILLBERRYHILLS LLC)