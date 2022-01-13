package user

import (
	"disspace/business/user"
	"time"
)

type User struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName  string    `json:"full_name" bson:"full_name"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	Email     string    `json:"email" bson:"email"`
	Gender    string    `json:"gender" bson:"gender"`
	Role      int       `json:"role" bson:"role"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type UserProfile struct {
	ID          string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string   `json:"username" bson:"username"`
	ProfilePict string   `json:"profile_pict" bson:"profile_pict"`
	Bio         string   `json:"bio" bson:"bio"`
	Following   []string `json:"following" bson:"following"`
	Followers   []string `json:"followers" bson:"followers"`
	Threads     []string `json:"threads" bson:"threads"`
	Reputation  int      `json:"reputation" bson:"reputation"`
}

type UserSession struct {
	Token    string `json:"token" bson:"token"`
	Username string `json:"username" bson:"username"`
}

func (record *UserSession) SessionToDomain() user.UserSessionDomain {
	return user.UserSessionDomain{
		Token:    record.Token,
		Username: record.Username,
	}
}

func SessionFromDomain(domain user.UserSessionDomain) UserSession {
	return UserSession{
		Token:    domain.Token,
		Username: domain.Username,
	}
}

func (record *User) UserToDomain() user.UserDomain {
	return user.UserDomain{
		ID:        record.ID,
		FullName:  record.FullName,
		Username:  record.Username,
		Password:  record.Password,
		Email:     record.Email,
		Gender:    record.Gender,
		Role:      record.Role,
		Status:    record.Status,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func (record *UserProfile) UserProfileToDomain() user.UserProfileDomain {
	return user.UserProfileDomain{
		ID:          record.ID,
		Username:    record.Username,
		ProfilePict: record.ProfilePict,
		Bio:         record.Bio,
		Following:   record.Following,
		Followers:   record.Followers,
		Threads:     record.Threads,
		Reputation:  record.Reputation,
	}
}

func UserFromDomain(domain user.UserDomain) User {
	return User{
		ID:        domain.ID,
		FullName:  domain.FullName,
		Username:  domain.Username,
		Password:  domain.Password,
		Email:     domain.Email,
		Gender:    domain.Gender,
		Role:      domain.Role,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func UserProfileFromDomain(domain user.UserProfileDomain) UserProfile {
	return UserProfile{
		ID:          domain.ID,
		Username:    domain.Username,
		ProfilePict: domain.ProfilePict,
		Bio:         domain.Bio,
		Following:   domain.Following,
		Followers:   domain.Followers,
		Threads:     domain.Threads,
		Reputation:  domain.Reputation,
	}
}
