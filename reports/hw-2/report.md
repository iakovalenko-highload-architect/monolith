# Индекс
Использовался индекс `GIN`, поскольку позволяет эффективно искать с оператором `LIKE`
```sql
create index first_name_second_name on users using gin(first_name gin_trgm_ops, second_name gin_trgm_ops);

explain select id, password, first_name, second_name, birthday, city, biography
        from users
        where first_name LIKE '%рам%' and second_name LIKE '%дам%'
        order by id

Sort  (cost=36.03..36.03 rows=1 width=147)
      Sort Key: id
  ->  Bitmap Heap Scan on users  (cost=32.00..36.02 rows=1 width=147)
        Recheck Cond: ((first_name ~~ '%рам%'::text) AND (second_name ~~ '%дам%'::text))
        ->  Bitmap Index Scan on first_name_second_name  (cost=0.00..32.00 rows=1 width=0)
              Index Cond: ((first_name ~~ '%рам%'::text) AND (second_name ~~ '%дам%'::text))

```

# Результаты замеров

## Без индекса

### 1-10 одновременных запросов
#### Количество одновременных запросов
![1-10-1.png](./images/1-10-1.png)
#### Время ответа
![1-10-2.png](./images/1-10-2.png)
#### Пропускная способность
![1-10-3.png](./images/1-10-3.png)

### 10-100 одновременных запросов
#### Количество одновременных запросов
![10-100-1.png](./images/10-100-1.png)
#### Время ответа
![10-100-2.png](./images/10-100-2.png)
#### Пропускная способность
![10-100-3.png](./images/10-100-3.png)

### 100-1000 одновременных запросов
#### Количество одновременных запросов
![100-1000-1.png](./images/100-1000-1.png)
#### Время ответа
![100-1000-2.png](./images/100-1000-2.png)
#### Пропускная способность
![100-1000-3.png](./images/100-1000-3.png)

## С индексом
### 1-10 одновременных запросов
#### Количество одновременных запросов
![1-10-2-1.png](./images/1-10-2-1.png)
#### Время ответа
![1-10-2-2.png](./images/1-10-2-2.png)
#### Пропускная способность
![1-10-2-3.png](./images/1-10-2-3.png)

### 10-100 одновременных запросов
#### Количество одновременных запросов
![10-100-2-1.png](./images/10-100-2-1.png)
#### Время ответа
![10-100-2-2.png](./images/10-100-2-2.png)
#### Пропускная способность
![10-100-2-3.png](./images/10-100-2-3.png)

### 100-1000 одновременных запросов
#### Количество одновременных запросов
![100-1000-2-1.png](./images/100-1000-2-1.png)
#### Время ответа
![100-1000-2-2.png](./images/100-1000-2-2.png)
#### Пропускная способность
![100-1000-2-3.png](
./images/100-1000-2-3.png)
