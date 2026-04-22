package service

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tristaamne/flowershopbe-v4/common/config"
	"github.com/tristaamne/flowershopbe-v4/common/mailer"
	"github.com/tristaamne/flowershopbe-v4/common/pagination"
	"github.com/tristaamne/flowershopbe-v4/common/security/jwt"
	"github.com/tristaamne/flowershopbe-v4/common/security/otp"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/users/model"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	repo   repository.UserRepository
	cfg    *config.Config
	rdb    *redis.Client
	mailer mailer.Mailer
	jwt    jwt.JwtSrv
	otp    otp.OTP
}

type Service interface {
	DeleteUserById(ctx context.Context, id primitive.ObjectID) error
	UpdateUserById(ctx context.Context, req *model.UserRequest, userIDStr any) (primitive.ObjectID, error)
	EmailVerify(ctx context.Context, email, otp string) error
	GetUserByCondition(ctx context.Context, rawQuery map[string]interface{}, pg pagination.PaginationQuery) ([]model.User, error)
	Login(req any) (string, error)
	Register(ctx context.Context, req model.UserRequest) (primitive.ObjectID, error)
}

func NewUserService(repo repository.UserRepository, cfg *config.Config, rdb *redis.Client, mailer mailer.Mailer, jwtSrv jwt.JwtSrv, otp otp.OTP) Service {
	return &service{repo: repo, cfg: cfg, rdb: rdb, mailer: mailer, jwt: jwtSrv, otp: otp}
}

func (s *service) DeleteUserById(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	opts := options.Delete()

	err := s.repo.DeleteAUser(ctx, filter, opts)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateUserById(ctx context.Context, req *model.UserRequest, userIDStr any) (primitive.ObjectID, error) {
	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		return userID, fmt.Errorf("invalid ID")
	}

	filter := bson.M{"_id": userID}

	setData := bson.M{}
	if req.Name != "" {
		setData["name"] = req.Name
	}
	if req.Password != "" {
		hashedPassword, err := utils.PasswordHasher(req.Password)
		if err != nil {
			return userID, fmt.Errorf("invalid password")
		}
		setData["password"] = hashedPassword
	}
	if !req.Birthday.IsZero() {
		setData["birthday"] = req.Birthday
	}
	if req.Email != "" {
		if err := s.mailer.EmailValidate(req.Email); err != nil {
			return userID, fmt.Errorf("invalid email")
		}

		if err := s.mailer.OTPSender(ctx, req.Email); err != nil {
			return userID, fmt.Errorf(err.Error())
		}
		setData["email"] = req.Email
		setData["email_verified"] = false
	}

	if len(req.DeliveryAddresses) > 0 {
		setData["delivery_addresses"] = req.DeliveryAddresses
	}

	if len(setData) == 0 {
		return userID, fmt.Errorf("no field to update")
	}
	setData["updated_at"] = time.Now()

	update := bson.M{"$set": setData}
	er := s.repo.UpdateAUser(ctx, filter, update)
	if er != nil {
		return userID, fmt.Errorf(er.Error())
	}
	return userID, nil
}

func (s *service) EmailVerify(ctx context.Context, email, otp string) error {
	isValid, err := s.otp.VerifyOTP(ctx, email, otp)
	if err != nil || !isValid {
		return fmt.Errorf("OTP not match")
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"email_verified": true}}

	err = s.repo.UpdateAUser(ctx, filter, update)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	s.rdb.Del(ctx, email)
	return nil
}

func (s *service) GetUserByCondition(ctx context.Context, rawQuery map[string]interface{}, pg pagination.PaginationQuery) ([]model.User, error) {
	filter := utils.MapToBSon(rawQuery)

	excludedFields := []string{"page", "limit", "sort_by", "order"}
	for _, field := range excludedFields {
		delete(filter, field)
	}

	opts := pagination.ParsePagingOption(pg)

	userData, err := s.repo.GetUserByCondition(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *service) Login(req any) (string, error) {
	user, ok := req.([]model.User)
	if !ok {
		return "", fmt.Errorf("missing login context")
	}
	tokenString, err := s.jwt.GenerateToken(&user[0])
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *service) Register(ctx context.Context, req model.UserRequest) (primitive.ObjectID, error) {
	filter := bson.M{
		"email": req.Email,
	}

	er := s.mailer.EmailValidate(req.Email)
	if er != nil {
		return primitive.NilObjectID, fmt.Errorf("email format is invalid")
	}

	existingUser, err := s.repo.GetUserByCondition(ctx, filter, nil)
	if err == nil && existingUser != nil {
		return primitive.NilObjectID, fmt.Errorf("user already exists")
	}

	hashedPassword, err := utils.PasswordHasher(req.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}

	user := &model.User{
		Name:              req.Name,
		Password:          hashedPassword,
		Birthday:          req.Birthday,
		Email:             req.Email,
		DeliveryAddresses: req.DeliveryAddresses,
		Role:              0,
		ProviderID:        "manual",
		EmailVerified:     false,
	}
	id, err := s.repo.RegisterUser(ctx, user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid user")
	}

	err = s.mailer.OTPSender(ctx, user.Email)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}
