package models

type basic struct{
	device_per_user int `json:"device___per___user,omitempty"`
	max_admin       int `json:"max___admin,omitempty"`
	max_admin_sz    int `json:"max___admin___sz,omitempty"`
	max_users       int `json:"max___users,omitempty"`
	sz_users        int `json:"sz___users,omitempty"`
}

func NewBasic() basic{
	basicLicense := basic{}
	basicLicense.device_per_user = 1
	basicLicense.max_admin = 1
	basicLicense.max_admin_sz = 1
	basicLicense.max_users = 5
	basicLicense.sz_users = 0
	return basicLicense
}

type business struct{
	device_per_user int    `json:"device___per___user,omitempty"`
	max_admin       int    `json:"max___admin,omitempty"`
	max_admin_sz    int    `json:"max___admin___sz,omitempty"`
	max_users       int    `json:"max___users,omitempty"`
	sz_users        int    `json:"sz___users,omitempty"`
	shared_sz       int    `json:"shared___sz,omitempty"`
	support         string `json:"support,omitempty"`
}

func NewBusinnes() business{
	businnesLicense := business{}
	businnesLicense.device_per_user = -1
	businnesLicense.max_admin = 3
	businnesLicense.max_admin_sz = 1
	businnesLicense.max_users = 5
	businnesLicense.sz_users = 0
	businnesLicense.shared_sz = 1
	businnesLicense.support = "basic"
	return businnesLicense
}