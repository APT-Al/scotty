main.go dosyası

<pre>go build</pre>

komutuyla derlenir. Çıkan executable dosyanın çalıştırılabilmesi için ise 
belli paketlerin ``go get`` ile kurulması gerekmektedir.

Bu paketler

<pre>
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
</pre>
Bunlar arasında muhtemelen log ve net/http paketleri kuruludur ama 
yine de kontrol edilmesinde fayda olacakır.
