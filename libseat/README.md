1. 用户信息:
id,用户名，电话，密码，admin，信誉分,头像
2. 座位信息:
id,分馆,楼层,编号,Num,状态,图片,二维码
3. 座位预约签到表:
用户id，座位id，开始时间，结束时间，签到，签退，is预约
4. 反馈表:
用户id，管理员id，用户信息，管理员信息，isok


普通用户功能:
1. 注册成为普通用户
    手机号注册短信验证，注册成功跳转到登录界面，失败在次界面不变
2. 普通用户找回密码
    忘记密码使用手机号验证码验证，成功则重新设置密码，之后退出在进行登录操作
3. 普通用户登录
    用户名密码登录或手机号密码登录
    登录成功设置Cookie，之后均携带CooKie进行操作
4. 普通用户记住密码  --不实现
    <!-- 选择记住密码使得登录后的Cookie保存30天
    30 日(day)[d] = 2592000 秒(second)[s] -->
    不选则是保存一周
    1 周(week) = 604800 秒(second)[s]
5. 查看个人信息

 3. 普通用户修改个人信息
id作为唯一键，手机号不可以重复注册。
可以修改用户的用户名，密码，头像图片，不可修改手机号，admin和integral，修改确认后不需要重新登陆

13. 普通用户登出账号
账号登出，删除Cookie

<!-- -------------------------------------------------------------------seat -->
4. 普通用户查看座位信息
可查看目前总馆分馆剩余空位，点开可查看每层楼剩余空位
        1.  查看图书馆数量 
        2.  查看各图书馆位量总数和空余座位量
        3.  查看各个图书馆楼层数量
        4. 




5. 普通用户自主选择座位 
可以输入条件查看座位信息，条件如下:
主馆/分馆   可选
楼层  可选
时间  必选
？？？？？？？？
6. 普通用户推荐算法选择座位
推荐算法:通过排序
7. 普通用户取消座位预约
在签到时间前10分钟之内可以取消预约，否则取消预约扣除信誉分2分
8. 普通用户座位扫码签到
扫码签到
9. 普通用户座位扫码签退
扫码签退
10.普通用户查看个人历史记录
        已完成的
        过程中的
        未完成的(用户信用过低禁止预约座位)
        已取消的
11. 普通用户信息反馈
用户发送信息给管理员，申报反馈问题
12. 普通用户接收回复消息
接收管理员返回信息解决情况，使用消息提示

14. 普通用户查看应用版权信息等


admin用户:
1. 登录为管理员
2. 创建座位信息
批量？？？，单个创建座位
3. 查看座位信息
查看座位信息同普通用户
4. 修改座位信息
座位空闲时，设置座位状态为修改，并修改作为信息，id不变
5. 删除座位信息
座位空闲时，设置座位状态为修改，并删除座位


6. 回复用户消息
接收提醒，回复用户消息，可扣除，增加用户信誉积分
7. 管理用户信息信誉积分