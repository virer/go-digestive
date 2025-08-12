# go-digestive 

Dictionary cracking tool for HTTP Digest challenge/response hashes

Launches dictionary attack vs. captured HTTP Digest credentials (taken from a PCAP, Burp or ZAP proxy, etc.)

```Authorization: Digest username="conrad", realm="Security542", nonce="q5mFt62KBQA=30b96361b3061fc88ad88a19170b873073ccb930", uri="/digest/", algorithm=MD5, response="36bf5df32af14d751ee901fcd2d72479", qop=auth, nc=00000001, cnonce="763eb7656a737513"```

## Resulting command line (using John the Ripper's wordlist)
 
```go-digestive --username conrad --wordlist /opt/john/run/password.lst --method GET --uri /digest/ --nc 00000001 --qop auth --realm Security542 --cnonce 763eb7656a737513  --nonce q5mFt62KBQA=30b96361b3061fc88ad88a19170b873073ccb930 --response 36bf5df32af14d751ee901fcd2d72479```

## About

This project is based on the work in Python2 version from Eric Conrad (@eric_conrad)
This project is available here: https://github.com/eric-conrad/digestive

Also, I've published a Python3 compatible fork here: https://github.com/virer/digestive
