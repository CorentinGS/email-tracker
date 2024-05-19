package domain

// Permission is the permission level of a user.
type Permission int64

const (
	PermissionNone  Permission = iota // PermissionNone is the default permission level.
	PermissionUser                    // PermissionUser is the permission level of a user.
	PermissionVip                     // PermissionVip is the permission level of a vip.
	PermissionAdmin                   // PermissionAdmin is the permission level of an admin.
)

func (p Permission) String() string {
	return [...]string{"none", "user", "vip", "admin"}[p]
}

func (p Permission) IsValid() bool {
	return p >= PermissionNone && p <= PermissionAdmin
}

type Nonce struct {
	HtmxNonce           string
	PicoCSSNonce        string
	HyperscriptNonce    string
	PreloadNonce        string
	UmamiNonce          string
	CSSScopeInlineNonce string
}
