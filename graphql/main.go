package main

type GraphqlRequest struct {
	Query string `json:"query"`
}

type GraphqlResponse struct {
	Data struct {
		User struct {
			Name      string `json:"name"`
			Following struct {
				Nodes []struct {
					Login   string `json:"login"`
					HtmlUrl string `json:"html_url"`
				} `json:"nodes"`
			} `json:"following"`
			Followers struct {
				Nodes []struct {
					Login   string `json:"login"`
					HtmlUrl string `json:"html_url"`
				} `json:"nodes"`
			} `json:"followers"`
		} `json:"user"`
	} `json:"data"`
}

func main() {
	
}
