package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"regexp"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jwt "github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nagokos/connefut_backend/db"
	"github.com/nagokos/connefut_backend/graph/model"
	"github.com/nagokos/connefut_backend/graph/utils"
	"github.com/nagokos/connefut_backend/logger"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	host      = "mailhog:1025"
	resUser   model.User
	SecretKey = []byte("secretKey")
)

type User struct {
	ID                              string
	Name                            string
	Email                           string
	Password                        string
	EmailVerificationStatus         bool
	EmailVerificationToken          string
	EmailVerificationTokenExpiresAt time.Time
}

type NullableUser struct {
	ID     *string
	Name   *string
	Avatar *string
}

func checkExistsEmail() validation.RuleFunc {
	return func(v interface{}) error {
		var err error

		email := v.(string)
		dbPool := db.DatabaseConnection()

		cmd := "SELECT COUNT(DISTINCT id) FROM users WHERE email = $1"
		row := dbPool.QueryRow(context.Background(), cmd, email)

		var count int
		err = row.Scan(&count)

		if err != nil {
			logger.NewLogger().Error(err.Error())
			return err
		}

		if count == 1 {
			logger.NewLogger().Error("This email address is already exists")
			err = errors.New("このメールアドレスは既に存在します")
		}

		return err
	}
}

// ** validation **
func (u User) CreateUserValidate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.Required.Error("名前を入力してください"),
			validation.RuneLength(1, 20).Error("名前は50文字以内で入力してください"),
		),
		validation.Field(
			&u.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			validation.RuneLength(1, 100).Error("メールアドレスは100文字以内で入力してください"),
			validation.Match(regexp.MustCompile(`^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`)).
				Error("メールアドレスを正しく入力してください"),
			validation.By(checkExistsEmail()),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("パスワードを入力してください"),
			validation.RuneLength(8, 100).Error("パスワードは8~100文字で入力してください"),
			validation.Match(regexp.MustCompile("[a-z]")).Error("パスワードを正しく入力してください"),
			validation.Match(regexp.MustCompile(`\d`)).Error("パスワードを正しく入力してください"),
		),
	)
}

func (u User) AuthenticateUserValidate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			validation.Match(regexp.MustCompile(`^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`)).
				Error("メールアドレスを正しく入力してください"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("パスワードを入力してください"),
			validation.RuneLength(8, 100).Error("パスワードは8文字以上で入力してください"),
			validation.Match(regexp.MustCompile("[a-z]")).Error("パスワードを正しく入力してください"),
			validation.Match(regexp.MustCompile(`\d`)).Error("パスワードを正しく入力してください"),
		),
	)
}

// ** utils **
func HashGenerate(password string) string {
	b := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(b, 12)
	if err != nil {
		logger.NewLogger().Error(err.Error())
	}
	return string(hash)
}

func CheckPasswordHash(passwordDigest, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(password))
	return err
}

func (u *User) GenerateEmailVerificationToken() string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(u.Email)))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		logger.NewLogger().Error(err.Error())
		return "", err
	}

	return tokenString, nil
}

func SendVerifyEmail(emailToken string) error {
	verifyURL := fmt.Sprintf("http://localhost:8080/accounts/email_verification/%s", emailToken)
	message := strings.NewReader(verifyURL)
	transformer := japanese.ISO2022JP.NewEncoder()
	newMessage, _ := ioutil.ReadAll(transform.NewReader(message, transformer))
	err := smtp.SendMail(host, nil, "connefut@example.com", []string{"connefut@example.com"}, newMessage)
	return err
}

// ** データベース伴う処理 **
func (u *User) Insert(ctx context.Context, dbPool *pgxpool.Pool) (string, error) {
	var ID string

	pwdHash := HashGenerate(u.Password)
	emailToken := u.GenerateEmailVerificationToken()
	tokenExpiresAt := time.Now().Add(24 * time.Hour)

	cmd := `
		INSERT INTO users
			( id, name, email, password_digest, email_verification_token, 
				email_verification_token_expires_at, last_sign_in_at, created_at, updated_at
			) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING id
		`

	row := dbPool.QueryRow(
		ctx, cmd,
		xid.New().String(), u.Name, u.Email, pwdHash, emailToken, tokenExpiresAt, time.Now().Local(), time.Now().Local(), time.Now().Local(),
	)

	err := row.Scan(&ID)

	if err != nil {
		logger.NewLogger().Error(err.Error())
		return ID, err
	}

	// 本番環境と開発環境では違う
	err = SendVerifyEmail(emailToken)

	if err != nil {
		logger.NewLogger().Sugar().Errorf("fail send email: %s", err)
		return ID, err
	}

	return ID, nil
}

func (u *User) Authenticate(ctx context.Context, dbPool *pgxpool.Pool) (string, error) {
	var ID string
	var passwordDigest string

	cmd := "SELECT id, password_digest FROM users WHERE email = $1"
	row := dbPool.QueryRow(ctx, cmd, u.Email)

	err := row.Scan(&ID, &passwordDigest)

	if err != nil {
		logger.NewLogger().Sugar().Errorf("user not found: %s", err)
		utils.NewAuthenticationErorr("メールアドレスが正しくありません", utils.WithField("email")).AddGraphQLError(ctx)
		return ID, errors.New("フォームに不備があります")
	}

	err = CheckPasswordHash(passwordDigest, u.Password)
	if err != nil {
		logger.NewLogger().Sugar().Errorf("password is incorrect: %s", err)
		utils.NewAuthenticationErorr("パスワードが正しくありません", utils.WithField("password")).AddGraphQLError(ctx)
		return ID, errors.New("フォームに不備があります")
	}

	return ID, nil
}

// ** メール認証 **
func EmailVerification(w http.ResponseWriter, r *http.Request) {
	dbPool := db.DatabaseConnection()
	defer dbPool.Close()

	token := chi.URLParam(r, "token")
	if token == "" {
		w.Write([]byte("無効なURLです"))
		logger.NewLogger().Error("Invalid URL")
		return
	}

	ctx := context.Background()

	cmd := "SELECT id, email_verification_token_expires_at FROM users WHERE email_verification_token = $1"
	row := dbPool.QueryRow(ctx, cmd, token)

	var ID string
	var tokenExpiresAt time.Time
	err := row.Scan(&ID, &tokenExpiresAt)

	if err != nil {
		w.Write([]byte("ユーザーが見つかりませんでした"))
		logger.NewLogger().Sugar().Errorf("user not found: %s", err)
		return
	}

	if time.Now().After(tokenExpiresAt) {
		w.Write([]byte("有効期限が切れています"))
		logger.NewLogger().Error("token expires at expired")
		return
	}

	cmd = `
	  UPDATE users AS u
		SET email_verification_status = $1, email_verification_token = $2, updated_at = $3
		WHERE u.id = $4
	`
	_, err = dbPool.Exec(ctx, cmd, "verified", nil, time.Now().Local(), ID)
	if err != nil {
		logger.NewLogger().Error(err.Error())
		w.Write([]byte("メールアドレスの認証に失敗しました"))
		return
	}

	jwt, _ := CreateToken(ID)
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(time.Hour * 1),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "http://localhost:3000/", http.StatusMovedPermanently)
}
