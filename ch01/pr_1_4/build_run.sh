#!/bin/bash

echo "--- build"
echo " :"
echo ""

go build -o practice_1_4

if [ $? -ne 0 ]; then
    echo ""
    echo "Oops...😣"
    exit 1
fi

echo "Build succeeded🎉"
echo ""

cat << EOS > sample1.txt
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
🍉
Orange
Orange
EOS

cat << EOS > sample2.txt
apple
🍕
apple
🍕
🍉
EOS

echo "--- run"
./practice_1_4 sample1.txt sample2.txt
