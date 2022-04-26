package tokens

type Token struct {
	Token string `json:"token" gorm:"primaryKey"`
	Id    int64  `json:"id"`
}
