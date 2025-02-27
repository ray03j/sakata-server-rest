package model

type Game struct {
    ID         string `json:"gameId" gorm:"primaryKey"`
    HumanScore int    `json:"humanScore"`
    AIScore    int    `json:"aiScore"`
}
