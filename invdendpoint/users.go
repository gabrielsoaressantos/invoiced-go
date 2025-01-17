package invdendpoint

const UsersEndpoint = "/members"

type UserRequests []UserRequest

type UserRequest struct {
	Email           string              `json:"email,omitempty"`
	FirstName       string              `json:"first_name,omitempty"`
	LastName        string              `json:"last_name,omitempty"`
	Role            string              `json:"role,omitempty"`
	RestrictionMode string              `json:"restriction_mode,omitempty"`
	Restrictions    map[string][]string `json:"restrictions,omitempty"`
}

type UserResponse struct {
	CreatedAt            int64               `json:"created_at,omitempty"`	//Timestamp when created
	UpdatedAt            int64               `json:"updated_at,omitempty"`
	EmailUpdateFrequency string              `json:"email_update_frequency,omitempty"`
	Id                   int64               `json:"id,omitempty"`
	LastSignedIn		 int64				 `json:"last_accessed,omitempty"`
	RestrictionMode      string              `json:"restriction_mode,omitempty"`
	Restrictions         map[string][]string `json:"restrictions,omitempty"`
	Role                 string              `json:"role,omitempty"`
	User                 *User               `json:"user,omitempty"`
}

type UserEmailUpdateRequest struct {
	Id                   int64               `json:"id"`
	EmailUpdateFrequency string              `json:"email_update_frequency"`
}

type UserInviteRequest struct {
	Id                   int64               `json:"id"`
}

type User struct {
	Email            string `json:"email,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	Id               int64  `json:"id,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Registered       bool   `json:"registered,omitempty"`
	TwoFactorEnabled bool   `json:"two_factor_enabled,omitempty"`
}
