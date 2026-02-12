-- SQLite3 compatible SQL dump
-- Converted from MySQL dump for dilu-plus database

PRAGMA foreign_keys = OFF;

-- Table structure for table admin
DROP TABLE IF EXISTS admin;
CREATE TABLE admin (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT,
  phone TEXT,
  email TEXT,
  password TEXT,
  nickname TEXT,
  name TEXT,
  avatar TEXT,
  bio TEXT,
  birthday DATE,
  gender INTEGER DEFAULT 2,
  role_id INTEGER,
  dept_path TEXT,
  dept_id INTEGER,
  remark TEXT,
  lock_time DATETIME,
  status INTEGER,
  client_id TEXT,
  reg_ip TEXT,
  ip_location TEXT,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME
);

-- Dumping data for table admin
INSERT INTO admin VALUES (1,'diluplus',NULL,NULL,'$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a',NULL,'aaa',NULL,NULL,NULL,2,1,NULL,1,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,'2025-03-18 14:37:38.723');
INSERT INTO admin VALUES (2,'test','','','$2a$10$lHRg6AcfVs2xZ0kzJ5RohuBvrfFHGk6SikAuCgF4ZgMT1UKHqtTYe','','test','','',NULL,1,2,'/1/2/',2,'',NULL,1,'','','',0,'2025-03-19 11:45:21.606','2025-03-20 15:53:18.794');
INSERT INTO admin VALUES (3,'test2','','','$2a$10$zUzRSIFSzWH5/IXnv5OfVuBWbdv70Kn7Uy5gwN5H4aqY4iuO.0sHe','','test2','','',NULL,0,2,NULL,0,'',NULL,1,'','','',0,'2025-03-19 11:47:52.958','2025-03-19 11:47:52.958');

