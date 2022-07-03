package authentication

type Company struct {
	Id                  int64   `db:"id"`
	Name                *string `db:"name"`
	AllowedAdminDevices int32   `db:"allowedadmindevices"`
	AllowedUserDevices  int32   `db:"alloweduserdevices"`
	AdminHashedPass     *string `db:"adminhashedpass"`
	UserHashedPass      *string `db:"userhashedpass"`
}
