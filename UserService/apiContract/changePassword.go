package apicontract

type ChangePassword struct {
	Id          uint
	OldPassword string
	NewPassword string
}
