package req

import (
	"fmt"
	"testing"
)

//sudo docker run -p 8081:80 kennethreitz/httpbin
func TestGet(t *testing.T) {
	header := Header{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
	}
	r, err := Get("http://127.0.0.1:8081/get", header)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.ToBytes())
}

func TestPost(t *testing.T) {
	header := Header{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
	}

	param := Param{
		"username": "test",
		"password": "test",
	}

	r, err := Post("http://127.0.0.1:8081/post", header, param)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.ToString())
}

func TestRetry(t *testing.T) {
	header := Header{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
	}

	r, err := Get("http://127.0.0.1:8081/basic-auth/admin/admin", header, DisableAllowRedirect, BasicAuth{Username: "admin", Password: "admin"})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.ToString())
	fmt.Println(r.resp.Location())
}
