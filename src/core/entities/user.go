package entities

type User struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId    string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	FirstName    string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MidName      string `protobuf:"bytes,4,opt,name=mid_name,json=midName,proto3" json:"mid_name,omitempty"`
	SurName      string `protobuf:"bytes,5,opt,name=sur_name,json=surName,proto3" json:"sur_name,omitempty"`
	AvatarUrl    string `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	IsActive     bool   `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Alias        string `protobuf:"bytes,8,opt,name=alias,proto3" json:"alias,omitempty"`
	BasicAddress string `protobuf:"bytes,9,opt,name=basic_address,json=basicAddress,proto3" json:"basic_address,omitempty"`
	Email        string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Phone        string `protobuf:"bytes,11,opt,name=phone,proto3" json:"phone,omitempty"`
}
