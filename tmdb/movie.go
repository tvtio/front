package tmdb

// Movie ...
type Movie struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		BackdropPath string `json:"backdrop_path"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
	} `json:"belongs_to_collection"`
	Budget  int `json:"budget"`
	Credits struct {
		Cast []struct {
			CastID      int    `json:"cast_id"`
			Character   string `json:"character"`
			CreditID    string `json:"credit_id"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Order       int    `json:"order"`
			ProfilePath string `json:"profile_path"`
		} `json:"cast"`
		Crew []struct {
			CreditID    string `json:"credit_id"`
			Department  string `json:"department"`
			ID          int    `json:"id"`
			Job         string `json:"job"`
			Name        string `json:"name"`
			ProfilePath string `json:"profile_path"`
		} `json:"crew"`
	} `json:"credits"`
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage string `json:"homepage"`
	ID       int    `json:"id"`
	Images   struct {
		Backdrops []struct {
			AspectRatio float64     `json:"aspect_ratio"`
			FilePath    string      `json:"file_path"`
			Height      int         `json:"height"`
			Iso639_1    interface{} `json:"iso_639_1"`
			VoteAverage float64     `json:"vote_average"`
			VoteCount   int         `json:"vote_count"`
			Width       int         `json:"width"`
		} `json:"backdrops"`
		Posters []struct {
			AspectRatio float64 `json:"aspect_ratio"`
			FilePath    string  `json:"file_path"`
			Height      int     `json:"height"`
			Iso639_1    string  `json:"iso_639_1"`
			VoteAverage float64 `json:"vote_average"`
			VoteCount   int     `json:"vote_count"`
			Width       int     `json:"width"`
		} `json:"posters"`
	} `json:"images"`
	ImdbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int    `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}
