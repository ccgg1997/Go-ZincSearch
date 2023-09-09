package zincsearch

import (
	"testing"
	"os"
)

func NEWCLIENT(url, user, password string) *NewZincSearchClient {
	return &NewZincSearchClient{
		Url:      url,
		User:     user,
		Password: password,
	}
}

func Test_checkClient(t *testing.T) {
	r := NEWCLIENT(os.Getenv("ZINC_API_URL"), "user", "password")
	err := r.checkClient()

	if err != nil {
		t.Errorf("la validación falló con el error: %v", err)
		t.Fail()
	}
}
