package zincsearch

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

type ZincSearchClient struct {
	Url      string
	User     string
	Password string
}
type NewZincSearchClient struct {
	Url      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (n *NewZincSearchClient) checkClient() error {
	req, err := http.Get(n.Url)
	if err != nil {
		fmt.Printf("error al conectarse a %s: %v", n.Url, err)
		return err // Retorna el error
	}
	defer req.Body.Close()
	fmt.Println(req.StatusCode)
	return nil

}
