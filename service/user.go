package service

import (
	"fmt"
	"webGo/database"
	"webGo/model"
)

func Register(user model.UserLogin)(string,error)  {
	username:=user.UserName
	has,err:=database.Engine.Table("user").Where("user_name = ?",username).Exist()
	if nil!=err{
		return "", err
	}else {
		if has{
			return "用户已存在", nil
		}else {
			_,err:=database.Engine.Table("user").Insert(&user)
			if nil!=err{
				return "", err
			}else {
				return "注册成功",nil
			}
		}
	}

}




func UserLogin(user model.UserLogin) (string,error) {
	has, err := database.Engine.Table("user").Where("pwd=?", user.Pwd).And("user_name=?", user.UserName).Exist()
	if err != nil {
		fmt.Println("err",err)
		return "", err
	}else {
		if has {
			fmt.Println("success")
			return "登录成功", nil
		}else {
			return "用户名或密码错误", nil
		}
	}


	//if user.UserNmae=="admin" && user.Password=="123456"{
	//	return (map[string]interface{}{"msg": "登录成功"}), nil
	//}else {
	//	return (map[string]interface{}{"msg": "用户名或密码错误"}), nil
	//}

}