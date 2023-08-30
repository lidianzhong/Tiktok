package util


type Author struct {
	Id              string `json:"id"`
	Name            string `json:"name,omitempty"`
	FollowCount     string `json:"follow_count"`
	FollowCounter   string `json:"follow_counter"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       string `json:"work_count"`
	FavoriteCount   string `json:"favorite_count"`
}

// TableName 表示配置操作数据库的表名称
func (Author) TableName() string {
	return "author"
}
