package web

type DetailCreateRequest struct {
	YoutubeId string `json:"youtube_id"`
}


//func (request *DetailCreateRequest) ToDomainService() servuser.User {
//	u := servuser.User{}
//	u.Name = request.Name
//	u.Password = request.Password
//	u.Email = request.Email
//	u.Birthday = request.Birthday
//	u.Gender = request.Gender
//
//	return u
//}