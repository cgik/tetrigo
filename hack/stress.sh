#!/bin/bash

# Relies on hey (https://github.com/rakyll/hey)
hey -n 100000 -c 100 http://localhost:8080/api/v1/game/init

wait
