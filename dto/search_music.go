package dto

import "time"

type SearchMusicResponse struct {
	PlayTime          time.Time `json:"play_time"`           // 音樂時長
	ReleaseDate       time.Time `json:"release_date"`        // 發行日期
	AverageScore      float64   `json:"average_score"`       // 平均分數
	AverageScoreCount int       `json:"average_score_count"` // 平均分數樣本數
	Singers           []string  `json:"singers"`             // 歌手
	Lyricists         []string  `json:"lyricists"`           // 作詞家
	Composers         []string  `json:"composers"`           // 作曲家
	Arrangers         []string  `json:"arrangers"`           // 編曲家
	GamesIncluded     []string  `json:"games_included"`      // 收錄遊戲
	AlbumsIncluded    []string  `json:"albums_included"`     // 收錄專輯
}
