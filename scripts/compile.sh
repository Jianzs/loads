#!/bin/bash

COMPILE_DIR=${1-''}
BIN_DIR=${2-'bin'}

for group_dir in ./* 
do
    if [ $group_dir = "scripts" ] || [ $group_dir = "bin" ] || ! [ -d $group_dir ]
    then
        continue
    fi
    if ! [ -z $COMPILE_DIR ] && ! [ $group_dir = "./$COMPILE_DIR" ]
    then
        continue
    fi
    echo "compiling $group_dir" 

    for load_dir in $group_dir/*
    do
        if ! [ -d $load_dir ]
        then
            continue
        fi

        if [ -f $load_dir/compile.sh ]
        then
            echo "  compiling $load_dir" 
            cd $load_dir
            bash compile.sh "../../$BIN_DIR"
        fi
    done
done