# tbolimpiada_semifinals_DevOps_2024
![banner.png](banner.png)

[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MisterZurg/tbolimpiada-semifinals-DevOps-2024/)
[![Go Report Card](https://goreportcard.com/badge/github.com/MisterZurg/tbolimpiada_semifinals_DevOps_2024)](https://goreportcard.com/report/github.com/MisterZurg/tbolimpiada_semifinals_DevOps_2024/)
[![codecov](https://codecov.io/gh/MisterZurg/tbolimpiada_semifinals_DevOps_2024/branch/dungeon-master/graph/badge.svg)](https://codecov.io/gh/MisterZurg/tbolimpiada_semifinals_DevOps_2024)

> [!NOTE]
> Задание - игра «Заливка»
> В игре есть два игрока (Игрок1, Игрок2).
> Размерность поля 47 x 21 гексагональных клеток в виде прямоугольника.
> На поле могут находится стены, в заранее известной конфигурации.
> Очевидно, что стены не находятся в позициях старта игроков, и симметричны для обоих игроков относительно центра.
> От стартовых позиций игроков можно проложить маршрут друг к другу.
> Каждая клетка поля имеет один из 10 цветов, заданных случайным образом:
> 1. Белый
> 2. Ярко-красный
> 3. Зеленый
> 4. Ярко-зеленый
> 5. Синий
> 6. Светло-синий
> 7. Желтый
> 8. Розовый
> 9. Оранжевый
> 10. Черный (стены, цвет не доступен для выбора)
> 
> Игроки начинают игру в противоположных углах.
> 
> Первый игрок начинает игру в левом нижнем углу.
> 
> Второй игрок - в правом верхнем.
> Игроки выбирают цвет по очереди, начиная с Игрока1.
> 
> Для выбора цвета разрешены те цвета, с которыми есть сопряжения (соприкосновения).
> Нельзя выбрать цвет соперника.
> При выборе нового цвета к пространству игрока добавляются клетки, соприкасающиеся с полем игрока, в которых указан данный цвет (т.е. при заливке цветом поле постепенно растет не менее, чем на 1 клетку).
> Игра останавливается в тот момент, когда игрок не может выбрать цвет и совершить ход.
> Побеждает тот игрок, у которого в момент окончания игры больше занятых клеток. 

## Установка
```shell
go get github.com/MisterZurg/tbolimpiada-semifinals-DevOps-2024
```