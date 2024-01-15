package logic

import (
	"errors"
	"time"

	"github.com/forgoer/openssl"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/md"
	"party-affairs/pkg/util"
)

type Auth struct {
	conf *config.Config
}

func NewAuth(conf *config.Config) *Auth {
	return &Auth{
		conf: conf,
	}
}

type AuthInfo struct {
	AuthId  string
	UnionId string
	Time    int64
}

const (
	PlatformYiBan = "yiban"
	BindEmailName = "bind"
)

func (a *Auth) Login(ctx kratosx.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	var err error
	var authInfo *AuthInfo

	switch in.Platform {
	case "yiban":
		authInfo, err = a.yiBanAuthInfo(in.Code)
	}

	if err != nil {
		return nil, v1.GetAuthInfoErrorFormat(err.Error())
	}

	// 超过授权时间2s
	if time.Now().Unix()-authInfo.Time > 2 {
		return nil, v1.GetAuthInfoErrorFormat("授权已失效")
	}

	// 获取授权信息
	auth := model.Auth{}
	if err := auth.OneByAuthId(ctx, authInfo.AuthId); err != nil {
		return nil, v1.UnBindError()
	}

	// 生成用户token
	token, err := ctx.JWT().NewToken(md.New(auth.UserID))
	if err != nil {
		return nil, v1.LoginError()
	}

	return &v1.LoginReply{Token: token}, nil
}

func (a *Auth) Bind(ctx kratosx.Context, in *v1.BindRequest) (*v1.BindReply, error) {
	var err error
	var authInfo *AuthInfo

	switch in.Platform {
	case PlatformYiBan:
		authInfo, err = a.yiBanAuthInfo(in.Code)
	}

	if err != nil {
		return nil, v1.GetAuthInfoErrorFormat(err.Error())
	}

	// 获取授权信息
	auth := model.Auth{}
	if auth.OneByAuthId(ctx, authInfo.AuthId) == nil {
		return nil, v1.AlreadyBindError()
	}

	user := model.User{}
	if util.IsPhone(in.Username) {
		// todo 增加短信验证
		err = user.OneByPhone(ctx, in.Username)
	} else if util.IsEmail(in.Username) {
		if err := ctx.Captcha().VerifyEmail(BindEmailName, ctx.ClientIP(), in.CaptchaId, in.Captcha); err != nil {
			return nil, v1.CaptchaError()
		}
		err = user.OneByEmail(ctx, in.Username)
	} else {
		return nil, v1.UsernameFormatError()
	}

	// 查询不到用户信息
	if err != nil {
		return nil, v1.UsernameNotExistError()
	}

	// 绑定用户信息
	auth = model.Auth{
		AuthID:   authInfo.AuthId,
		UnionID:  authInfo.UnionId,
		Platform: in.Platform,
		UserID:   user.ID,
	}
	if err := auth.Create(ctx); err != nil {
		return nil, v1.BindError()
	}

	// 生成用户token
	token, err := ctx.JWT().NewToken(md.New(auth.UserID))
	if err != nil {
		return nil, v1.LoginError()
	}

	return &v1.BindReply{Token: token}, nil
}

func (a *Auth) SendBindEmail(ctx kratosx.Context, in *v1.SendBindEmailRequest) (*v1.SendBindEmailReply, error) {
	user := model.User{}
	if err := user.OneByEmail(ctx, in.Email); err != nil {
		return nil, v1.UsernameNotExistError()
	}

	resp, err := ctx.Captcha().Email(BindEmailName, ctx.ClientIP(), user.Email)
	if err != nil {
		return nil, v1.SendEmailCaptchaErrorFormat(err.Error())
	}

	return &v1.SendBindEmailReply{
		Uuid:   resp.ID(),
		Expire: uint32(resp.Expire().Seconds()),
	}, nil
}

func (a *Auth) yiBanAuthInfo(code string) (*AuthInfo, error) {
	authInfo := a.conf.Auth.YiBan
	src := util.HexToByte(code)
	iv := []byte(authInfo.AppId)
	key := []byte(authInfo.AppSecret)

	body, _ := openssl.AesCBCDecrypt(src, key, iv, openssl.ZEROS_PADDING)
	if body == nil {
		return nil, errors.New("decrypt error")
	}

	// 解析数据
	res := struct {
		VisitTime int64 `json:"visit_time"`
		VisitUser struct {
			UserId string `json:"userid"`
		} `json:"visit_user"`
	}{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &AuthInfo{
		AuthId: res.VisitUser.UserId,
		Time:   res.VisitTime,
	}, nil
}

// RefreshToken 刷新用户token
func (a *Auth) RefreshToken(ctx kratosx.Context) (*v1.RefreshTokenReply, error) {
	// 进行token续期
	token, err := ctx.JWT().Renewal(ctx.Ctx())
	if err != nil {
		return nil, v1.RefreshTokenErrorFormat(err.Error())
	}

	return &v1.RefreshTokenReply{
		Token: token,
	}, nil
}
