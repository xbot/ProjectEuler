<?php
$names = explode(',', str_replace('"', '', file_get_contents('names.txt')));
sort($names, SORT_STRING);
$cal_alpha_value = function($name, $i) {
    $cal_alpha_index = function($char){return ord($char) - 64;};
    return array_sum(array_map($cal_alpha_index, str_split($name, 1))) * $i;
};
echo array_sum(array_map($cal_alpha_value, $names, range(1, count($names))));
?>
