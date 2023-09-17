package models

type Query struct {
	Query struct {
		Bool struct {
			Should []struct {
				MatchPhrase map[string]struct {
					Query string  
					Boost float64 
				} 
			} 
		} 
	} 
	Size int 
}

type CreateQueryCMD struct {
	Query struct {
		Bool struct {
			Should []struct {
				MatchPhrase map[string]struct {
					Query string  `json:"query"`
					Boost float64 `json:"boost"`
				} `json:"match_phrase"`
			} `json:"should"`
		} `json:"bool"`
	} `json:"query"`
	Size int `json:"size"`
}
