<?php
$digits = array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9);

function factorial($n) {
    $result = 1;
    for ($i = 1;$i <= $n;$i++) {
        $result *= $i;
    }
    return $result;
}

function get_perm($digits, $number) {
    $perm = "";
    $len = count($digits);
    $counter = factorial($len - 1);
    for ($i = 0;$i < $len;$i++) {
        if ($counter >= $number) {
            $segment = array_splice($digits, $i, 1);
            $perm .= $segment[0] . get_perm($digits, $number);
            break;
        }
        $number -= $counter;
    }
    return $perm;
}

$startTime = microtime(true);
$perm = get_perm($digits, 1000000);
$costs = (microtime(true) - $startTime) * 1000;
echo "$perm ${costs}ms\n";
