package hummingbird

type Anime struct {
	Id               uint64
	Slug             string
	Status           string
	Url              string
	Title            string
	Alternate_title  string
	Episode_count    uint64
	Episode_length   uint64
	Cover_image      string
	Synopsis         string
	Show_type        string
	Started_airing   string
	Finished_airing  string
	Community_rating float64
	Age_rating       string
	Genres           []Genre `json:",omitempty"`
}

type Genre struct {
	Name string
}

type LibraryEntry struct {
	Id               uint64
	Episodes_watched uint64
	Last_watched     string
	Updated_at       string
	Rewatched_times  uint64
	Notes            string
	Notes_present    bool
	Status           string
	Private          bool
	Rewatching       bool
	Anime            Anime
	Rating           LibraryEntryRating
}

type LibraryEntryRating struct {
	Type  string
	Value string
}

type Favorite struct {
	Id         uint64
	User_id    uint64
	Item_id    uint64
	Item_type  string
	Created_at string
	Updated_at string
	fav_rank   uint64
}

type User struct {
	Name                      string
	Waifu                     string
	Waifu_or_husbando         string
	Waifu_slug                string
	Waifu_char_id             string
	Location                  string
	Website                   string
	Avatar                    string
	Cover_image               string
	About                     string
	Bio                       string
	Karma                     uint64
	Life_spent_on_anime       uint64
	Show_adult_content        bool
	Title_language_preference string
	Last_library_update       string
	Online                    bool
	Following                 bool
	Favorites                 []Favorite
}
