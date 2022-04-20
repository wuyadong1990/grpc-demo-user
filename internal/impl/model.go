package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/wuyadong1990/grpc-demo-user/cinit"
	"github.com/wuyadong1990/grpc-demo-user/internal/utils"

	"github.com/asaskevich/govalidator"
	"github.com/xiaomeng79/go-log"
	"gorm.io/gorm"
)

type User struct {
	ID         int64      `json:"id" db:"id" valid:"int~用户id类型为int"`
	UserName   string     `json:"user_name" db:"user_name" valid:"required~用户名称必须存在"`
	Password   string     `json:"password" db:"password" valid:"required~密码必须存在"`
	Iphone     string     `json:"iphone" db:"iphone" valid:"required~手机号码必须存在"`
	Sex        int32      `json:"sex" db:"sex" valid:"required~性别必须存在"`
	IsUsable   int32      `json:"-" db:"is_usable"`
	Page       utils.Page `gorm:"-" json:"-"`
	gorm.Model `json:"-"`
}

// 性别
const (
	SexMan   = 1
	SexWoman = 2
	SexOther = 3
)

// nolint: unused
var (
	sexTypes = []int32{
		SexMan,
		SexWoman,
		SexOther,
	}
)

// 从缓存获取数据
// nolint: unused
// 验证参数
func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

// 验证参数
func (m *User) validate() error {
	_, err := govalidator.ValidateStruct(m)
	return err
}

// 验证id
func (m *User) validateID() error {
	if m.ID <= 0 {
		return errors.New("id必须大于0")
	}
	return nil
}

// 验证性别类型
// nolint: unused
func (m *User) validateSexType() error {
	b := false
	for _, v := range sexTypes {
		if m.Sex == v {
			b = true
			break
		}
	}
	if !b {
		return errors.New("性别类型不合法")
	}
	return nil
}

// 添加之前
// nolint: unparam
func (m *User) beforeAdd(ctx context.Context) error {
	// 验证参数
	err := utils.V(m.validate)
	if err != nil {
		return err
	}
	return nil
}

// 修改之前
// nolint: unparam
func (m *User) beforeUpdate(ctx context.Context) error {
	err := utils.V(m.validate, m.validateID)
	if err != nil {
		return err
	}
	return nil
}

// 删除之前
// nolint: unparam
func (m *User) beforeDelete(ctx context.Context) error {
	err := utils.V(m.validateID)
	if err != nil {
		return err
	}
	return nil
}

// 添加之后,异步操作
func (m *User) afterAdd(ctx context.Context) error {
	//go msgNotify(ctx, "添加用户:"+m.UserName)
	go CacheSet(ctx, m.ID, m)
	return nil
}

// 修改之后,异步操作
func (m *User) afterUpdate(ctx context.Context) error {

	//go CacheDel(ctx, m.ID)
	// 修改缓存
	go CacheSet(ctx, m.ID, m)
	//go msgNotify(ctx, "修改用户:"+m.UserName)
	return nil
}

// 删除之后,异步操作
func (m *User) afterDelete(ctx context.Context) error {
	// 删除缓存
	go CacheDel(ctx, m.ID)
	return nil
}

// 添加
func (m *User) Add(ctx context.Context) error {
	err := m.beforeAdd(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	fmt.Printf("GormDB=%#v\n", cinit.GormDB)
	r := cinit.GormDB.Create(m)
	if r.Error != nil {
		log.Error(r.Error.Error(), ctx)
		return r.Error
	}
	return m.afterAdd(ctx)
}

// 修改
func (m *User) Update(ctx context.Context) error {
	err := m.beforeUpdate(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}
	r := cinit.GormDB.Model(m).Updates(m)
	if r.Error != nil {
		log.Error(r.Error.Error(), ctx)
		return r.Error
	}
	return m.afterUpdate(ctx)
}

// 删除
func (m *User) Delete(ctx context.Context) error {
	err := m.beforeDelete(ctx)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}

	r := cinit.GormDB.Delete(m)
	if r.Error != nil {
		log.Error(r.Error.Error(), ctx)
		return r.Error
	}

	return m.afterDelete(ctx)
}

// 查询一个
func (m *User) QueryOne(ctx context.Context) error {
	err := utils.V(m.validateID)
	if err != nil {
		log.Info(err.Error(), ctx)
		return err
	}

	/*缓存查询*/
	r, err := CacheGet(ctx, m.ID)
	if err == nil && len(r) > 0 {
		log.Debugf("[QueryOne] load from cache")
		utils.Map2Struct(r, m)
	} else {
		result := cinit.GormDB.First(m)
		err = result.Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info(err.Error(), ctx)
			return err
		}
	}
	return nil
}

// 查询全部
func (m *User) QueryAll(ctx context.Context) ([]*User, utils.Page, error) {
	var err error
	all := make([]*User, 0, m.Page.PageSize)

	tx := cinit.GormDB.Model(m)
	if m.Sex > 0 {
		tx.Where("sex = ?", m.Sex)
	}
	if len(m.UserName) > 0 {
		tx.Where("user_name = ?", m.UserName)
	}

	var total int64
	tx.Count(&total)
	log.Debugf("总页数:%v", total, ctx)

	m.Page.InitPage(total)
	// 加页数
	tx.Limit(int(m.Page.PageSize)).Offset(int((m.Page.PageIndex - 1) * m.Page.PageSize)).Find(&all)
	return all, m.Page, err
}
