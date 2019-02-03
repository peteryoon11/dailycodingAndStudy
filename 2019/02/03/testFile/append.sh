#!/bin/bash
echo "test"

for i in {1..60}
do 
    sleep 1s
    NOW=$(date '+%G-%m-%e %H:%M:%S')
    echo "[$NOW]"
done 