graw
====
[![Build Status](https://travis-ci.org/abo/graw.svg)](https://travis-ci.org/abo/graw)
[![GoDoc](https://godoc.org/github.com/abo/graw/patn?status.svg)](https://godoc.org/github.com/abo/graw/patn)
[![Go Report Card](https://goreportcard.com/badge/github.com/abo/graw)](https://goreportcard.com/report/github.com/abo/graw)
[![Coverage Status](https://coveralls.io/repos/github/abo/graw/badge.svg?branch=master)](https://coveralls.io/github/abo/graw?branch=master)

A pattern generator and searcher

* Auto generate pattern according sample
* format output by template


Installation
------------

Install graw using the "go get" command:

    go get github.com/abo/graw

make sure the $GOPATH/bin is added into $PATH


Example
-------

    ~ abo$ cat nginx_access.log
    74.125.19.106 - - [15/May/2015:18:27:55] "GET /product.screen?productId=DC-SG-G02&JSESSIONID=SD10SL4FF4ADFF4976 HTTP 1.1" 200 3279 "http://www.buttercupgames.com/oldlink?itemId=EST-26" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.46 Safari/536.5" 559
    203.223.0.20 - - [15/May/2015:18:30:03] "POST /category.screen?categoryId=NULL&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 406 1944 "http://www.buttercupgames.com/product.screen?productId=SF-BVS-G01" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 417
    203.223.0.20 - - [15/May/2015:18:30:05] "GET /category.screen?categoryId=ACCESSORIES&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 200 2006 "http://www.buttercupgames.com/oldlink?itemId=EST-17" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 509
    203.223.0.20 - - [15/May/2015:18:30:06] "GET /product.screen?productId=BS-AG-G09&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 200 1051 "http://www.buttercupgames.com/cart.do?action=addtocart&itemId=EST-6&productId=BS-AG-G09" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 179
    125.89.78.6 - - [15/May/2015:18:34:10] "GET /oldlink?itemId=EST-7&JSESSIONID=SD8SL10FF1ADFF5003 HTTP 1.1" 200 1140 "http://www.yahoo.com" "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B206 Safari/7534.48.3" 471
    125.89.78.6 - - [15/May/2015:18:34:12] "GET /oldlink?itemId=EST-17&JSESSIONID=SD8SL10FF1ADFF5003 HTTP 1.1" 200 620 "http://www.buttercupgames.com/oldlink?itemId=EST-17" "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B206 Safari/7534.48.3" 590
    125.89.78.6 - - [15/May/2015:18:34:13] "GET /product.screen?productId=WC-SH-T02&JSESSIONID=SD8SL10FF1ADFF5003 HTTP 1.1" 200 2147 "http://www.buttercupgames.com/cart.do?action=remove&itemId=EST-27&productId=WC-SH-T02" "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B206 Safari/7534.48.3" 291
    125.89.78.6 - - [15/May/2015:18:34:15] "GET /oldlink?itemId=EST-7&JSESSIONID=SD8SL10FF1ADFF5003 HTTP 1.1" 200 1578 "http://www.buttercupgames.com/cart.do?action=view&itemId=EST-7&productId=WC-SH-G04" "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B206 Safari/7534.48.3" 933
    ~ abo$ cat nginx_access.log | graw 18:27:55
    18:27:55
    18:30:03
    18:30:05
    18:30:06
    18:34:10
    18:34:12
    18:34:13
    18:34:15
    ~ abo$ cat nginx_access.log | graw -format="{{index . 1}} -- {{index . 0}}" /product.screen 74.125.19.106
    74.125.19.106 -- /product.screen
    203.223.0.20 -- /category.screen
    203.223.0.20 -- /category.screen
    203.223.0.20 -- /product.screen
    125.89.78.6 -- /oldlink
    125.89.78.6 -- /oldlink
    125.89.78.6 -- /product.screen
    125.89.78.6 -- /oldlink


License
-------

graw is available under the [The MIT License (MIT)](https://opensource.org/licenses/MIT).    