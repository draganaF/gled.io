package apicontract

type ChangePassword struct {
	UserId      uint
	OldPassword string
	NewPassword string
}
