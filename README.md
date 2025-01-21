* [个人博客](https://www.codezhou.top)
* [CSDN](https://blog.csdn.net/cativen)

## 技能清单

1. 雪花算法

2. gin框架

3. Viper配置管理

4. 单点登录

5. Go语言操作MySQL **(gorm)**

6. Go语言操作Redis **(go-redis)**

7. Docker部署

8. Vue框架

9. ElementUI

10. axios

11. 泛型

12. 支付宝支付

13. 发送邮件

14. Excel下载



## 项目环境

前端：vue2  node版本 14.21.3

后端：go 版本 1.23.2

​			java 版本 1.8

## 项目目录结构

### 后端结构树

```bash
│  build.sh
│  Dockerfile
│  go.mod
│  go.sum
│  main
│  main.go
│  README.md
│  start.sh
│  tree.txt
│      
├─common
│      alipay.go
│      const.go
│      db.go
│      log.go
│      redis.go
│      
├─config
│      application.yml
│      
├─controller
│      dept.go
│      employee.go
│      exchange_point_products_records.go
│      goods.go
│      goods_category.go
│      login.go
│      member.go
│      menu.go
│      personal.go
│      point_product.go
│      role.go
│      sale_record.go
│      statistic_sale.go
│      storage_situation.go
│      store.go
│      store_in.go
│      store_out.go
│      supplier.go
│      
├─dao
│      dept.go
│      detail_sale_records.go
│      detail_store_goods.go
│      exchange_point_products.go
│      goods.go
│      goods_category.go
│      goods_store.go
│      member.go
│      menu.go
│      point_products.go
│      role.go
│      sale_records.go
│      store.go
│      supplier.go
│      user.go
│      
├─dto
│      base_query.go
│      dept_query.go
│      emp_query.go
│      menu_query.go
│      query_detail_storage_situation.go
│      query_detail_store_goods_out.go
│      query_edit_pwd.go
│      query_exchange_point_products_record.go
│      query_goods.go
│      query_goods_category.go
│      query_member.go
│      query_point_products.go
│      query_statistic_sale.go
│      query_storage_situation.go
│      query_store.go
│      query_supplier.go
│      role_query.go
│      sale_record_page.go
│      user.go
│      
├─log
│      nohup.log
│      
├─model
│      department.go
│      detail_sale_record.go
│      detail_store_goods.go
│      employee.go
│      emp_role.go
│      exchange_point_products.go
│      goods.go
│      goods_category.go
│      goods_store.go
│      member.go
│      menu.go
│      point_product.go
│      role.go
│      sale_record.go
│      store.go
│      supplier.go
│      
├─response
│      response.go
│      
├─route
│      routes.go
│      
├─service
│      dept.go
│      detail_store_goods.go
│      employee.go
│      exchange_point_products.go
│      goods.go
│      goods_category.go
│      goods_store.go
│      member.go
│      menu.go
│      point_products.go
│      role.go
│      sale_records.go
│      store.go
│      supplier.go
│      user.go
│      
├─sql
│      supermarket.sql
│      
├─static
│  └─img
│          1729582375884_screenshot-20240919-162435.png
│          1729595691328_programming.png
│          1729688055264_programming.png
│          
├─utils
│      convert_util.go
│      dateformat_util.go
│      email.go
│      page_util.go
│      path_util.go
│      snow_flake_util.go
│      
└─vo
        detail_storage_situation.go
        detail_store_goods_out.go
        goods_list.go
        information.go
        page.go
        sales_statistics.go
        sale_goods.go
        storage_situation.go
```

### 前端结构树

```bash
│  .gitignore
│  babel.config.js
│  jsconfig.json
│  package-lock.json
│  package.json
│  tree.txt
│  vue.config.js
│      
├─public
│      favicon.ico
│      index.html
│      
└─src
    │  App.vue
    │  main.js
    │  
    ├─api
    │  ├─goods_management
    │  │  ├─goods
    │  │  │      goodsApi.js
    │  │  │      
    │  │  ├─goods_category
    │  │  │      goodsCategoryApi.js
    │  │  │      
    │  │  ├─goods_store
    │  │  │      goodsStoreApi.js
    │  │  │      
    │  │  ├─point_products
    │  │  │      pointProductApi.js
    │  │  │      
    │  │  └─statistic
    │  │          statisticSaleApi.js
    │  │          
    │  ├─inventory_management
    │  │  ├─detail_store_goods
    │  │  │      detail_store_goodsApi.js
    │  │  │      
    │  │  ├─notice
    │  │  │      noticeApi.js
    │  │  │      
    │  │  ├─storage_situation
    │  │  │      storageSituationApi.js
    │  │  │      
    │  │  ├─store
    │  │  │      storeApi.js
    │  │  │      
    │  │  └─supplier
    │  │          supplierApi.js
    │  │          
    │  ├─login
    │  │      loginApi.js
    │  │      
    │  ├─member_management
    │  │  └─member
    │  │          memberApi.js
    │  │          
    │  ├─personal
    │  │      personApi.js
    │  │      
    │  ├─personnel_management
    │  │  ├─department
    │  │  │      deptApi.js
    │  │  │      
    │  │  └─employee
    │  │          empApi.js
    │  │          
    │  ├─sale_management
    │  │  │  saleRecordsApi.js
    │  │  │  
    │  │  └─exchange_point_products
    │  │          exchangePointProductsApi.js
    │  │          
    │  └─system
    │      ├─menu
    │      │      menuApi.js
    │      │      
    │      └─role
    │              roleApi.js
    │              
    ├─assets
    │  ├─css
    │  │      globel.css
    │  │      
    │  ├─font
    │  │      iconfont.css
    │  │      iconfont.js
    │  │      iconfont.json
    │  │      iconfont.ttf
    │  │      iconfont.woff
    │  │      iconfont.woff2
    │  │      
    │  ├─img
    │  │      01.jpg
    │  │      03.jpg
    │  │      13.jpg
    │  │      
    │  └─js
    │          auth.js
    │          common.js
    │          
    ├─router
    │      index.js
    │      
    ├─store
    │      index.js
    │      
    └─views
        │  Home.vue
        │  index.vue
        │  Login.vue
        │  
        ├─goods_management
        │  ├─goods
        │  │      List.vue
        │  │      
        │  ├─goods_category
        │  │      List.vue
        │  │      
        │  ├─goods_store
        │  │      List.vue
        │  │      
        │  ├─point_products
        │  │      List.vue
        │  │      
        │  └─statistic_sale
        │          List.vue
        │          
        ├─inventory_management
        │  ├─detail_store_goods_in
        │  │  │  List.vue
        │  │  │  
        │  │  └─notice
        │  │          List.vue
        │  │          
        │  ├─detail_store_goods_out
        │  │  │  List.vue
        │  │  │  
        │  │  └─notice
        │  │          List.vue
        │  │          
        │  ├─store
        │  │      List.vue
        │  │      StorageSituation.vue
        │  │      
        │  └─supplier
        │          List.vue
        │          
        ├─member_management
        │  └─member
        │          List.vue
        │          
        ├─personal
        │      Edit_pwd.vue
        │      Information.vue
        │      
        ├─personnel_management
        │  ├─dept
        │  │      List.vue
        │  │      
        │  └─employee
        │          List.vue
        │          
        ├─sale_management
        │  ├─exchange_point_products_records
        │  │      List.vue
        │  │      
        │  ├─sale_cmd
        │  │      List.vue
        │  │      
        │  └─sale_records
        │          List.vue
        │          
        └─system
            ├─menu
            │      List.vue
            │      
            └─role
                    List.vue
```

## 项目预览图

![捕获](https://codegym.oss-cn-shenzhen.aliyuncs.com/uiiujhj/202501212249512.PNG)

## 前端运行方法

```golang
npm install 

npm run serve
```



## 后端运行方法

```golang
go run main.go 
```



### docker部署

```bash
在docker环境下，先给build.sh 和 start.sh权限
chomd 777 build.sh
chomd 777 start.sh
执行脚本
./build.sh
./start.sh
```

