<?php 

function loadX ($route)
{
    $handler = fopen($route, 'r');
    $lines = [];
    if ($handler) {
        while(($line = fgets($handler)) !== false) {
            if ($line !== '') {
                $lines[] = trim($line);
            }
        }
        fclose($handler);
    }
    return $lines;
}

function loadCorpus($route)
{
    $handler = fopen($route, 'r');
    $entries = [];

    if ($handler) {
        while(($line = fgets($handler)) !== false) {
            if ($line !== '') {
                $entry = explode(' ', $line);
                $entries[] = [
                    'x' => trim($entry[0]),
                    'y' => intval($entry[1]),
                    'value' => trim($entry[2])
                ];
            }
        }
        fclose($handler);
    }

    shuffle($entries);
    return $entries;
}

// End of loader.php
