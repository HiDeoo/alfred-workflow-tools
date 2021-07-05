package main

type BSShows struct {
	Shows []BSShow `json:"shows"`
	BSErrors
}

type BSShow struct {
	ID             int    `json:"id"`
	ThetvdbID      int    `json:"thetvdb_id"`
	ImdbID         string `json:"imdb_id"`
	Title          string `json:"title"`
	OriginalTitle  string `json:"original_title"`
	Description    string `json:"description"`
	Seasons        string `json:"seasons"`
	SeasonsDetails []struct {
		Number   int `json:"number"`
		Episodes int `json:"episodes"`
	} `json:"seasons_details"`
	Episodes   string `json:"episodes"`
	Followers  string `json:"followers"`
	Comments   string `json:"comments"`
	Similars   string `json:"similars"`
	Characters string `json:"characters"`
	Creation   string `json:"creation"`
	Showrunner struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	} `json:"showrunner"`
	Genres struct {
		Adventure      string `json:"Adventure"`
		Comedy         string `json:"Comedy"`
		ScienceFiction string `json:"Science Fiction"`
	} `json:"genres"`
	Length   string `json:"length"`
	Network  string `json:"network"`
	Country  string `json:"country"`
	Rating   string `json:"rating"`
	Status   string `json:"status"`
	Language string `json:"language"`
	Notes    struct {
		Total int     `json:"total"`
		Mean  float64 `json:"mean"`
		User  int     `json:"user"`
	} `json:"notes"`
	InAccount bool `json:"in_account"`
	Images    struct {
		Show   string `json:"show"`
		Banner string `json:"banner"`
		Box    string `json:"box"`
		Poster string `json:"poster"`
	} `json:"images"`
	Aliases struct {
		Num6094  string `json:"6094"`
		Num8073  string `json:"8073"`
		Num18286 string `json:"18286"`
		Num22724 string `json:"22724"`
		Num22729 string `json:"22729"`
		Num36744 string `json:"36744"`
		Num36745 string `json:"36745"`
		Num36746 string `json:"36746"`
		Num62382 string `json:"62382"`
	} `json:"aliases"`
	SocialLinks []struct {
		Type       string `json:"type"`
		ExternalID string `json:"external_id"`
	} `json:"social_links"`
	User struct {
		Archived  bool    `json:"archived"`
		Favorited bool    `json:"favorited"`
		Remaining int     `json:"remaining"`
		Status    float64 `json:"status"`
		Last      string  `json:"last"`
		Tags      string  `json:"tags"`
		Next      struct {
			ID    int    `json:"id"`
			Code  string `json:"code"`
			Date  string `json:"date"`
			Title string `json:"title"`
			Image string `json:"image"`
		} `json:"next"`
		FriendsWatching []interface{} `json:"friends_watching"`
	} `json:"user"`
	NextTrailer string `json:"next_trailer"`
	ResourceURL string `json:"resource_url"`
	Platforms   struct {
		Svods []struct {
			ID        string      `json:"id"`
			Name      string      `json:"name"`
			Tag       interface{} `json:"tag"`
			LinkURL   string      `json:"link_url"`
			Available struct {
				Last  int `json:"last"`
				First int `json:"first"`
			} `json:"available"`
			Logo string `json:"logo"`
		} `json:"svods"`
		Svod struct {
			ID        string      `json:"id"`
			Name      string      `json:"name"`
			Tag       interface{} `json:"tag"`
			LinkURL   string      `json:"link_url"`
			Available struct {
				Last  int `json:"last"`
				First int `json:"first"`
			} `json:"available"`
			Logo string `json:"logo"`
		} `json:"svod"`
	} `json:"platforms"`
}

type BSErrors struct {
	Errors []BSError `json:"errors"`
}

type BSError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
