package hummingbird

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
