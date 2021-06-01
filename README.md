# WEB课程大作业
## 小组成员任务分工（按字母顺序）
+ 倪嘉伟——组长&中期汇报
+ 陈君鹏——论坛后端技术实现
+ 兰佳薇——task3&task9
+ 刘琼———task8&前端debug&github提交
+ 李俞颉——task5&task11&最终汇报PPT
+ 孙文昕——功能及接口说明&前端debug&task2
+ 朱锋———中期汇报PPT&task7&task12
+ 张耀仁——前端代码框架&结课展示
## 使用技术
+ 本论坛采用JAVA+Web前端+MySQL编写
+ 使用Tomcat——一个免费的开放源代码的Web 应用服务器，属于轻量级应用服务器。
+ Apache 为HTML页面服务，而Tomcat 实际上运行JSP 页面和Servlet。
+ 使用war exploded模式这种热部署的方式。（直接把文件夹、jsp页面 、classes等等移到Tomcat 部署文件夹里面，进行加载部署。）
+ Hibernate是一个开放源代码的对象关系映射框架，它对JDBC进行了非常轻量级的对象封装。
+ SLF4J是为各种loging APIs提供一个简单统一的接口，从而使得最终用户能够在部署的时候配置自己希望的loging APIs实现。 
## 业务流程及相关截图
+ 本系统有两类人群可以使用：用户&管理员
+ 游客页面：可以浏览帖子，但无法发帖和回复帖子。
可以进行帖子搜索（使用模糊搜索）
当需要回复帖子时，页面跳转至登录界面。
 ![image](https://user-images.githubusercontent.com/80044424/120337629-11382d80-c326-11eb-89ca-a0efb2996fbd.png)
 
 
+ 注册
 ![image](https://user-images.githubusercontent.com/80044424/120337687-1e551c80-c326-11eb-8222-ac8c9c7f3643.png)
限制特定项目不为空，密码相同重复，邮箱是否重复和有效。
+ 登录
 ![image](https://user-images.githubusercontent.com/80044424/120337721-2614c100-c326-11eb-9853-5382a43dc1c5.png)
登录操作可以和数据库进行匹配，成功即可登录。
登陆后可以进行：
+ 修改个人信息&查看我的帖子和申请精华帖，管理员查看申请后认同。相比游客而言，增加了发布帖子和回复帖子。
 ![image](https://user-images.githubusercontent.com/80044424/120337754-2e6cfc00-c326-11eb-81b0-778027ebb54e.png)
![image](https://user-images.githubusercontent.com/80044424/120337776-3331b000-c326-11eb-9e37-3c056d7c0d0b.png)
![image](https://user-images.githubusercontent.com/80044424/120337795-375dcd80-c326-11eb-8324-7354b554bfdb.png) 
+ 管理员页面：登录
 ![image](https://user-images.githubusercontent.com/80044424/120337846-42186280-c326-11eb-9f75-24bcb55c1305.png)
![image](https://user-images.githubusercontent.com/80044424/120337874-48a6da00-c326-11eb-8463-f2af322be512.png)

 
+ 管理员可以发布公告
 ![image](https://user-images.githubusercontent.com/80044424/120337900-4e042480-c326-11eb-8bef-c6628f7167ac.png)
+ 修改资料
 ![image](https://user-images.githubusercontent.com/80044424/120337995-63794e80-c326-11eb-809f-5b4ff2192c30.png)
+ 查看帖子或删除帖子，同意精华帖的申请操作。
 ![image](https://user-images.githubusercontent.com/80044424/120338033-6bd18980-c326-11eb-8c5d-6b00ad5de347.png)
+ 通过模糊搜索得到账号，拉黑用户
 ![image](https://user-images.githubusercontent.com/80044424/120338079-755af180-c326-11eb-9947-6dc17182c554.png)
+ 创建讨论区
![image](https://user-images.githubusercontent.com/80044424/120338102-7be96900-c326-11eb-86fb-0ec160e97bff.png)
+ 后台数据库ER图
 ![image](https://user-images.githubusercontent.com/80044424/120338127-8277e080-c326-11eb-9a3b-f81bf049741f.png)
