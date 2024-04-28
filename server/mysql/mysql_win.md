=====================================================================
Windows使用管理员进入CMD
=====================================================================
https://downloads.mysql.com/archives/community/


D:\BaiduNetdiskDownload\mysql-5.7.44-winx64>
D:\BaiduNetdiskDownload\mysql-5.7.44-winx64>.\bin\mysqld.exe install
Service successfully installed.

D:\BaiduNetdiskDownload\mysql-5.7.44-winx64>.\bin\mysqld.exe --initialize --console
2024-04-27T07:52:04.101219Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
2024-04-27T07:52:04.827219Z 0 [Warning] InnoDB: New log files created, LSN=45790
2024-04-27T07:52:04.937264Z 0 [Warning] InnoDB: Creating foreign key constraint system tables.
2024-04-27T07:52:05.042638Z 0 [Warning] No existing UUID has been found, so we assume that this is the first time that this server has been started. Generating a new UUID: 0a2ac60d-046b-11ef-a79e-e0d55eee4cc8.
2024-04-27T07:52:05.047371Z 0 [Warning] Gtid table is not ready to be used. Table 'mysql.gtid_executed' cannot be opened.
2024-04-27T07:52:05.517640Z 0 [Warning] A deprecated TLS version TLSv1 is enabled. Please use TLSv1.2 or higher.
2024-04-27T07:52:05.517872Z 0 [Warning] A deprecated TLS version TLSv1.1 is enabled. Please use TLSv1.2 or higher.
2024-04-27T07:52:05.521338Z 0 [Warning] CA certificate ca.pem is self signed.
2024-04-27T07:52:05.683268Z 1 [Note] A temporary password is generated for root@localhost: kRt=8iCqwnY)

D:\BaiduNetdiskDownload\mysql-5.7.44-winx64>

D:\BaiduNetdiskDownload\mysql-5.7.44-winx64>c:\Windows\System32\net.exe start mysql

=====================================================================
=====================================================================




=====================================================================
重置密码
=====================================================================
第一步：在mysql安装目录中找到my.ini文件，在文档内搜索mysqld定位到[mysqld]文本段，在[mysqld]后面任意一行添加“skip-grant-tables”用来跳过密码验证的过程。具体如下：
skip-grant-tables

第二步：重启MySQL服务：
c:\Windows\System32\net.exe start mysql
c:\Windows\System32\net.exe stop mysql

这种情况通过重置密码即可。

第一步：重置密码的第一步就是跳过MySQL的密码认证过程，具体方法看情况一。
第二步：使用sql来修改root的密码，方法如下：
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'newpassword';
mysql> flush privileges;
mysql>  mysql > \q;
到这里root账户就已经重置成新的密码了。
第三步：修改my.cnf,去掉刚才添加的内容，然后重启MySQL。大功告成！
=====================================================================
=====================================================================


=====================================================================
数据库操作
=====================================================================
create DATABASE gva;
use gva;
SELECT * FROM sys_users;


