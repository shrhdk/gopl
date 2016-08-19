#!/bin/sh

go get gopl.shiro.be/ch08/ex05/surface

for i in 1 2 3 4 5 6 7 8
do
    echo $i
    export GOMAXPROCS=$i
    time $GOPATH/bin/surface > /dev/null
    echo
done
