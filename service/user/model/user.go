package model

import (
	"HIMGo/pkg/baseModel"
	"HIMGo/pkg/fxerror"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	baseModel.UidModel
	NickName string
	Account  string `gorm:"not null"`
}
type Auth struct {
	baseModel.IdModel
	Uid        string `gorm:"type:uuid;not null"`
	LoginType  int    `gorm:"type:smallint;not null"`
	Identifier string `gorm:"not null"`
	Credential string `gorm:"not null"`
}

type UserModel struct {
	Db *gorm.DB
}

func (u *UserModel) GetUserInfo(uid string) (*User, error) {
	userD := &User{UidModel: baseModel.UidModel{Uid: uid}}
	err := u.Db.Take(&userD).Error
	if err != nil {
		logx.Errorf("uid：%v查找失败,err:%v", uid, err)
		return nil, fxerror.NewDefaultError("系统错误，用户不存在")
	}
	return userD, err
}

func (u *UserModel) PasswordLogin(accout, password string) (*User, error) {

	auth := Auth{Identifier: accout, LoginType: 0}
	err := u.Db.Take(&auth).Error
	userD := &User{Account: accout}
	if err != nil {
		//如果账户不存在就创建账户
		if errors.Is(err, gorm.ErrRecordNotFound) {

			err = u.register(userD, password)
			if err != nil {
				logx.Error(err)
				return nil, fxerror.NewDefaultError("用户创建失败")
			}
			return userD, nil
		}
		return nil, fxerror.NewDefaultError("用户名不存在")
	}
	if !comparePasswords(auth.Credential, password) {
		return nil, fxerror.NewDefaultError("密码错误")
	}
	err = u.Db.Find(&userD).Error
	if err != nil {
		logx.Errorf("auth表有数据，user表无数据；用户：", accout)
		return nil, fxerror.NewDefaultError("登录失败")
	}
	return userD, nil
}

//密码hash存储
func hashCode(text string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// 验证密码
func comparePasswords(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

func (u *UserModel) register(userD *User, password string) error {
	return u.Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(userD).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		fmt.Println(password)
		if err := tx.Create(&Auth{
			Uid:        userD.Uid,
			LoginType:  0,
			Identifier: userD.Account,
			Credential: hashCode(password),
		}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}
