# sudoku
sudoku solver

Маленькая программа - нахождение решения головоломки судоку

## Install | Установка
```
git clone https://github.com/igiant/sudoku.git
cd sudoku
```
## Build | Компиляция
```
go build main.go
```
## Run | Запуск
```
./sudoku -i file_name -o file_name
```
- **-i**: name of csv-file with sudoku data (default:matrix.csv) | входящий файл с данными (по умолчанию: matrix.csv)
- **-o**: name of file for result (default:result.csv) | файл с результатом решения (по умолчанию: result.csv)
## Format of file with data
csv-file with space-separate and with dot for empty place | csv-файл с пробелом в качестве разделителя и точкой для незаполненой клетки

