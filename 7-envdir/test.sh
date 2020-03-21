#!/usr/bin/env bash
rm -r ./env
mkdir -p env/1/

#Тест на удаление переменных
touch env/SHELL

#Тест на нечитаемые директории 
echo "error" > env/1/test1
echo "error" > env/1=2

#тест просто на чтение(позитивный)
echo "ok" > env/test

#тест на обрезание пустых строк
echo -ne "ok\n\t" > env/test_trailing_spaces

#тест на замену null byte
echo -ne "null\0byte" > env/test_nullbyte

go run . env env > result.txt
