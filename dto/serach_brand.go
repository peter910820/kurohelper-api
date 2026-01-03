package dto

// SearchBrandResponse 品牌搜尋回應結構
type SearchBrandResponse struct {
	ID          string   `json:"id"`          // 品牌 ID
	Name        string   `json:"name"`        // 品牌名稱
	Aliases     []string `json:"aliases"`     // 品牌別名列表
	Description string   `json:"description"` // 品牌描述
	Link        Link     `json:"link"`        // 品牌相關連結
	VN          []VN     `json:"vn"`          // 視覺小說資訊
}

// Link 品牌相關連結結構
type Link struct {
	OfficialWebsite string `json:"official_website"` // 官方網站
	Wikipedia       string `json:"wikipedia"`        // 維基百科連結
	Xitter          string `json:"xitter"`           // X (Twitter) 連結
	Steam           string `json:"steam"`            // Steam 連結
}

// VN 視覺小說資訊結構
type VN struct {
	Title         string  `json:"title"`          // 標題
	AltTitle      string  `json:"alttitle"`       // 別名標題
	Average       float64 `json:"average"`        // 平均評分
	Rating        float64 `json:"rating"`         // 評分
	VoteCount     int     `json:"votecount"`      // 投票數
	LengthMinutes int     `json:"length_minutes"` // 遊玩時長（分鐘）
	LengthVotes   int     `json:"length_votes"`   // 時長投票數
}
