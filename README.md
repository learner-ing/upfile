# upfile
服务器使用nc接收文件
 `nc -lvvp 8888 >file.tar.gz`
 
客户端上传文件
`upfile -dir /var/www -host 192.168.1.1:8888`

upfile -dir /var/www -host 192.168.1.1:8888
upfile -dir /var/www -host https://192.168.1.1 还未开发
  -dir string
        Compressed folder,[/var/www]
  -host string
        Server IP port,[192.168.1.1:8888]
