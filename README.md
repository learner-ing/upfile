# upfile
服务器使用nc接收文件
 `nc -lvvp 8888 >file.tar.gz`
 
客户端上传文件
`upfile -dir /var/www -host 192.168.1.1:8888`

```
upfile -dir /var/www -host 192.168.1.1:8888
upfile -dir /var/www -host 192.168.1.1:8888 -err flase
upfile -dir /var/www -host 192.168.1.1:8888 -skip "/var/www/log,/var/www/upload,.pdf,.png"
  -dir string
        Compressed folder,[/var/www]
  -err
        Exit in case of exception,Default true (default true)
  -host string
        Server IP port,[192.168.1.1:8888]
  -skip string
        skip folder
```
下个版本将添加的功能</br>
排除文件后缀
