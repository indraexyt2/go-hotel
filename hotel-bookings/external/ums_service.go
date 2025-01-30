package external

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type External struct {
}

func NewExternal() *External {
	return &External{}
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKe;autoIncrement"`
	Username  string    `json:"username" gorm:"type:varchar(50);uniqueIndex" validate:"required" form:"username"`
	Email     string    `json:"email" gorm:"type:varchar(50);uniqueIndex" validate:"required,email" form:"email"`
	Role      string    `json:"role" gorm:"type:user_role;default:guest"`
	FullName  string    `json:"full_name" type:"varchar(50)" form:"full_name"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdateAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

type UMSResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

func (ex *External) ValidateUser(ctx context.Context, token string) (*User, error) {
	url := os.Getenv("UMS_URL_USER")
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	var umsResponse UMSResponse
	err = json.NewDecoder(response.Body).Decode(&umsResponse)
	if err != nil {
		return nil, err
	}

	return &umsResponse.Data, nil
}
