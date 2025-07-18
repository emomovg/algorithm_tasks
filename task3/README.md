# Карточки

## Условие задачи

Среди ваших 𝑛 друзей стало популярно коллекционирование редчайших карточек. Производитель выпустил 𝑚 различных видов карточек, пронумерованных от 1 до 𝑚. Эти карточки настолько редкие, что их продает только один человек. Известно, что у него осталось всего 𝑚 карточек, по одной каждого вида.

Вам известно, что у 𝑖-го из ваших друзей есть все карточки с номерами от 1 до 𝑎𝑖 включительно. Вы хотите сделать подарок всем своим друзьям, подарив 𝑖-му из них карточку 𝑏𝑖, которой у него еще нет, то есть такую, что 𝑏𝑖>𝑎𝑖.

## Входные данные

Первая строка содержит два целых числа 𝑛 и 𝑚 (1≤𝑛,𝑚≤10^5) — количество друзей и количество карточек.

Вторая строка содержит 𝑛 целых чисел 𝑎𝑖 (1≤𝑎𝑖≤𝑚). Решения, работающие правильно при 𝑛,𝑚≤100, получат 10 баллов.

## Выходные данные

Выведите массив 𝑏𝑖 или −1, если ответа не существует. Если ответов несколько, выведите любой.

Пояснение к первому примеру:

У нас есть 5 друзей и 7 различных карточек.
- Мы дарим 1 другу — 5 карточку, так как ее у него еще нет.

Пока неподаренными остались карточки [1, 2, 3, 4, 6, 7].
- Мы дарим 2 другу — 4 карточку, так как ее у него еще нет.

Пока неподаренными остались карточки [1, 2, 3, 6, 7].
- Мы дарим 3 другу — 3 карточку, так как ее у него еще нет.

Пока неподаренными остались карточки [1, 2, 6, 7].
- Мы дарим 4 другу — 7 карточку, так как ее у него еще нет.

Пока неподаренными остались карточки [1, 2, 6].
- Мы дарим 5 другу — 6 карточку, так как ее у него еще нет.

Пока неподаренными остались карточки [1, 2]. Также возможны и другие варианты раздачи карточек. Во втором тесте никак нельзя выдать карточки правильно, поэтому и ответ — -1.

## Пример теста 1

### Входные данные

```
5 7
3 3 2 6 5

```

### Выходные данные

```
4 5 3 7 6

```

## Пример теста 2

### Входные данные

```
4 4
2 1 2 2

```

### Выходные данные

```
-1

```

## Пример теста 3

### Входные данные

```
5 6
3 1 2 3 5

```

### Выходные данные

```
4 2 3 5 6

```