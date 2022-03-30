package repository

import (
	"bytes"
	"encoding/json"
	"enigmacamp.com/goacc/delivery/appreq"
	"enigmacamp.com/goacc/delivery/commonresp"
	"enigmacamp.com/goacc/utils"
	"errors"
)

type userCredentialRepoImpl struct {
	intraClient *utils.IntraClient
}

func (u *userCredentialRepoImpl) GetByUserNameAndPassword(user appreq.AuthRequest) error {
	//var isUserExist int
	//err := u.userCredDb.Get(&isUserExist, "select count(id) from user_credentials where user_name=$1 and user_password=$2 and is_blocked=$3", user.GetUserName(), user.GetUserPassword(), false)
	//if err != nil {
	//	return errors.New(err.Error())
	//}
	//if isUserExist == 0 {
	//	return utils.DataNotFoundError()
	//}
	//return nil
	postBody, _ := json.Marshal(user)
	requestBody := bytes.NewBuffer(postBody)
	response, err := u.intraClient.IntraPost("http://localhost:3334/auth", requestBody)
	if err != nil {
		return errors.New(err.Error())
	}
	resp := commonresp.ApiResponse{}
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return err
	}
	//if response.StatusCode == http.StatusOK {}
	if resp.ResponseCode != "00" {
		return utils.UnauthorizedError()
	}
	return nil

}

func NewUserCredentialRepo(intraClient *utils.IntraClient) UserCredentialRepo {
	return &userCredentialRepoImpl{
		intraClient: intraClient,
	}
}
