羽球之家架构设计
===============

1. 模版目录:

`/static/tpl/` 为模版目录.

2. Api接口:

	1. 用户接口:

		1. 添加用户

			`PUT /api/user/reg`
			{
				"username":string
				"password":string  //两次输入密码在前端校验
				"email":string
			}
		2. 用户登陆

			`POST /api/user/login`

		3. `GET /api/user`





　　