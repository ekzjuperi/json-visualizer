# json-visualizer
Утилита табличного вывода массива данных, хронящихся в файле формата JSON.

В качестве аргумента утилите передаётся путь к файлу, который в ключе items содержит массив объектов. Каждый объект имеет следующие ключи:

name — фамилия и имя студента (строчный тип)
group — номер группы (строчный/целочисленный тип)
avg — средний балл (строчный/целочисленный/вещественный тип)
debt — список задолженностей (строчный/перечислительный тип)

Утилиту можно запустить с аргументом -dir указав путь до JSON файла.(по умолчанию используется корневой каталог).
