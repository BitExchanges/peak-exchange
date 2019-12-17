# Poor Exchange

ApiBaseUrl:  /api/:platform/v1


> 用户
```bash
整体用户环节当前的步骤是。用户首先注册，会发送一封邮件验证码。验证码通过后会再次发送一封激活邮件
用户需要去邮箱点击激活链接后才能正常登录。
目前整体环节还未完善包括：

1.未做全局邮件发送开关，注册时邮件必填
2.邮件发送未做当天次数判断
3.登录会判断是否在常用登录地址范围，如果不在，会再次发送邮件提醒，未做地址归属地解析
4.发送邮件提醒用户登录地异常，需要发送链接确认地址正常。
```
- ApiBaseUrl/user/register           注册   √
- ApiBaseUrl/user/login              登录 √
- ApiBaseUrl/user/logout             退出登录  ×
- ApiBaseUrl/user/forgetLoginPwd     忘记登录密码 ×
- ApiBaseUrl/user/active             激活用户 √
- ApiBaseUrl/user/updateProfile      更新用户信息 ×
- ApiBaseUrl/user/changeLoginPwd     修改登录密码 ×
- ApiBaseUrl/user/changeTradePwd     修改交易密码 ×




