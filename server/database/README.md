Dockerfile dosyasında belirtilen ve docker-compose.yml dosyasında yinelenen 
kullanıcı adı ve parola bilgileri geçerli olacak şekilde 
bir veritabanı ve görselleştiricisi (phpmyadmin) ayaklandırmak için 

<pre>docker-compose up</pre>

komutunun docker-compose.yml dosyasının bulunduğu dizinde 
kullanılması gerekmektedir.

Docker compose dosyasında da belirtildii gibi phpmyadmin erişimi sunucu adresi localhost ise:

<pre>localhost:8081</pre> şeklinde gerçekleşmektedir.

Not: phpmyadmin servisi 8081 portunu kullanıyor şu anda, go.api ise 8080
