#!/bin/bash

echo "go run . --reverse example00.txt"
go run . --reverse example00.txt
echo
echo "---------------------------------------------------------"

go run . --output=example00.txt "Hello World"
echo "go run . --reverse=example00.txt"
go run . --reverse=example00.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example01.txt "123"
echo "go run . --reverse=example01.txt"
go run . --reverse=example01.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example02.txt "#=\["
echo "go run . --reverse=example02.txt"
go run . --reverse=example02.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example03.txt "something&234"
echo "go run . --reverse=example03.txt"
go run . --reverse=example03.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example04.txt "abcdefghijklmnopqrstuvwxyz"
echo "go run . --reverse=example04.txt"
go run . --reverse=example04.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example05.txt "\!\" #$%&'()*+,-./"
echo "go run . --reverse=example05.txt"
go run . --reverse=example05.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example06.txt ":;{=}?@"
echo "go run . --reverse=example06.txt"
go run . --reverse=example06.txt 
echo
echo "---------------------------------------------------------"

go run . --output=example07.txt "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
echo "go run . --reverse=example07.txt"
go run . --reverse=example07.txt  
echo
echo "---------------------------------------------------------"
