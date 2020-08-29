package response

type SysCaptchaResponse struct {
	CaptchaId string `json:"captchaid"`
	PicPath   string `json:"picpath"`
}