-- Table structure for table admin_dept
DROP TABLE IF EXISTS admin_dept;
CREATE TABLE admin_dept (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  parent_id INTEGER,
  dept_path TEXT,
  name TEXT,
  type INTEGER,
  principal TEXT,
  phone TEXT,
  email TEXT,
  sort INTEGER,
  status INTEGER,
  remark TEXT,
  create_by INTEGER,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table admin_dept
INSERT INTO admin_dept VALUES (1,0,'/1/','技术部',2,'','','',0,1,'',0,0,'2025-03-17 22:37:05.256','2025-03-17 22:37:05.256',NULL);
INSERT INTO admin_dept VALUES (2,1,'/1/2/','后端',2,'','','',0,1,'',0,0,'2025-03-17 22:37:25.663','2025-03-17 22:37:25.663',NULL);
INSERT INTO admin_dept VALUES (3,0,'/3/','产品部',2,'','','',0,1,'',0,0,'2025-03-17 22:37:50.104','2025-03-17 22:37:50.104',NULL);

-- Table structure for table admin_role
DROP TABLE IF EXISTS admin_role;
CREATE TABLE admin_role (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  role_key TEXT,
  role_sort INTEGER,
  status INTEGER,
  remark TEXT,
  create_by INTEGER,
  menu_ids TEXT,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table admin_role
INSERT INTO admin_role VALUES (1,'超级管理员','',0,0,'',0,'[401, 426, 430, 429, 428, 427, 432, 436, 435, 434, 433, 402, 406, 405, 404, 403]',1,'2025-03-17 22:36:35.607','2025-03-19 11:21:23.426',NULL);
INSERT INTO admin_role VALUES (2,'操作员','',0,0,'',1,'[394, 2, 6, 5, 4, 3, 52, 55, 53, 54, 56, 17, 21, 20, 19, 18, 12, 13, 16, 15, 14, 57, 58, 59, 60, 61, 117, 121, 120, 119, 118, 126, 127, 130, 129, 128, 396, 400, 399, 398, 397]',0,'2025-03-19 11:21:45.712','2025-03-19 11:21:45.712',NULL);

-- Table structure for table admin_role_menu
DROP TABLE IF EXISTS admin_role_menu;
CREATE TABLE admin_role_menu (
  role_id INTEGER NOT NULL,
  menu_id INTEGER NOT NULL,
  PRIMARY KEY (role_id, menu_id)
);

-- Dumping data for table admin_role_menu
INSERT INTO admin_role_menu VALUES (1,401);
INSERT INTO admin_role_menu VALUES (1,402);
INSERT INTO admin_role_menu VALUES (1,403);
INSERT INTO admin_role_menu VALUES (1,404);
INSERT INTO admin_role_menu VALUES (1,405);
INSERT INTO admin_role_menu VALUES (1,406);
INSERT INTO admin_role_menu VALUES (1,426);
INSERT INTO admin_role_menu VALUES (1,427);
INSERT INTO admin_role_menu VALUES (1,428);
INSERT INTO admin_role_menu VALUES (1,429);
INSERT INTO admin_role_menu VALUES (1,430);
INSERT INTO admin_role_menu VALUES (1,432);
INSERT INTO admin_role_menu VALUES (1,433);
INSERT INTO admin_role_menu VALUES (1,434);
INSERT INTO admin_role_menu VALUES (1,435);
INSERT INTO admin_role_menu VALUES (1,436);

-- Table structure for table gen_columns
DROP TABLE IF EXISTS gen_columns;
CREATE TABLE gen_columns (
  column_id INTEGER PRIMARY KEY AUTOINCREMENT,
  table_id INTEGER,
  column_name TEXT,
  column_comment TEXT,
  column_type TEXT,
  go_type TEXT,
  go_field TEXT,
  json_field TEXT,
  is_pk TEXT,
  is_increment TEXT,
  is_required TEXT,
  is_insert TEXT,
  is_edit TEXT,
  is_list TEXT,
  is_query TEXT,
  query_type TEXT,
  html_type TEXT,
  dict_type TEXT,
  sort INTEGER,
  list TEXT,
  pk INTEGER,
  required INTEGER,
  super_column INTEGER,
  usable_column INTEGER,
  increment INTEGER,
  insert INTEGER,
  edit INTEGER,
  query INTEGER,
  remark TEXT,
  fk_table_name TEXT,
  fk_table_name_class TEXT,
  fk_table_name_package TEXT,
  fk_label_id TEXT,
  fk_label_name TEXT,
  create_by INTEGER,
  update_By INTEGER,
  created_at DATETIME,
  updated_at DATETIME
);

-- Dumping data for table gen_columns
-- Note: Only showing first few rows due to length limitations
INSERT INTO gen_columns VALUES (1065,94,'column_id','','bigint','int','ColumnId','columnId','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.731','2025-03-15 16:16:37.731');
INSERT INTO gen_columns VALUES (1066,94,'table_id','','bigint','int','TableId','tableId','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.736','2025-03-15 16:16:37.736');
INSERT INTO gen_columns VALUES (1067,94,'column_name','','varchar(128)','string','ColumnName','columnName','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.743','2025-03-15 16:16:37.743');

-- Table structure for table gen_tables
DROP TABLE IF EXISTS gen_tables;
CREATE TABLE gen_tables (
  table_id INTEGER PRIMARY KEY AUTOINCREMENT,
  db_name TEXT,
  table_name TEXT,
  table_comment TEXT,
  class_name TEXT,
  tpl_category TEXT,
  package_name TEXT,
  module_name TEXT,
  module_front_name TEXT,
  business_name TEXT,
  function_name TEXT,
  function_author TEXT,
  pk_column TEXT,
  pk_go_field TEXT,
  pk_json_field TEXT,
  options TEXT,
  tree_code TEXT,
  tree_parent_code TEXT,
  tree_name TEXT,
  tree INTEGER DEFAULT 0,
  crud INTEGER DEFAULT 1,
  remark TEXT,
  is_data_scope INTEGER,
  is_actions INTEGER,
  is_auth INTEGER,
  is_logical_delete TEXT,
  logical_delete INTEGER,
  logical_delete_column TEXT,
  created_at DATETIME,
  updated_at DATETIME,
  create_by INTEGER,
  update_by INTEGER
);

-- Dumping data for table gen_tables
INSERT INTO gen_tables VALUES (94,'dilu-plus','gen_columns','生成列','GenColumns','crud','sys','gen-columns','','genColumns','生成列','baowk','column_id','ColumnId','columnId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:37.724','2025-03-15 16:16:37.724',0,0);
INSERT INTO gen_tables VALUES (95,'dilu-plus','gen_tables','生成表','GenTables','crud','sys','gen-tables','','genTables','生成表','baowk','table_id','TableId','tableId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.003','2025-03-15 16:16:38.003',0,0);
INSERT INTO gen_tables VALUES (96,'dilu-plus','sys_api','接口','SysApi','crud','sys','sys-api','','sysApi','接口','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.213','2025-03-15 16:16:38.213',0,0);

-- Table structure for table sys_api
DROP TABLE IF EXISTS sys_api;
CREATE TABLE sys_api (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  method TEXT,
  path TEXT,
  perm_type INTEGER,
  status INTEGER,
  update_by INTEGER,
  updated_at DATETIME
);

-- Dumping data for table sys_api
INSERT INTO sys_api VALUES (1,'分页获取用户','POST','/api/sys/user/page',3,3,0,'2023-09-26 13:46:59.488');
INSERT INTO sys_api VALUES (2,'根据id获取用户','POST','/api/sys/user/get',3,3,0,'2023-09-26 13:46:59.518');
INSERT INTO sys_api VALUES (3,'创建用户','POST','/api/sys/user/create',3,3,0,'2023-09-26 13:46:59.532');
INSERT INTO sys_api VALUES (4,'修改用户','POST','/api/sys/user/update',3,3,0,'2023-09-26 13:46:59.539');
INSERT INTO sys_api VALUES (5,'删除用户','POST','/api/sys/user/del',3,3,0,'2023-09-26 13:46:59.550');

-- Table structure for table sys_cfg
DROP TABLE IF EXISTS sys_cfg;
CREATE TABLE sys_cfg (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  type TEXT,
  key TEXT,
  val TEXT,
  remark TEXT,
  status INTEGER,
  team_id INTEGER,
  update_by INTEGER,
  updated_at DATETIME
);

-- Dumping data for table sys_cfg
INSERT INTO sys_cfg VALUES (1,NULL,'1111','222','2222','法大师傅',1,-1,1,'2025-03-21 17:23:35.936');

-- Table structure for table sys_email
DROP TABLE IF EXISTS sys_email;
CREATE TABLE sys_email (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  email TEXT,
  code TEXT,
  type TEXT,
  status INTEGER,
  use_status INTEGER,
  created_at INTEGER,
  updated_at INTEGER
);

-- Dumping data for table sys_email
-- No data to insert

-- Table structure for table sys_job
DROP TABLE IF EXISTS sys_job;
CREATE TABLE sys_job (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  job_name TEXT,
  job_group TEXT,
  job_type INTEGER,
  cron_expression TEXT,
  invoke_target TEXT,
  args TEXT,
  misfire_policy INTEGER,
  concurrent INTEGER,
  status INTEGER,
  entry_id INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME,
  create_by INTEGER,
  update_by INTEGER
);

-- Dumping data for table sys_job
-- No data to insert

-- Table structure for table sys_menu
DROP TABLE IF EXISTS sys_menu;
CREATE TABLE sys_menu (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  menu_name TEXT,
  title TEXT,
  icon TEXT,
  path TEXT,
  platform_type INTEGER,
  menu_type INTEGER,
  permission TEXT,
  parent_id INTEGER,
  no_cache INTEGER,
  component TEXT,
  sort INTEGER,
  hidden INTEGER,
  create_by INTEGER,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table sys_menu

INSERT INTO `sys_menu` VALUES (1,'','系统管理','setting','/sys',1,1,'',0,0,'Layout',50,0,1,0,'2023-09-26 13:46:59.480','2024-03-13 20:49:05.368',NULL);
INSERT INTO sys_menu VALUES (2,'SysUserManage','用户管理','flUser','/sys/sys-user',1,2,'sys:sysUser:list',394,0,'/sys/sys-user/index',0,0,1,1,'2023-09-26 13:46:59.493','2025-03-15 16:41:45.297',NULL);
INSERT INTO sys_menu VALUES (3,'','用户详情','','sys_user_detail',1,3,'sys:sysUser:query',2,0,'',0,0,1,1,'2023-09-26 13:46:59.522','2023-09-26 13:46:59.526',NULL);
INSERT INTO sys_menu VALUES (4,'','用户创建','','sys_user_create',1,3,'sys:sysUser:add',2,0,'',0,0,1,1,'2023-09-26 13:46:59.534','2023-09-26 13:46:59.536',NULL);
INSERT INTO sys_menu VALUES (5,'','用户修改','','sys_user_update',1,3,'sys:sysUser:edit',2,0,'',0,0,1,1,'2023-09-26 13:46:59.543','2023-09-26 13:46:59.546',NULL);
INSERT INTO sys_menu VALUES (6,'','用户删除','','sys_user_del',1,3,'sys:sysUser:remove',2,0,'',0,0,1,1,'2023-09-26 13:46:59.553','2023-09-26 13:46:59.555',NULL);
INSERT INTO sys_menu VALUES (7,'SysMenuManage','菜单管理','menu','/sys/sys-menu',1,2,'sys:sysMenu:list',91,0,'/sys/sys-menu/index',0,0,1,1,'2023-09-26 13:47:37.025','2024-05-24 22:47:05.611',NULL);
INSERT INTO sys_menu VALUES (8,'','菜单详情','','sys_menu_detail',1,3,'sys:sysMenu:query',7,0,'',0,0,1,1,'2023-09-26 13:47:37.042','2023-09-26 13:47:37.055',NULL);
INSERT INTO sys_menu VALUES (9,'','菜单创建','','sys_menu_create',1,3,'sys:sysMenu:add',7,0,'',0,0,1,1,'2023-09-26 13:47:37.067','2023-09-26 13:47:37.069',NULL);
INSERT INTO sys_menu VALUES (10,'','菜单修改','','sys_menu_update',1,3,'sys:sysMenu:edit',7,0,'',0,0,1,1,'2023-09-26 13:47:37.076','2023-09-26 13:47:37.079',NULL);
INSERT INTO sys_menu VALUES (11,'','菜单删除','','sys_menu_del',1,3,'sys:sysMenu:remove',7,0,'',0,0,1,1,'2023-09-26 13:47:37.084','2023-09-26 13:47:37.086',NULL);
INSERT INTO sys_menu VALUES (12,'SysRoleManage','角色管理','role','/sys/sys-role',1,2,'sys:sysRole:list',394,0,'/sys/sys-role/index',14,1,1,1,'2023-09-26 13:47:39.688','2025-03-17 16:18:52.560',NULL);
INSERT INTO sys_menu VALUES (13,'','角色详情','','sys_role_detail',1,3,'sys:sysRole:query',12,0,'',0,0,1,1,'2023-09-26 13:47:39.714','2023-09-26 13:47:39.717',NULL);
INSERT INTO sys_menu VALUES (14,'','角色创建','','sys_role_create',1,3,'sys:sysRole:add',12,0,'',0,0,1,1,'2023-09-26 13:47:39.747','2023-09-26 13:47:39.752',NULL);
INSERT INTO sys_menu VALUES (15,'','角色修改','','sys_role_update',1,3,'sys:sysRole:edit',12,0,'',0,0,1,1,'2023-09-26 13:47:39.765','2023-09-26 13:47:39.769',NULL);
INSERT INTO sys_menu VALUES (16,'','角色删除','','sys_role_del',1,3,'sys:sysRole:remove',12,0,'',0,0,1,1,'2023-09-26 13:47:39.776','2023-09-26 13:47:39.778',NULL);
INSERT INTO sys_menu VALUES (17,'SysDeptManage','部门管理','dept','/sys/sys-dept',1,2,'sys:sysDept:list',394,0,'/sys/sys-dept/index',12,1,1,1,'2023-09-26 13:47:42.140','2025-03-17 16:18:44.940',NULL);
INSERT INTO sys_menu VALUES (18,'','部门详情','','sys_dept_detail',1,3,'sys:sysDept:query',17,0,'',0,0,1,1,'2023-09-26 13:47:42.155','2023-09-26 13:47:42.167',NULL);
INSERT INTO sys_menu VALUES (19,'','部门创建','','sys_dept_create',1,3,'sys:sysDept:add',17,0,'',0,0,1,1,'2023-09-26 13:47:42.180','2023-09-26 13:47:42.182',NULL);
INSERT INTO sys_menu VALUES (20,'','部门修改','','sys_dept_update',1,3,'sys:sysDept:edit',17,0,'',0,0,1,1,'2023-09-26 13:47:42.188','2023-09-26 13:47:42.190',NULL);
INSERT INTO sys_menu VALUES (21,'','部门删除','','sys_dept_del',1,3,'sys:sysDept:remove',17,0,'',0,0,1,1,'2023-09-26 13:47:42.200','2023-09-26 13:47:42.202',NULL);
INSERT INTO sys_menu VALUES (52,'SysTeamManage','团队管理','team','/sys/sys-team',1,2,'sys:sysTeam:list',394,0,'/sys/sys-team/index',10,0,1,1,'2023-09-29 08:44:07.932','2025-03-15 16:41:55.374',NULL);
INSERT INTO sys_menu VALUES (53,'','团队详情','','sys_team_detail',1,3,'sys:sysTeam:query',52,0,'',0,0,1,1,'2023-09-29 08:44:07.941','2023-09-29 08:44:07.945',NULL);
INSERT INTO sys_menu VALUES (54,'','团队创建','','sys_team_create',1,3,'sys:sysTeam:add',52,0,'',0,0,1,1,'2023-09-29 08:44:07.951','2023-09-29 08:44:07.954',NULL);
INSERT INTO sys_menu VALUES (55,'','团队修改','','sys_team_update',1,3,'sys:sysTeam:edit',52,0,'',0,0,1,1,'2023-09-29 08:44:07.958','2023-09-29 08:44:07.961',NULL);
INSERT INTO sys_menu VALUES (56,'','团队删除','','sys_team_del',1,3,'sys:sysTeam:remove',52,0,'',0,0,1,1,'2023-09-29 08:44:07.966','2023-09-29 08:44:07.971',NULL);
INSERT INTO sys_menu VALUES (57,'SysMemberManage','成员管理','openArm','/sys/sys-member',1,2,'sys:sysMember:list',394,0,'/sys/sys-member/index',16,1,1,1,'2023-09-29 08:44:33.484','2025-03-17 16:18:58.556',NULL);
INSERT INTO sys_menu VALUES (58,'','成员详情','','sys_member_detail',1,3,'sys:sysMember:query',57,0,'',0,0,1,1,'2023-09-29 08:44:33.517','2023-09-29 08:44:33.520',NULL);
INSERT INTO sys_menu VALUES (59,'','成员创建','','sys_member_create',1,3,'sys:sysMember:add',57,0,'',0,0,1,1,'2023-09-29 08:44:33.526','2023-09-29 08:44:33.529',NULL);
INSERT INTO sys_menu VALUES (60,'','成员修改','','sys_member_update',1,3,'sys:sysMember:edit',57,0,'',0,0,1,1,'2023-09-29 08:44:33.534','2023-09-29 08:44:33.538',NULL);
INSERT INTO sys_menu VALUES (61,'','会员删除','','sys_member_del',1,3,'sys:sysMember:remove',57,0,'',0,0,1,1,'2023-09-29 08:44:33.543','2023-09-29 08:44:33.546',NULL);
INSERT INTO sys_menu VALUES (91,'','工具管理','setUp','/tools',1,1,NULL,0,0,'Layout',111,0,1,1,'2023-10-28 09:41:26.000','2024-05-24 22:47:52.057',NULL);
INSERT INTO sys_menu VALUES (93,'GenTablesManage','代码生成','table','/sys/gen-tables',1,2,'sys:genTables:list',91,0,'/tool/gen-tables/index',80,0,1,1,'2023-10-28 16:31:47.693','2025-03-15 16:36:31.947',NULL);
INSERT INTO sys_menu VALUES (94,'','GenTables详情','','gen_tables_detail',1,3,'sys:genTables:query',93,0,'',0,0,1,1,'2023-10-28 16:31:47.704','2023-10-28 16:31:47.704',NULL);
INSERT INTO sys_menu VALUES (95,'','GenTables创建','','gen_tables_create',1,3,'sys:genTables:add',93,0,'',0,0,1,1,'2023-10-28 16:31:47.714','2023-10-28 16:31:47.714',NULL);
INSERT INTO sys_menu VALUES (96,'','GenTables修改','','gen_tables_update',1,3,'sys:genTables:edit',93,0,'',0,0,1,1,'2023-10-28 16:31:47.723','2023-10-28 16:31:47.723',NULL);
INSERT INTO sys_menu VALUES (97,'','GenTables删除','','gen_tables_del',1,3,'sys:genTables:remove',93,0,'',0,0,1,1,'2023-10-28 16:31:47.731','2023-10-28 16:31:47.731',NULL);
INSERT INTO sys_menu VALUES (103,'SysApiManage','接口管理','swapLine','/sys/sys-api',1,2,'sys:sysApi:list',91,0,'/tool/sys-api/index',90,0,1,1,'2023-10-28 16:51:44.924','2025-03-15 16:36:24.788',NULL);
INSERT INTO sys_menu VALUES (104,'','接口详情','','sys_api_detail',1,3,'sys:sysApi:query',103,0,'',0,0,1,1,'2023-10-28 16:51:44.938','2023-10-28 16:51:44.938',NULL);
INSERT INTO sys_menu VALUES (105,'','接口创建','','sys_api_create',1,3,'sys:sysApi:add',103,0,'',0,0,1,1,'2023-10-28 16:51:44.947','2023-10-28 16:51:44.947',NULL);
INSERT INTO sys_menu VALUES (106,'','接口修改','','sys_api_update',1,3,'sys:sysApi:edit',103,0,'',0,0,1,1,'2023-10-28 16:51:44.955','2023-10-28 16:51:44.955',NULL);
INSERT INTO sys_menu VALUES (107,'','接口删除','','sys_api_del',1,3,'sys:sysApi:remove',103,0,'',0,0,1,1,'2023-10-28 16:51:44.964','2023-10-28 16:51:44.964',NULL);
INSERT INTO sys_menu VALUES (111,'SysCfgManage','配置管理','operation','/sys/sys-cfg',1,2,'sys:sysCfg:list',1,0,'/sys/sys-cfg/index',21,0,1,1,'2023-11-08 20:34:50.875','2025-03-15 16:31:48.457',NULL);
INSERT INTO sys_menu VALUES (112,'','配置详情','','sys_cfg_detail',1,3,'sys:sysCfg:query',111,0,'',0,0,1,1,'2023-11-08 20:34:50.889','2023-11-08 20:34:50.889',NULL);
INSERT INTO sys_menu VALUES (113,'','配置创建','','sys_cfg_create',1,3,'sys:sysCfg:add',111,0,'',0,0,1,1,'2023-11-08 20:34:50.903','2023-11-08 20:34:50.903',NULL);
INSERT INTO sys_menu VALUES (114,'','配置修改','','sys_cfg_update',1,3,'sys:sysCfg:edit',111,0,'',0,0,1,1,'2023-11-08 20:34:50.916','2023-11-08 20:34:50.916',NULL);
INSERT INTO sys_menu VALUES (115,'','配置删除','','sys_cfg_del',1,3,'sys:sysCfg:remove',111,0,'',0,0,1,1,'2023-11-08 20:34:50.929','2023-11-08 20:34:50.929',NULL);
INSERT INTO sys_menu VALUES (117,'SysOperaLogManage','操作日志管理','generate','/sys/sys-opera-log',1,2,'sys:sysOperaLog:list',394,0,'/sys/sys-opera-log/index',19,0,1,1,'2023-11-09 22:07:17.257','2025-03-15 16:42:31.824',NULL);
INSERT INTO sys_menu VALUES (118,'','操作日志详情','','sys_opera_log_detail',1,3,'sys:sysOperaLog:query',117,0,'',0,0,1,1,'2023-11-09 22:07:17.272','2023-11-09 22:07:17.272',NULL);
INSERT INTO sys_menu VALUES (119,'','操作日志创建','','sys_opera_log_create',1,3,'sys:sysOperaLog:add',117,0,'',0,0,1,1,'2023-11-09 22:07:17.288','2023-11-09 22:07:17.288',NULL);
INSERT INTO sys_menu VALUES (120,'','操作日志修改','','sys_opera_log_update',1,3,'sys:sysOperaLog:edit',117,0,'',0,0,1,1,'2023-11-09 22:07:17.301','2023-11-09 22:07:17.301',NULL);
INSERT INTO sys_menu VALUES (121,'','操作日志删除','','sys_opera_log_del',1,3,'sys:sysOperaLog:remove',117,0,'',0,0,1,1,'2023-11-09 22:07:17.313','2023-11-09 22:07:17.313',NULL);
INSERT INTO sys_menu VALUES (122,'My','我的',NULL,'/my',1,1,'',0,0,'Layout',123,1,1,1,'2023-11-18 15:25:44.000','2024-05-24 23:19:32.049',NULL);
INSERT INTO sys_menu VALUES (123,NULL,'修改企业信息',NULL,'my_team',1,3,'my:change:team',122,0,NULL,0,0,1,1,'2023-11-18 15:33:28.000','2023-11-18 15:33:31.000',NULL);
INSERT INTO sys_menu VALUES (124,'Monitor','监控','monitor','/tools/monitor',1,2,'sys:monitor',91,0,'/tool/monitor',100,0,1,1,'2023-11-20 19:46:04.000','2025-03-15 16:36:13.924',NULL);
INSERT INTO sys_menu VALUES (126,'UserNoticeManage','用户通知管理','bell','/notice/user-notice',1,2,'notice:userNotice:list',394,0,'/notice/user-notice/index',20,0,1,1,'2023-11-22 21:45:33.866','2025-03-15 16:42:50.590',NULL);
INSERT INTO sys_menu VALUES (127,'','用户通知详情','','user_notice_detail',1,3,'notice:userNotice:query',126,0,'',0,0,1,1,'2023-11-22 21:45:33.889','2023-11-22 21:45:33.889',NULL);
INSERT INTO sys_menu VALUES (128,'','用户通知创建','','user_notice_create',1,3,'notice:userNotice:add',126,0,'',0,0,1,1,'2023-11-22 21:45:33.907','2023-11-22 21:45:33.907',NULL);
INSERT INTO sys_menu VALUES (129,'','用户通知修改','','user_notice_update',1,3,'notice:userNotice:edit',126,0,'',0,0,1,1,'2023-11-22 21:45:33.925','2023-11-22 21:45:33.925',NULL);
INSERT INTO sys_menu VALUES (130,'','用户通知删除','','user_notice_del',1,3,'notice:userNotice:remove',126,0,'',0,0,1,1,'2023-11-22 21:45:33.941','2023-11-22 21:45:33.941',NULL);
INSERT INTO sys_menu VALUES (132,'PubNoticeManage','公用通知管理','bellFill','/notice/pub-notice',1,2,'notice:pubNotice:list',1,0,'/notice/pub-notice/index',28,0,1,1,'2023-11-22 21:45:36.391','2025-03-15 16:32:22.933',NULL);
INSERT INTO sys_menu VALUES (133,'','公用通知详情','','pub_notice_detail',1,3,'notice:pubNotice:query',132,0,'',0,0,1,1,'2023-11-22 21:45:36.405','2023-11-22 21:45:36.405',NULL);
INSERT INTO sys_menu VALUES (134,'','公用通知创建','','pub_notice_create',1,3,'notice:pubNotice:add',132,0,'',0,0,1,1,'2023-11-22 21:45:36.416','2023-11-22 21:45:36.416',NULL);
INSERT INTO sys_menu VALUES (135,'','公用通知修改','','pub_notice_update',1,3,'notice:pubNotice:edit',132,0,'',0,0,1,1,'2023-11-22 21:45:36.427','2023-11-22 21:45:36.427',NULL);
INSERT INTO sys_menu VALUES (136,'','公用通知删除','','pub_notice_del',1,3,'notice:pubNotice:remove',132,0,'',0,0,1,1,'2023-11-22 21:45:36.439','2023-11-22 21:45:36.439',NULL);
INSERT INTO sys_menu VALUES (138,'TaskManage','Task管理','task','/notice/task',1,2,'notice:task:list',1,0,'/notice/task/index',25,0,1,1,'2023-11-24 15:39:24.103','2025-03-15 16:32:14.006',NULL);
INSERT INTO sys_menu VALUES (139,'','Task详情','','task_detail',1,3,'notice:task:query',138,0,'',0,0,1,1,'2023-11-24 15:39:24.118','2023-11-24 15:39:24.118',NULL);
INSERT INTO sys_menu VALUES (140,'','Task创建','','task_create',1,3,'notice:task:add',138,0,'',0,0,1,1,'2023-11-24 15:39:24.131','2023-11-24 15:39:24.131',NULL);
INSERT INTO sys_menu VALUES (141,'','Task修改','','task_update',1,3,'notice:task:edit',138,0,'',0,0,1,1,'2023-11-24 15:39:24.142','2023-11-24 15:39:24.142',NULL);
INSERT INTO sys_menu VALUES (142,'','Task删除','','task_del',1,3,'notice:task:remove',138,0,'',0,0,1,1,'2023-11-24 15:39:24.155','2023-11-24 15:39:24.155',NULL);
INSERT INTO sys_menu VALUES (394,'','租户管理','team','/tenant',1,1,'',0,0,'Layout',30,0,1,1,'2023-09-26 13:46:59.480','2025-03-15 16:41:33.128',NULL);
INSERT INTO sys_menu VALUES (395,'','用户登录记录','log','/sys',1,1,'',0,0,'Layout',0,0,1,0,'2025-03-15 16:43:48.097','2025-03-15 16:43:48.097','2025-03-15 16:44:16.727');
INSERT INTO sys_menu VALUES (396,'SysUserLoginLogManage','用户登录记录管理','log','/sys/sys-user-login-log',1,2,'sys:sysUserLoginLog:list',394,0,'/sys/sys-user-login-log/index',50,0,1,1,'2025-03-15 16:43:48.112','2025-03-15 16:48:04.165',NULL);
INSERT INTO sys_menu VALUES (397,'','用户登录记录详情','','sys_user_login_log_detail',1,3,'sys:sysUserLoginLog:query',396,0,'',0,0,1,1,'2025-03-15 16:43:48.128','2025-03-15 16:43:48.128',NULL);
INSERT INTO sys_menu VALUES (398,'','用户登录记录创建','','sys_user_login_log_create',1,3,'sys:sysUserLoginLog:add',396,0,'',0,0,1,1,'2025-03-15 16:43:48.150','2025-03-15 16:43:48.150',NULL);
INSERT INTO sys_menu VALUES (399,'','用户登录记录修改','','sys_user_login_log_update',1,3,'sys:sysUserLoginLog:edit',396,0,'',0,0,1,1,'2025-03-15 16:43:48.169','2025-03-15 16:43:48.169',NULL);
INSERT INTO sys_menu VALUES (400,'','用户登录记录删除','','sys_user_login_log_del',1,3,'sys:sysUserLoginLog:remove',396,0,'',0,0,1,1,'2025-03-15 16:43:48.183','2025-03-15 16:43:48.183',NULL);
INSERT INTO sys_menu VALUES (401,'','内部管理','userSetting','/sys/internal',1,1,'',0,0,'Layout',30,0,1,1,'2025-03-17 09:25:50.030','2025-03-17 20:46:42.129',NULL);
INSERT INTO sys_menu VALUES (402,'AdminManage','超管管理','userSetting','/sys/admin',1,2,'sys:admin:list',401,0,'/sys/admin/index',0,0,1,1,'2025-03-17 09:25:50.050','2025-03-17 10:36:27.593',NULL);
INSERT INTO sys_menu VALUES (403,'','用户详情','','admin_detail',1,3,'sys:admin:query',402,0,'',0,0,1,1,'2025-03-17 09:25:50.068','2025-03-17 09:25:50.068',NULL);
INSERT INTO sys_menu VALUES (404,'','用户创建','','admin_create',1,3,'sys:admin:add',402,0,'',0,0,1,1,'2025-03-17 09:25:50.088','2025-03-17 09:25:50.088',NULL);
INSERT INTO sys_menu VALUES (405,'','用户修改','','admin_update',1,3,'sys:admin:edit',402,0,'',0,0,1,1,'2025-03-17 09:25:50.099','2025-03-17 09:25:50.099',NULL);
INSERT INTO sys_menu VALUES (406,'','用户删除','','admin_del',1,3,'sys:admin:remove',402,0,'',0,0,1,1,'2025-03-17 09:25:50.117','2025-03-17 09:25:50.117',NULL);
INSERT INTO sys_menu VALUES (407,'','系统管理','setting','/sys',3,1,'',0,0,'Layout',50,0,1,0,'2023-09-26 13:46:59.480','2024-03-13 20:49:05.368',NULL);
INSERT INTO sys_menu VALUES (408,'SysRoleManage','角色管理','role','/sys/sys-role',3,2,'sys:sysRole:list',407,0,'/sys/sys-role/index',14,0,1,1,'2023-09-26 13:47:39.688','2025-03-15 16:42:09.754',NULL);
INSERT INTO sys_menu VALUES (409,'','角色详情','','sys_role_detail',3,3,'sys:sysRole:query',408,0,'',0,0,1,1,'2023-09-26 13:47:39.714','2023-09-26 13:47:39.717',NULL);
INSERT INTO sys_menu VALUES (410,'','角色创建','','sys_role_create',3,3,'sys:sysRole:add',408,0,'',0,0,1,1,'2023-09-26 13:47:39.747','2023-09-26 13:47:39.752',NULL);
INSERT INTO sys_menu VALUES (411,'','角色修改','','sys_role_update',3,3,'sys:sysRole:edit',408,0,'',0,0,1,1,'2023-09-26 13:47:39.765','2023-09-26 13:47:39.769',NULL);
INSERT INTO sys_menu VALUES (412,'','角色删除','','sys_role_del',3,3,'sys:sysRole:remove',408,0,'',0,0,1,1,'2023-09-26 13:47:39.776','2023-09-26 13:47:39.778',NULL);
INSERT INTO sys_menu VALUES (413,'SysDeptManage','部门管理','dept','/sys/sys-dept',3,2,'sys:sysDept:list',407,0,'/sys/sys-dept/index',12,0,1,1,'2023-09-26 13:47:42.140','2025-03-15 16:42:02.161',NULL);
INSERT INTO sys_menu VALUES (414,'','部门详情','','sys_dept_detail',3,3,'sys:sysDept:query',413,0,'',0,0,1,1,'2023-09-26 13:47:42.155','2023-09-26 13:47:42.167',NULL);
INSERT INTO sys_menu VALUES (415,'','部门创建','','sys_dept_create',3,3,'sys:sysDept:add',413,0,'',0,0,1,1,'2023-09-26 13:47:42.180','2023-09-26 13:47:42.182',NULL);
INSERT INTO sys_menu VALUES (416,'','部门修改','','sys_dept_update',3,3,'sys:sysDept:edit',413,0,'',0,0,1,1,'2023-09-26 13:47:42.188','2023-09-26 13:47:42.190',NULL);
INSERT INTO sys_menu VALUES (417,'','部门删除','','sys_dept_del',3,3,'sys:sysDept:remove',413,0,'',0,0,1,1,'2023-09-26 13:47:42.200','2023-09-26 13:47:42.202',NULL);
INSERT INTO sys_menu VALUES (418,'','团队详情','','sys_team_detail',3,3,'sys:sysTeam:query',413,0,'',0,0,1,1,'2023-09-29 08:44:07.941','2023-09-29 08:44:07.945',NULL);
INSERT INTO sys_menu VALUES (419,'','团队修改','','sys_team_update',3,3,'sys:sysTeam:edit',413,0,'',0,0,1,1,'2023-09-29 08:44:07.958','2023-09-29 08:44:07.961',NULL);
INSERT INTO sys_menu VALUES (420,'SysMemberManage','成员管理','openArm','/sys/sys-member',3,2,'sys:sysMember:list',407,0,'/sys/sys-member/index',16,0,1,1,'2023-09-29 08:44:33.484','2025-03-15 16:42:17.377',NULL);
INSERT INTO sys_menu VALUES (421,'','成员详情','','sys_member_detail',3,3,'sys:sysMember:query',420,0,'',0,0,1,1,'2023-09-29 08:44:33.517','2023-09-29 08:44:33.520',NULL);
INSERT INTO sys_menu VALUES (422,'','成员创建','','sys_member_create',3,3,'sys:sysMember:add',420,0,'',0,0,1,1,'2023-09-29 08:44:33.526','2023-09-29 08:44:33.529',NULL);
INSERT INTO sys_menu VALUES (423,'','成员修改','','sys_member_update',3,3,'sys:sysMember:edit',420,0,'',0,0,1,1,'2023-09-29 08:44:33.534','2023-09-29 08:44:33.538',NULL);
INSERT INTO sys_menu VALUES (424,'','成员删除','','sys_member_del',3,3,'sys:sysMember:remove',420,0,'',0,0,1,1,'2023-09-29 08:44:33.543','2023-09-29 08:44:33.546',NULL);
INSERT INTO sys_menu VALUES (425,'','部门','dept','/sys',1,1,'',0,0,'Layout',0,0,1,0,'2025-03-17 20:44:16.535','2025-03-17 20:44:16.535','2025-03-17 20:45:33.142');
INSERT INTO sys_menu VALUES (426,'AdminDeptManage','部门管理','dept','/sys/admin-dept',1,2,'sys:adminDept:list',401,0,'/sys/admin-dept/index',0,0,1,1,'2025-03-17 20:44:16.556','2025-03-17 20:45:13.260',NULL);
INSERT INTO sys_menu VALUES (427,'','部门详情','','admin_dept_detail',1,3,'sys:adminDept:query',426,0,'',0,0,1,1,'2025-03-17 20:44:16.577','2025-03-17 20:44:16.577',NULL);
INSERT INTO sys_menu VALUES (428,'','部门创建','','admin_dept_create',1,3,'sys:adminDept:add',426,0,'',0,0,1,1,'2025-03-17 20:44:16.594','2025-03-17 20:44:16.594',NULL);
INSERT INTO sys_menu VALUES (429,'','部门修改','','admin_dept_update',1,3,'sys:adminDept:edit',426,0,'',0,0,1,1,'2025-03-17 20:44:16.611','2025-03-17 20:44:16.611',NULL);
INSERT INTO sys_menu VALUES (430,'','部门删除','','admin_dept_del',1,3,'sys:adminDept:remove',426,0,'',0,0,1,1,'2025-03-17 20:44:16.628','2025-03-17 20:44:16.628',NULL);
INSERT INTO sys_menu VALUES (431,'','角色','role','/sys',1,1,'',0,0,'Layout',0,0,1,0,'2025-03-17 20:44:28.381','2025-03-17 20:44:28.381','2025-03-17 20:45:30.165');
INSERT INTO sys_menu VALUES (432,'AdminRoleManage','角色管理','role','/sys/admin-role',1,2,'sys:adminRole:list',401,0,'/sys/admin-role/index',0,0,1,1,'2025-03-17 20:44:28.400','2025-03-17 20:45:26.529',NULL);
INSERT INTO sys_menu VALUES (433,'','角色详情','','admin_role_detail',1,3,'sys:adminRole:query',432,0,'',0,0,1,1,'2025-03-17 20:44:28.419','2025-03-17 20:44:28.419',NULL);
INSERT INTO sys_menu VALUES (434,'','角色创建','','admin_role_create',1,3,'sys:adminRole:add',432,0,'',0,0,1,1,'2025-03-17 20:44:28.439','2025-03-17 20:44:28.439',NULL);
INSERT INTO sys_menu VALUES (435,'','角色修改','','admin_role_update',1,3,'sys:adminRole:edit',432,0,'',0,0,1,1,'2025-03-17 20:44:28.459','2025-03-17 20:44:28.459',NULL);
INSERT INTO sys_menu VALUES (436,'','角色删除','','admin_role_del',1,3,'sys:adminRole:remove',432,0,'',0,0,1,1,'2025-03-17 20:44:28.477','2025-03-17 20:44:28.477',NULL);
INSERT INTO sys_menu VALUES (437,'teamDeptAll','部门管理','','/team/dept',3,3,'team:dept:list',413,0,'',0,0,1,0,'2025-03-21 09:21:51.880','2025-03-21 09:21:51.880','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (438,'UserNoticeManage','团队用户通知管理','bell','/notice/user-notice',3,2,'notice:userNotice:list',407,0,'/notice/user-notice/index',20,0,1,1,'2023-11-22 21:45:33.866','2025-03-15 16:42:50.590','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (439,'','团队用户通知详情','','user_notice_detail',3,3,'notice:userNotice:query',438,0,'',0,0,1,1,'2023-11-22 21:45:33.889','2023-11-22 21:45:33.889','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (440,'','团队用户通知创建','','user_notice_create',3,3,'notice:userNotice:add',438,0,'',0,0,1,1,'2023-11-22 21:45:33.907','2023-11-22 21:45:33.907','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (441,'','团队用户通知修改','','user_notice_update',3,3,'notice:userNotice:edit',438,0,'',0,0,1,1,'2023-11-22 21:45:33.925','2023-11-22 21:45:33.925','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (442,'','团队用户通知删除','','user_notice_del',3,3,'notice:userNotice:remove',438,0,'',0,0,1,1,'2023-11-22 21:45:33.941','2023-11-22 21:45:33.941','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (443,'PubNoticeManage','团队公用通知管理','bellFill','/notice/pub-notice',3,2,'notice:pubNotice:list',407,0,'/notice/pub-notice/index',28,0,1,1,'2023-11-22 21:45:36.391','2025-03-15 16:32:22.933','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (444,'','团队公用通知详情','','pub_notice_detail',3,3,'notice:pubNotice:query',443,0,'',0,0,1,1,'2023-11-22 21:45:36.405','2023-11-22 21:45:36.405','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (445,'','团队公用通知创建','','pub_notice_create',3,3,'notice:pubNotice:add',443,0,'',0,0,1,1,'2023-11-22 21:45:36.416','2023-11-22 21:45:36.416','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (446,'','团队公用通知修改','','pub_notice_update',3,3,'notice:pubNotice:edit',443,0,'',0,0,1,1,'2023-11-22 21:45:36.427','2023-11-22 21:45:36.427','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (447,'','团队公用通知删除','','pub_notice_del',3,3,'notice:pubNotice:remove',443,0,'',0,0,1,1,'2023-11-22 21:45:36.439','2023-11-22 21:45:36.439','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (448,'TaskManage','团队Task管理','task','/notice/task',3,2,'notice:task:list',407,0,'/notice/task/index',25,0,1,1,'2023-11-24 15:39:24.103','2025-03-15 16:32:14.006','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (449,'','团队Task详情','','task_detail',3,3,'notice:task:query',448,0,'',0,0,1,1,'2023-11-24 15:39:24.118','2023-11-24 15:39:24.118','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (450,'','团队Task创建','','task_create',3,3,'notice:task:add',448,0,'',0,0,1,1,'2023-11-24 15:39:24.131','2023-11-24 15:39:24.131','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (451,'','团队Task修改','','task_update',3,3,'notice:task:edit',448,0,'',0,0,1,1,'2023-11-24 15:39:24.142','2023-11-24 15:39:24.142','2025-03-21 09:41:23.104');
INSERT INTO sys_menu VALUES (452,'','团队Task删除','','task_del',3,3,'notice:task:remove',448,0,'',0,0,1,1,'2023-11-24 15:39:24.155','2023-11-24 15:39:24.155','2025-03-21 09:41:23.104');


-- Table structure for table sys_menu_api_rule
DROP TABLE IF EXISTS sys_menu_api_rule;
CREATE TABLE sys_menu_api_rule (
  sys_menu_id INTEGER NOT NULL,
  sys_api_id INTEGER NOT NULL,
  PRIMARY KEY (sys_menu_id, sys_api_id)
);

-- Dumping data for table sys_menu_api_rule
INSERT INTO sys_menu_api_rule VALUES (2,1);
INSERT INTO sys_menu_api_rule VALUES (3,2);
INSERT INTO sys_menu_api_rule VALUES (4,3);
INSERT INTO sys_menu_api_rule VALUES (5,4);
INSERT INTO sys_menu_api_rule VALUES (6,5);

-- Table structure for table sys_opera_log
DROP TABLE IF EXISTS sys_opera_log;
CREATE TABLE sys_opera_log (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  business_type TEXT,
  business_types TEXT,
  method TEXT,
  request_method TEXT,
  operator_type TEXT,
  oper_name TEXT,
  dept_name TEXT,
  oper_url TEXT,
  oper_ip TEXT,
  oper_location TEXT,
  oper_param TEXT,
  status INTEGER,
  oper_time DATETIME,
  json_result TEXT,
  remark TEXT,
  latency_time TEXT,
  user_agent TEXT,
  created_at DATETIME,
  updated_at DATETIME,
  create_by INTEGER,
  update_by INTEGER
);

-- Dumping data for table sys_opera_log
-- No data to insert

-- Table structure for table sys_sms
DROP TABLE IF EXISTS sys_sms;
CREATE TABLE sys_sms (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  phone TEXT,
  code TEXT,
  type TEXT,
  status INTEGER,
  use_status INTEGER,
  created_at INTEGER,
  updated_at INTEGER
);

-- Dumping data for table sys_sms
-- No data to insert

-- Table structure for table sys_user
DROP TABLE IF EXISTS sys_user;
CREATE TABLE sys_user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT,
  phone TEXT,
  email TEXT,
  password TEXT,
  nickname TEXT,
  name TEXT,
  avatar TEXT,
  bio TEXT,
  birthday DATE,
  gender INTEGER DEFAULT 2,
  team_id INTEGER,
  remark TEXT,
  lock_time DATETIME,
  status INTEGER,
  src_id TEXT,
  inviter INTEGER,
  invite_code TEXT,
  client_id TEXT,
  reg_ip TEXT,
  ip_location TEXT,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME
);

-- Dumping data for table sys_user
INSERT INTO sys_user VALUES (2,'tangtang','13800138001',NULL,'$2a$10$hpISaioe30s6d3LYPdFaw.23rm29Yw9a44t8hl9G.tCNYKYoA72iW','糖糖','小唐',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-11-18 20:35:24.140');
INSERT INTO sys_user VALUES (3,'zcm',NULL,NULL,'$2a$10$8ZbmdrNu22DMnRFDjfBugOwvK0fVOixMX.AAeLFbgOkjS9frXHjVW','小梅','小梅',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,1,NULL,'2023-11-12 19:59:39.849');
INSERT INTO sys_user VALUES (4,'liyanlei',NULL,NULL,'$2a$10$sPJOyDdh/J4OFVa27z4LCuz2xtu3KwhyNegkp27hL.G9D3xMMYoKW','阿雷','阿雷',NULL,NULL,NULL,1,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,1,NULL,'2023-11-12 20:03:10.441');

-- Table structure for table sys_user_login_log
DROP TABLE IF EXISTS sys_user_login_log;
CREATE TABLE sys_user_login_log (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uid INTEGER,
  method INTEGER,
  ip TEXT,
  region TEXT,
  client_id TEXT,
  ver TEXT,
  os TEXT,
  os_ver TEXT,
  team_id INTEGER,
  updated_at DATETIME
);

-- Dumping data for table sys_user_login_log
INSERT INTO sys_user_login_log VALUES (33,1,3,'127.0.0.1','','','','','',NULL,'2025-03-15 16:53:02');
INSERT INTO sys_user_login_log VALUES (34,1,3,'127.0.0.1','','','','','',NULL,'2025-03-15 16:53:15');
INSERT INTO sys_user_login_log VALUES (35,1,3,'127.0.0.1','','','','','',NULL,'2025-03-15 21:07:19');

-- Table structure for table sys_user_third_login
DROP TABLE IF EXISTS sys_user_third_login;
CREATE TABLE sys_user_third_login (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER,
  platform INTEGER,
  open_id TEXT,
  union_id TEXT,
  third_data TEXT,
  created_at INTEGER,
  updated_at INTEGER
);

-- Dumping data for table sys_user_third_login
-- No data to insert

-- Table structure for table team
DROP TABLE IF EXISTS team;
CREATE TABLE team (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  owner INTEGER,
  status INTEGER,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME
);

-- Dumping data for table team
INSERT INTO team VALUES (1,'销售管理系统',2,3,NULL,'2023-09-30 00:22:43','2025-03-21 10:40:46');
INSERT INTO team VALUES (4,'大金 Team',8,3,0,'2025-03-20 21:48:06','2025-03-21 10:40:42');
INSERT INTO team VALUES (6,'小琪 Team',9,3,0,'2025-03-22 10:10:59','2025-03-22 10:10:59');

-- Table structure for table team_dept
DROP TABLE IF EXISTS team_dept;
CREATE TABLE team_dept (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  parent_id INTEGER,
  dept_path TEXT,
  name TEXT,
  type INTEGER,
  principal TEXT,
  phone TEXT,
  email TEXT,
  sort INTEGER,
  status INTEGER,
  remark TEXT,
  team_id INTEGER,
  create_by INTEGER,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table team_dept
INSERT INTO team_dept VALUES (1,0,'/0/1/','销售一部',2,'糖糖','13800138000',NULL,1,1,NULL,1,NULL,2,NULL,'2023-10-26 20:57:48.321',NULL);
INSERT INTO team_dept VALUES (2,1,'/0/1/2/','糖糖组',2,'','',NULL,NULL,1,NULL,1,NULL,2,NULL,'2023-10-26 21:03:59.644',NULL);
INSERT INTO team_dept VALUES (9,10,'/0/10/9/','哈哈组',2,'哈哈','','',0,1,'1',1,2,1,'2023-10-26 20:50:45.790','2025-03-17 16:40:29.951',NULL);

-- Table structure for table team_member
DROP TABLE IF EXISTS team_member;
CREATE TABLE team_member (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  team_id INTEGER,
  user_id INTEGER,
  nickname TEXT,
  name TEXT,
  py TEXT,
  phone TEXT,
  dept_path TEXT,
  dept_id INTEGER,
  role_id INTEGER,
  post_id INTEGER,
  entry_time DATETIME,
  retire_time DATETIME,
  gender INTEGER DEFAULT 2,
  birthday DATE,
  status INTEGER,
  create_by INTEGER,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME
);

-- Dumping data for table team_member
INSERT INTO team_member VALUES (1,1,2,'糖糖','小唐','xiao-tang','','/0/1/2/',2,1,1,'2023-02-20 00:00:00.000',NULL,2,'1991-02-01',1,NULL,NULL,NULL,'2025-03-21 11:23:17.249');
INSERT INTO team_member VALUES (2,1,3,'梅梅','小梅','xiao-mei','13800138000','/0/1/',1,8,6,'2021-01-13 00:00:00.000',NULL,2,NULL,-1,NULL,1,NULL,'2025-03-21 11:23:10.910');
INSERT INTO team_member VALUES (12,1,4,'阿雷','阿雷','a-lei','','',0,0,6,NULL,NULL,1,NULL,1,0,0,'2025-03-20 15:18:10.954','2025-03-21 11:22:54.038');

-- Table structure for table team_role
DROP TABLE IF EXISTS team_role;
CREATE TABLE team_role (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  role_key TEXT,
  role_sort INTEGER,
  status INTEGER,
  team_id INTEGER,
  remark TEXT,
  create_by INTEGER,
  update_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table team_role
-- No data to insert

PRAGMA foreign_keys = ON;