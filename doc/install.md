# 安装版本

	go : version go1.6.2 darwin/amd64
	bee : version 1.5.1
	beego :version 1.7.1
	mongodb : version: 3.0.7
	
## 包安装

	session : go get github.com/astaxie/beego/session  //作用域包
	mongodb : go get github.com/goinggo/beego-mgo		//mongodb工具包
	redis : go get github.com/garyburd/redigo/redis	//redis工具包
	mysql : go get github.com/go-sql-driver/mysql		//mysql工具包
	orm : go get github.com/astaxie/beego/orm			//beego orm工具包
	paypal : go get github.com/logpacker/PayPal-Go-SDK   //Paypal支付工具包
	

## 工具安装

	beego : go get github.com/astaxie/beego
	bee : go get github.com/beego/bee
	 

## 架构默认设置

	* MaxMemory 文件上传默认内存缓存大小，默认值是 1 << 26(64M)。
	* AutoRender 自动渲染模板，默认值为 true，这里暂时设置为True（因为暂时还没有做界面开发这一块）， 对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。
	* FlashName Flash数据设置时Cookie的名称，默认是 "BEEGO_FLASH"，现在设置为 "DCONLINE_FLASH"
	* FlashSeperator Flash数据的分隔符，默认是 BEEGOFLASH , 现设置为 "DCONLINEFLASH"
	* DirectoryIndex  是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。现暂住设置为True(因为前期没有前端开发，所以先把静态文件目录开放出来先)
	* StaticDir 静态文件目录设置，默认是static
		可配置单个或多个目录:

		单个目录, StaticDir = download. 相当于beego.SetStaticPath("/download","download")
		多个目录, StaticDir = download:down download2:down2. 相当于beego.SetStaticPath("/download","down")和beego.SetStaticPath("/download2","down2")
		beego.BConfig.WebConfig.StaticDir

	* StaticExtensionsToGzip 允许哪些后缀名的静态文件进行gzip压缩，默认支持 .css 和 .js
		beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"} 等价config文件中 StaticExtensionsToGzip = .css, .js