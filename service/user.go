package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"superMarket-backend/common"
	"superMarket-backend/dao"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"time"
)

func UserLogin(c *gin.Context, user *dto.User) {
	var userDao dao.UserDao = &dao.UserDaoImpl{}
	employee := userDao.SelectOneByUsername(user.Username)
	if employee.Phone == "" {
		response.Error(c, "用户不存在")
		return
	}

	rdb := common.GetRDB()
	ctx := context.Background()

	//判断账户是否被冻结
	if validateFailUser(user.Username, rdb, ctx) {
		response.Error(c, "该账户已被冻结，请6小时后再来登录")
		return
	}

	//比对密码是否一致
	if !(user.Password == employee.Password) {
		//密码错误
		validateUserPassword(user.Username, rdb, ctx, c)
		return
	}

	//登录成功
	//删除密码错误次数缓存
	rdb.Del(ctx, fmt.Sprintf("%s%s", common.LOGIN_ERRO_PWDNUM, user.Username))
	if employee.IsAdmin {
		//超级管理员处理方式
		allMenu := menuDao.SelectAllMenu()
		employee.Menus = allMenu
	} else {
		//非超级管理员处理
		allMenu := menuDao.SelectAllMenu()
		employee.Menus = allMenu
	}
	//生成token的key和value
	key := fmt.Sprintf("%s%s", common.LOGIN_USER, user.Username)
	b, err := json.Marshal(employee)
	if err != nil {
		log.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	value := fmt.Sprintf("%s", b)
	//存入redis缓存中
	rdb.Set(ctx, key, value, 1440*time.Minute).Err()
	var maps = make(map[string]interface{})
	maps["token"] = key
	employee.Menus = nil
	employee.UserName = employee.Phone
	employee.LeaveTime = time.Time{}
	maps["employee"] = employee
	response.Success(c, maps, "登录成功")
}

func CheckedToken(c *gin.Context, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	var maps = make(map[string]interface{})
	var hasKey = false
	_, err := rdb.Get(ctx, token).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			hasKey = false
		}
	} else {
		hasKey = true
	}
	if hasKey {
		val, err := rdb.Get(ctx, token).Result()
		if err != nil {
			log.Println(err)
		}
		var employee model.Employee
		if err := json.Unmarshal([]byte(val), &employee); err != nil {
			response.Error(c, "token已过期需要重新登录")
			return
		}
		maps["employee"] = employee
	} else {
		maps["employee"] = nil
	}
	maps["token"] = token
	response.Success(c, maps, "操作成功")
}

func Logout(c *gin.Context, token string, content string) {
	if "本人确定注销" == content {
		//判断是否是系统管理员
		rdb := common.GetRDB()
		ctx := context.Background()
		val, err := rdb.Get(ctx, token).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				log.Println(err)
				response.Error(c, "fail")
			}
		}
		var employee model.Employee
		if err := json.Unmarshal([]byte(val), &employee); err != nil {
			response.Error(c, "token已过期需要重新登录")
			return
		}
		if employee.IsAdmin {
			response.Error(c, "系统管理员账户不可被注销")
			return
		}
		//清除角色员工关系
		var roleService IRoleService = RoleServiceImpl{}
		roleService.ClearEmpPermission(employee.ID)
		//清除缓存数据
		rdb.Del(ctx, token)
		//删除用户
		var userDao dao.UserDao = &dao.UserDaoImpl{}
		userDao.DeleteByUserName(employee.UserName)
		response.Success(c, "success", "操作成功")
		return
	} else {
		response.Error(c, "内容输入有误")
		return
	}
}
func EmpMenu(c *gin.Context, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		log.Println(err)
	}
	var employee model.Employee
	if err := json.Unmarshal([]byte(val), &employee); err != nil {
		response.Error(c, "token已过期需要重新登录")
		return
	}
	menus := employee.Menus
	response.Success(c, menus, "操作成功")
}

func Exit(c *gin.Context, token string) {
	rdb := common.GetRDB()
	ctx := context.Background()
	val, err := rdb.Get(ctx, token).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Println(err)
		}
	}
	var employee model.Employee
	if err := json.Unmarshal([]byte(val), &employee); err != nil {
		response.Success(c, nil, "操作成功")
		return
	} else {
		//密码错误
		errorPassKey := fmt.Sprintf("%s%s", common.LOGIN_ERRO_PWDNUM, employee.UserName)
		//删除token
		rdb.Del(ctx, token)
		rdb.Del(ctx, errorPassKey)
		response.Success(c, nil, "操作成功")
		return
	}
}

func validateFailUser(username string, rdb *redis.Client, ctx context.Context) bool {
	failUserKey := fmt.Sprintf("%s%s", common.DISABLEUSER, username)
	_, err := rdb.Get(ctx, failUserKey).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}

func validateUserPassword(username string, rdb *redis.Client, ctx context.Context, c *gin.Context) {
	//密码错误
	errorPassKey := fmt.Sprintf("%s%s", common.LOGIN_ERRO_PWDNUM, username)
	//账户冻结
	disableKey := fmt.Sprintf("%s%s", common.DISABLEUSER, username)
	//失败次数
	failCount := 1
	val, err := rdb.Get(ctx, errorPassKey).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			failCount = 1
		}
	} else {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Println("转换错误:", err)
		} else {
			log.Println("转换后的整数:", num)
		}
		failCount = num + 1
	}

	if failCount == 6 {
		rdb.Set(ctx, disableKey, true, 6*time.Hour)
		rdb.Del(ctx, errorPassKey)
		response.Error(c, "账户被冻结6小时")
		return
	} else {
		// 直接执行命令获取错误
		err = rdb.Set(ctx, errorPassKey, failCount, time.Hour).Err()
		response.Error(c, "账号或密码错误,错误剩余"+strconv.Itoa(6-failCount)+"次")
		return
	}
}
