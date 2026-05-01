farg="2048"
$1 -f $farg
size=$(stat -c%s yeah_0-${farg}.txt)

if (( $size == $farg )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

farg="9M"
$1 -f $farg
size=$(stat -c%s yeah_0-${farg}.txt)

if (( $size == 9437184 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

echo -e "\n[*] TEST LIST"
farg="11M,308K,412"
$1 -f $farg

size=$(stat -c%s yeah_0-11M.txt)

if (( $size == 11534336 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

size=$(stat -c%s yeah_1-308K.txt)

if (( $size == 315392 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

size=$(stat -c%s yeah_2-412.txt)

if (( $size == 412 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi


echo -e "\n[*] SYNCHRON TEST LIST"
farg="3M,288K,120"
$1 -f $farg -s

size=$(stat -c%s yeah_0-3M.txt)

if (( $size == 3145728 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

size=$(stat -c%s yeah_1-288K.txt)

if (( $size == 294912 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

size=$(stat -c%s yeah_2-120.txt)

if (( $size == 120 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

echo -e "\n[*] LINEAR TEST LIST"
farg="1M,309K,412"
$1 -f $farg -l

size=$(stat -c%s yeah_0-1M.txt)

if (( $size == 1048576 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

size=$(stat -c%s yeah_1-309K.txt)

if (( $size == 316416 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

size=$(stat -c%s yeah_2-412.txt)

if (( $size == 412 )); then
    echo "${farg} PASS"
else 
    echo "${farg} FAIL"
    exit 1
fi

echo -e "\n[*] ALL TESTS PASS"