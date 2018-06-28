#!/bin/bash


ab -c 1024 -n 100000 -t 60 http://192.168.99.101/accounts/10020 &
ab -c 1024 -n 100000 -t 60 http://192.168.99.101/accounts/10040 &
ab -c 1024 -n 100000 -t 60 http://192.168.99.101/accounts/10060 &
ab -c 1024 -n 100000 -t 60 http://192.168.99.101/accounts/10080
