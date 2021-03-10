package models

type Counters struct {
	ID              int `bson:"_id" json:"id"`
	CountEpisode    int `bson:"count_episode" json:"count_episode"`
	CountLocation   int `bson:"count_location" json:"count_location"`
	CountCharacters int `bson:"count_character" json:"count_character"`
}
