package songs

type AddSongToInventory struct {
	Name      string `json:"name"`
	CompanyId int64
}
