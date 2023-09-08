package zincsearch

import (
	"testing"
)

func NEWCLIENT(url, user, password string) *NewZincSearchClient {
	return &NewZincSearchClient{
		Url:      url,
		User:     user,
		Password: password,
	}
}

func Test_checkClient(t *testing.T) {
	r := NEWCLIENT("http://zincsearch:4080/", "user", "password")
	err := r.checkClient()

	if err != nil {
		t.Errorf("la validación falló con el error: %v", err)
		t.Fail()
	}
}
