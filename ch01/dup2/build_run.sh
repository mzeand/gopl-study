#!/bin/bash

echo "--- build"
echo " :"
echo ""

go build -o dup2

if [ $? -ne 0 ]; then
    echo ""
    echo "Oops...😣"
    exit 1
fi

echo "Build succeeded🎉"
echo ""

cat << EOS > sample.txt
apple
apple
リンゴ
banana
banana
Orange
banana
リンゴ
Orange
リンゴ
Orange
Orange
Orange
Orange
Orange
EOS

echo "--- run"
./dup2 sample.txt
