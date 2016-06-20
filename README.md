graw
====
[![Build Status](https://travis-ci.org/abo/graw.svg)](https://travis-ci.org/abo/graw)
[![GoDoc](https://godoc.org/github.com/abo/graw/patn?status.svg)](https://godoc.org/github.com/abo/graw/patn)
[![Go Report Card](https://goreportcard.com/badge/github.com/abo/graw)](https://goreportcard.com/report/github.com/abo/graw)
[![Coverage Status](https://coveralls.io/repos/github/abo/graw/badge.svg?branch=master)](https://coveralls.io/github/abo/graw?branch=master)

A pattern generator and searcher





    cat access.log
    # 74.125.19.106 - - [15/May/2015:18:27:55] "GET /product.screen?productId=DC-SG-G02&JSESSIONID=SD10SL4FF4ADFF4976 HTTP 1.1" 200 3279 "http://www.buttercupgames.com/oldlink?itemId=EST-26" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.46 Safari/536.5" 559
    # 203.223.0.20 - - [15/May/2015:18:30:02] "GET /product.screen?productId=PZ-SG-G05&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 200 992 "http://www.google.com" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 718
    # 203.223.0.20 - - [15/May/2015:18:30:03] "POST /category.screen?categoryId=NULL&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 406 1944 "http://www.buttercupgames.com/product.screen?productId=SF-BVS-G01" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 417
    # 203.223.0.20 - - [15/May/2015:18:30:04] "GET /cart.do?action=view&itemId=EST-7&productId=SC-MG-G10&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 200 3970 "http://www.buttercupgames.com/cart.do?action=view&itemId=EST-7&productId=SC-MG-G10" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 854
    # 203.223.0.20 - - [15/May/2015:18:30:05] "GET /category.screen?categoryId=ACCESSORIES&JSESSIONID=SD4SL3FF4ADFF4991 HTTP 1.1" 200 2006 "http://www.buttercupgames.com/oldlink?itemId=EST-17" "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; MS-RTC LM 8; InfoPath.2)" 509

    cat access.log | go run cmd/graw/main.go -t "74.125.19.106 18:27:55 /product.screen" -f "{{index . 1}} -- {{index . 0}} {{index . 2}}"