羽球之家架构设计
===============

1. 模版目录:

`/static/tpl/` 为模版目录.

2. Api接口:
	
	1. 用户接口:  //用户先不做,api没写完

		1. 添加用户
            <pre>
			PUT /api/user/reg

			req:
			{
				"username":string,
				"password":string,  //两次输入密码在前端校验
				"email":string,
				"vfcodeid":string,
				"vfcode":string,
			}

			resp:
			{
		        "ok":bool,
				"errmsg":string, //如果不ok
			}
            </pre>
		2. 获取验证码
            <pre>
			GET /api/vfcode

			req:
			{}

			resp:
			{
				"token":string,
				"imgurl":string,
			}
            </pre>
		2. 用户登陆
            <pre>
			POST /api/user/login

			req:
			{
				"username":string,
				"password":string,
			}
			resp:
			{
				"ok":bool,
				"userid":int,
				"errmsg":string, // 如果不ok
			}
            </pre>
		3. 获取用户资料
            <pre>
			GET /api/user

			req:
			{}

			resp:
			{}
            </pre>
		4. 更新用户信息 
            <pre>
			PUT /api/user
            </pre>
		4. 删除用户
			
			`DELETE /api/user`

		6. 用户消息接口

			`GET /api/user/message`

	2. 页面内容接口
		
		1. 通用接口

			1. 获取主页菜单栏
                <pre>
				GET /api/common/mainBar

				req:
				{}

				resp:
				[
					{
						"title":string,
						"url":string,
						"subtitle":
							[
								{
									"title":string,
									"url":string,
								}
							]
					}
				]
                </pre>
			2. 上传文件接口
                <pre>
				PUT /api/common/upload

				req:
				{
					"type":string,
					"file":formdata,
				}

				resp:
				{
					"ok":bool,
					"errmsg":string //if error
					"fileId":string //if not error
				}
                </pre>
			3. 删除文件接口
                <pre>
				DELETE /api/common/upload

				req:
				{}

				resp:
				{
					"ok":bool,
					"errmsg":string, //if not ok
				}
				</pre>
			4. 获得文件接口
				<pre>
				GET /api/common/upload?id=xxxx

				req:{}
				resp:
				{
					"ok":bool,
					"errmsg":string //if not ok,
					"fileid":int ,
					"filepath": string,
					"filename":string
				}
		2. 主页接口

			1. 获得主页滚动公告接口
				<pre>
				GET /api/common/carousel?num=xxx
				</pre>

		3. 学打羽毛球接口

			1. 拉取文章列表接口

				`GET /api/learn/list` ?type ?range 

			2. 查看特定文章接口

				`GET /api/learn/article`

			3. 文章搜索接口

				`GET /api/learn/search` ?time ?range ?title
			
		4. 促销接口 

			1. 促销信息获取接口

				`GET /api/deal/list`

		5. 视频库接口

			1. 获取视频库列表接口

				`GET /api/video/list` ?type ?range ?sort ?sorttype

			2. 获取视频详细信息接口

				`GET /api/video/video`

		6. 广告接口

			1. 获取广告




　　