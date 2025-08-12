# go-digestive

### HTTP Digest Dictionary Cracking Tool

**go-digestive** is a tool for performing dictionary attacks against captured HTTP Digest authentication credentials. It's designed to find the correct password by testing a provided wordlist against a captured hash.

These credentials can be obtained from various sources, such as a **PCAP**, **Burp Suite**, or **ZAP** proxy.

**Example Captured Credentials**:

```

Authorization: Digest username="conrad", realm="Security542", nonce="q5mFt62KBQA=30b96361b3061fc88ad88a19170b873073ccb930", uri="/digest/", algorithm=MD5, response="36bf5df32af14d751ee901fcd2d72479", qop=auth, nc=00000001, cnonce="763eb7656a737513"

```

---

### Usage

Use the command-line flags to provide the captured credentials and the path to your wordlist. The following example demonstrates how to use the tool with credentials from the example above and a common wordlist.


```

$ ./go-digestive --username conrad --wordlist /opt/john/run/password.lst --method GET --uri /digest/ --nc 00000001 --qop auth --realm Security542 --cnonce 763eb7656a737513  --nonce q5mFt62KBQA=30b96361b3061fc88ad88a19170b873073ccb930 --response 36bf5df32af14d751ee901fcd2d72479

```

**Expected Output**:

```

Username = conrad
Password = stargate

```

---

### Credits

This project is a high-performance Go port based on the original Python 2 implementation by **Eric Conrad (@eric_conrad)**.

* **Original Python 2 Project**: [https://github.com/eric-conrad/digestive](https://github.com/eric-conrad/digestive)

A Python 3 compatible fork is also available:

* **Python 3 Fork**: [https://github.com/virer/digestive](https://github.com/virer/digestive)

