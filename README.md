# tcping

Simple CLI powered by [github.com/tevino/tcp-shaker](https://github.com/tevino/tcp-shaker) to perform TCP handshake without ACK, useful for health checking.

```sh
> tcping github.com:123
KO Connect to github.com:123 timed out

> tcping github.com:22
OK Connect to github.com:22 succeede
```