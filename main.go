package main

import(
	"fmt"
	"github.com/pkg/errors"
)

var (
    	ErrDataNotFound = errors.New("record not found")
    )

func main(){
	service := &UserService{}
	_, err := service.FindUserByID(uint64(1))
	if err != nil{
	fmt.Printf("%+v", err)
	}
}

type User struct {
        Id uint64 `json:"Id"`
        Name string `json:"name"`
        Age  int32  `json:"age"`
}

type Dao interface {
        Get(id uint64) interface{}
        List() interface{}
        Create() 
        Update()
        Delete(id uint64) 
}
 
// Dao层获取到底层错误，使用errors的Wrap进行包装
type UserDao struct {}
func(dao *UserDao) Get(id uint64) (*User, error) {
        user := &User{}
        //err := db.Where("id = ?",id).Find(user).Error
	//if err != nil{
        return user, errors.Wrapf(ErrDataNotFound, "error getting the result with user_id %d", id)
	//}
}

// 业务层获取到错误直接往上层抛
type UserService struct {}
    
func (s *UserService) FindUserByID(user_id uint64) (*User, error) {
	dao := &UserDao{}
	user, err := dao.Get(user_id)
	if err != nil{
	if errors.Is(err, ErrDataNotFound){
	return user, err
	}
	}
return user, nil
}
