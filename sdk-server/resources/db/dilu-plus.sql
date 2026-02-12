-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: dilu-plus
-- ------------------------------------------------------
-- Server version	8.0.41-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '昵称',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '姓名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '头像',
  `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '签名',
  `birthday` date DEFAULT NULL COMMENT '生日 格式 yyyy-MM-dd',
  `gender` tinyint(1) DEFAULT '2' COMMENT '性别 1男 2女 3未知',
  `role_id` int DEFAULT NULL COMMENT '角色id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门路径',
  `dept_id` int DEFAULT NULL COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `lock_time` datetime(3) DEFAULT NULL COMMENT '锁定结束时间',
  `status` tinyint DEFAULT NULL COMMENT '状态 1正常 ',
  `client_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '客户端id',
  `reg_ip` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '注册ip',
  `ip_location` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_user_update_by` (`update_by`) USING BTREE,
  KEY `admin_username_IDX` (`username`) USING BTREE,
  KEY `admin_phone_IDX` (`phone`) USING BTREE,
  KEY `admin_email_IDX` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'diluplus',NULL,NULL,'$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a',NULL,'aaa',NULL,NULL,NULL,2,1,NULL,1,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,'2025-03-18 14:37:38.723'),(2,'test','','','$2a$10$lHRg6AcfVs2xZ0kzJ5RohuBvrfFHGk6SikAuCgF4ZgMT1UKHqtTYe','','test','','',NULL,1,2,'/1/2/',2,'',NULL,1,'','','',0,'2025-03-19 11:45:21.606','2025-03-20 15:53:18.794'),(3,'test2','','','$2a$10$zUzRSIFSzWH5/IXnv5OfVuBWbdv70Kn7Uy5gwN5H4aqY4iuO.0sHe','','test2','','',NULL,0,2,NULL,0,'',NULL,1,'','','',0,'2025-03-19 11:47:52.958','2025-03-19 11:47:52.958');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_dept`
--

DROP TABLE IF EXISTS `admin_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_dept` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int unsigned DEFAULT NULL COMMENT '父id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门路径',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门名',
  `type` tinyint DEFAULT NULL COMMENT '类型 1分公司 2部门',
  `principal` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门领导',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `sort` tinyint DEFAULT NULL COMMENT '排序',
  `status` tinyint DEFAULT NULL COMMENT '状态 1正常 2关闭',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dept_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='部门';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_dept`
--

LOCK TABLES `admin_dept` WRITE;
/*!40000 ALTER TABLE `admin_dept` DISABLE KEYS */;
INSERT INTO `admin_dept` VALUES (1,0,'/1/','技术部',2,'','','',0,1,'',0,0,'2025-03-17 22:37:05.256','2025-03-17 22:37:05.256',NULL),(2,1,'/1/2/','后端',2,'','','',0,1,'',0,0,'2025-03-17 22:37:25.663','2025-03-17 22:37:25.663',NULL),(3,0,'/3/','产品部',2,'','','',0,1,'',0,0,'2025-03-17 22:37:50.104','2025-03-17 22:37:50.104',NULL);
/*!40000 ALTER TABLE `admin_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role`
--

DROP TABLE IF EXISTS `admin_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '角色名称',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '角色代码',
  `role_sort` int unsigned DEFAULT NULL COMMENT '排序',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `menu_ids` json DEFAULT NULL COMMENT '角色菜单',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_role_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role`
--

LOCK TABLES `admin_role` WRITE;
/*!40000 ALTER TABLE `admin_role` DISABLE KEYS */;
INSERT INTO `admin_role` VALUES (1,'超级管理员','',0,0,'',0,'[401, 426, 430, 429, 428, 427, 432, 436, 435, 434, 433, 402, 406, 405, 404, 403]',1,'2025-03-17 22:36:35.607','2025-03-19 11:21:23.426',NULL),(2,'操作员','',0,0,'',1,'[394, 2, 6, 5, 4, 3, 52, 55, 53, 54, 56, 17, 21, 20, 19, 18, 12, 13, 16, 15, 14, 57, 58, 59, 60, 61, 117, 121, 120, 119, 118, 126, 127, 130, 129, 128, 396, 400, 399, 398, 397]',0,'2025-03-19 11:21:45.712','2025-03-19 11:21:45.712',NULL);
/*!40000 ALTER TABLE `admin_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role_menu`
--

DROP TABLE IF EXISTS `admin_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role_menu` (
  `role_id` int unsigned NOT NULL COMMENT '主键',
  `menu_id` int unsigned NOT NULL COMMENT '主键',
  PRIMARY KEY (`role_id`,`menu_id`) USING BTREE,
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`) USING BTREE,
  CONSTRAINT `fk_sys_role_menu_sys_menu_copy` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role_copy` FOREIGN KEY (`role_id`) REFERENCES `team_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='角色菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role_menu`
--

LOCK TABLES `admin_role_menu` WRITE;
/*!40000 ALTER TABLE `admin_role_menu` DISABLE KEYS */;
INSERT INTO `admin_role_menu` VALUES (1,401),(1,402),(1,403),(1,404),(1,405),(1,406),(1,426),(1,427),(1,428),(1,429),(1,430),(1,432),(1,433),(1,434),(1,435),(1,436);
/*!40000 ALTER TABLE `admin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gen_columns`
--

DROP TABLE IF EXISTS `gen_columns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gen_columns` (
  `column_id` bigint NOT NULL AUTO_INCREMENT,
  `table_id` bigint DEFAULT NULL,
  `column_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `column_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `column_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `go_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `go_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `json_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_pk` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_increment` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_required` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_insert` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_edit` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_list` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_query` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `query_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `html_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  `list` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `pk` tinyint(1) DEFAULT NULL,
  `required` tinyint(1) DEFAULT NULL,
  `super_column` tinyint(1) DEFAULT NULL,
  `usable_column` tinyint(1) DEFAULT NULL,
  `increment` tinyint(1) DEFAULT NULL,
  `insert` tinyint(1) DEFAULT NULL,
  `edit` tinyint(1) DEFAULT NULL,
  `query` tinyint(1) DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `fk_table_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `fk_table_name_class` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `fk_table_name_package` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `fk_label_id` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `fk_label_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `create_by` mediumint DEFAULT NULL,
  `update_By` mediumint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1376 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='生成列';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gen_columns`
--

LOCK TABLES `gen_columns` WRITE;
/*!40000 ALTER TABLE `gen_columns` DISABLE KEYS */;
INSERT INTO `gen_columns` VALUES (1065,94,'column_id','','bigint','int','ColumnId','columnId','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.731','2025-03-15 16:16:37.731'),(1066,94,'table_id','','bigint','int','TableId','tableId','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.736','2025-03-15 16:16:37.736'),(1067,94,'column_name','','varchar(128)','string','ColumnName','columnName','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.743','2025-03-15 16:16:37.743'),(1068,94,'column_comment','','varchar(128)','string','ColumnComment','columnComment','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.749','2025-03-15 16:16:37.749'),(1069,94,'column_type','','varchar(128)','string','ColumnType','columnType','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.756','2025-03-15 16:16:37.756'),(1070,94,'go_type','','varchar(128)','string','GoType','goType','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.761','2025-03-15 16:16:37.761'),(1071,94,'go_field','','varchar(128)','string','GoField','goField','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.767','2025-03-15 16:16:37.767'),(1072,94,'json_field','','varchar(128)','string','JsonField','jsonField','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.775','2025-03-15 16:16:37.775'),(1073,94,'is_pk','','varchar(4)','string','IsPk','isPk','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.782','2025-03-15 16:16:37.782'),(1074,94,'is_increment','','varchar(4)','string','IsIncrement','isIncrement','0','','0','1','1','1','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.787','2025-03-15 16:16:37.787'),(1075,94,'is_required','','varchar(4)','string','IsRequired','isRequired','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.793','2025-03-15 16:16:37.793'),(1076,94,'is_insert','','varchar(4)','string','IsInsert','isInsert','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.800','2025-03-15 16:16:37.800'),(1077,94,'is_edit','','varchar(4)','string','IsEdit','isEdit','0','','0','1','1','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.805','2025-03-15 16:16:37.805'),(1078,94,'is_list','','varchar(4)','string','IsList','isList','0','','0','1','1','1','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.812','2025-03-15 16:16:37.812'),(1079,94,'is_query','','varchar(4)','string','IsQuery','isQuery','0','','0','1','1','1','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.818','2025-03-15 16:16:37.818'),(1080,94,'query_type','','varchar(128)','string','QueryType','queryType','0','','0','1','1','1','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.825','2025-03-15 16:16:37.825'),(1081,94,'html_type','','varchar(128)','string','HtmlType','htmlType','0','','0','1','1','1','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.833','2025-03-15 16:16:37.833'),(1082,94,'dict_type','','varchar(128)','string','DictType','dictType','0','','0','1','1','1','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.837','2025-03-15 16:16:37.837'),(1083,94,'sort','','bigint','int','Sort','sort','0','','0','1','1','1','','EQ','input','',19,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.843','2025-03-15 16:16:37.843'),(1084,94,'list','','varchar(1)','string','List','list','0','','0','1','1','1','','EQ','input','',20,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.848','2025-03-15 16:16:37.848'),(1085,94,'pk','','tinyint(1)','int8','Pk','pk','0','','0','1','1','1','','EQ','input','',21,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.855','2025-03-15 16:16:37.855'),(1086,94,'required','','tinyint(1)','int8','Required','required','0','','0','1','1','1','','EQ','input','',22,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.861','2025-03-15 16:16:37.861'),(1087,94,'super_column','','tinyint(1)','int8','SuperColumn','superColumn','0','','0','1','1','1','','EQ','input','',23,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.866','2025-03-15 16:16:37.866'),(1088,94,'usable_column','','tinyint(1)','int8','UsableColumn','usableColumn','0','','0','1','1','1','','EQ','input','',24,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.873','2025-03-15 16:16:37.873'),(1089,94,'increment','','tinyint(1)','int8','Increment','increment','0','','0','1','1','1','','EQ','input','',25,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.879','2025-03-15 16:16:37.879'),(1090,94,'insert','','tinyint(1)','int8','Insert','insert','0','','0','1','1','1','','EQ','input','',26,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.885','2025-03-15 16:16:37.885'),(1091,94,'edit','','tinyint(1)','int8','Edit','edit','0','','0','1','1','1','','EQ','input','',27,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.892','2025-03-15 16:16:37.892'),(1092,94,'query','','tinyint(1)','int8','Query','query','0','','0','1','1','1','','EQ','input','',28,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.899','2025-03-15 16:16:37.899'),(1093,94,'remark','','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',29,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.906','2025-03-15 16:16:37.906'),(1094,94,'fk_table_name','','longtext','string','FkTableName','fkTableName','0','','0','1','1','1','','EQ','input','',30,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.911','2025-03-15 16:16:37.911'),(1095,94,'fk_table_name_class','','longtext','string','FkTableNameClass','fkTableNameClass','0','','0','1','1','1','','EQ','input','',31,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.917','2025-03-15 16:16:37.917'),(1096,94,'fk_table_name_package','','longtext','string','FkTableNamePackage','fkTableNamePackage','0','','0','1','1','1','','EQ','input','',32,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.926','2025-03-15 16:16:37.926'),(1097,94,'fk_label_id','','longtext','string','FkLabelId','fkLabelId','0','','0','1','1','1','','EQ','input','',33,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.938','2025-03-15 16:16:37.938'),(1098,94,'fk_label_name','','varchar(255)','string','FkLabelName','fkLabelName','0','','0','1','1','1','','EQ','input','',34,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.948','2025-03-15 16:16:37.948'),(1099,94,'create_by','','mediumint','int','CreateBy','createBy','0','','0','1','','1','','EQ','input','',35,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.957','2025-03-15 16:16:37.957'),(1100,94,'update_By','','mediumint','int','UpdateBy','updateBy','0','','0','1','1','1','','EQ','input','',36,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.968','2025-03-15 16:16:37.968'),(1101,94,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',37,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.976','2025-03-15 16:16:37.976'),(1102,94,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',38,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:37.987','2025-03-15 16:16:37.987'),(1103,95,'table_id','','bigint','int','TableId','tableId','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.009','2025-03-15 16:16:38.009'),(1104,95,'db_name','','varchar(64)','string','DbName','dbName','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.018','2025-03-15 16:16:38.018'),(1105,95,'table_name','','varchar(128)','string','TableName','tableName','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.026','2025-03-15 16:16:38.026'),(1106,95,'table_comment','','varchar(128)','string','TableComment','tableComment','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.032','2025-03-15 16:16:38.032'),(1107,95,'class_name','','varchar(128)','string','ClassName','className','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.039','2025-03-15 16:16:38.039'),(1108,95,'tpl_category','','varchar(128)','string','TplCategory','tplCategory','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.045','2025-03-15 16:16:38.045'),(1109,95,'package_name','','varchar(128)','string','PackageName','packageName','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.053','2025-03-15 16:16:38.053'),(1110,95,'module_name','','varchar(128)','string','ModuleName','moduleName','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.058','2025-03-15 16:16:38.058'),(1111,95,'module_front_name','前端文件名','varchar(255)','string','ModuleFrontName','moduleFrontName','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.064','2025-03-15 16:16:38.064'),(1112,95,'business_name','','varchar(255)','string','BusinessName','businessName','0','','0','1','1','1','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.069','2025-03-15 16:16:38.069'),(1113,95,'function_name','','varchar(255)','string','FunctionName','functionName','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.075','2025-03-15 16:16:38.075'),(1114,95,'function_author','','varchar(255)','string','FunctionAuthor','functionAuthor','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.081','2025-03-15 16:16:38.081'),(1115,95,'pk_column','','varchar(255)','string','PkColumn','pkColumn','0','','0','1','1','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.087','2025-03-15 16:16:38.087'),(1116,95,'pk_go_field','','varchar(255)','string','PkGoField','pkGoField','0','','0','1','1','1','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.094','2025-03-15 16:16:38.094'),(1117,95,'pk_json_field','','varchar(255)','string','PkJsonField','pkJsonField','0','','0','1','1','1','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.101','2025-03-15 16:16:38.101'),(1118,95,'options','','varchar(255)','string','Options','options','0','','0','1','1','1','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.107','2025-03-15 16:16:38.107'),(1119,95,'tree_code','','varchar(255)','string','TreeCode','treeCode','0','','0','1','1','1','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.114','2025-03-15 16:16:38.114'),(1120,95,'tree_parent_code','','varchar(255)','string','TreeParentCode','treeParentCode','0','','0','1','1','1','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.119','2025-03-15 16:16:38.119'),(1121,95,'tree_name','','varchar(255)','string','TreeName','treeName','0','','0','1','1','1','','EQ','input','',19,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.124','2025-03-15 16:16:38.124'),(1122,95,'tree','','tinyint(1)','int8','Tree','tree','0','','0','1','1','1','','EQ','input','',20,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.131','2025-03-15 16:16:38.131'),(1123,95,'crud','','tinyint(1)','int8','Crud','crud','0','','0','1','1','1','','EQ','input','',21,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.137','2025-03-15 16:16:38.137'),(1124,95,'remark','','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',22,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.144','2025-03-15 16:16:38.144'),(1125,95,'is_data_scope','','tinyint','int8','IsDataScope','isDataScope','0','','0','1','1','1','','EQ','input','',23,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.152','2025-03-15 16:16:38.152'),(1126,95,'is_actions','','tinyint','int8','IsActions','isActions','0','','0','1','1','1','','EQ','input','',24,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.159','2025-03-15 16:16:38.159'),(1127,95,'is_auth','','tinyint','int8','IsAuth','isAuth','0','','0','1','1','1','','EQ','input','',25,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.164','2025-03-15 16:16:38.164'),(1128,95,'is_logical_delete','','varchar(1)','string','IsLogicalDelete','isLogicalDelete','0','','0','1','1','1','','EQ','input','',26,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.170','2025-03-15 16:16:38.170'),(1129,95,'logical_delete','','tinyint(1)','int8','LogicalDelete','logicalDelete','0','','0','1','1','1','','EQ','input','',27,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.176','2025-03-15 16:16:38.176'),(1130,95,'logical_delete_column','','varchar(128)','string','LogicalDeleteColumn','logicalDeleteColumn','0','','0','1','1','1','','EQ','input','',28,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.182','2025-03-15 16:16:38.182'),(1131,95,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',29,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.187','2025-03-15 16:16:38.187'),(1132,95,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',30,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.192','2025-03-15 16:16:38.192'),(1133,95,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',31,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.197','2025-03-15 16:16:38.197'),(1134,95,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',32,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.203','2025-03-15 16:16:38.203'),(1135,96,'id','主键编码','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.218','2025-03-15 16:16:38.218'),(1136,96,'title','标题','varchar(128)','string','Title','title','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.224','2025-03-15 16:16:38.224'),(1137,96,'method','请求类型','varchar(16)','string','Method','method','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.229','2025-03-15 16:16:38.229'),(1138,96,'path','请求地址','varchar(128)','string','Path','path','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.237','2025-03-15 16:16:38.237'),(1139,96,'perm_type','权限类型（1：无需认证 2:须token 3：须鉴权）','bigint','int','PermType','permType','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.244','2025-03-15 16:16:38.244'),(1140,96,'status','状态 3 DEF 2 OK 1 del','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.250','2025-03-15 16:16:38.250'),(1141,96,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.255','2025-03-15 16:16:38.255'),(1142,96,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.262','2025-03-15 16:16:38.262'),(1143,97,'id','主键编码','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.278','2025-03-15 16:16:38.278'),(1144,97,'name','名字','varchar(128)','string','Name','name','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.284','2025-03-15 16:16:38.284'),(1145,97,'key','key','varchar(128)','string','Key','key','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.291','2025-03-15 16:16:38.291'),(1146,97,'value','Value','json','string','Value','value','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.295','2025-03-15 16:16:38.295'),(1147,97,'type','Type','varchar(64)','string','Type','type','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.301','2025-03-15 16:16:38.301'),(1148,97,'remark','Remark','varchar(128)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.306','2025-03-15 16:16:38.306'),(1149,97,'status','Status','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.312','2025-03-15 16:16:38.312'),(1150,97,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.318','2025-03-15 16:16:38.318'),(1151,97,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.324','2025-03-15 16:16:38.324'),(1152,98,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.342','2025-03-15 16:16:38.342'),(1153,98,'parent_id','父id','int unsigned','uint','ParentId','parentId','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.347','2025-03-15 16:16:38.347'),(1154,98,'dept_path','部门路径','varchar(255)','string','DeptPath','deptPath','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.353','2025-03-15 16:16:38.353'),(1155,98,'name','部门名','varchar(128)','string','Name','name','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.359','2025-03-15 16:16:38.359'),(1156,98,'type','类型 1分公司 2部门','tinyint','int8','Type','type','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.363','2025-03-15 16:16:38.363'),(1157,98,'principal','部门领导','varchar(128)','string','Principal','principal','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.371','2025-03-15 16:16:38.371'),(1158,98,'phone','手机号','varchar(11)','string','Phone','phone','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.378','2025-03-15 16:16:38.378'),(1159,98,'email','邮箱','varchar(128)','string','Email','email','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.383','2025-03-15 16:16:38.383'),(1160,98,'sort','排序','tinyint','int8','Sort','sort','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.389','2025-03-15 16:16:38.389'),(1161,98,'status','状态 1正常 2关闭','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.397','2025-03-15 16:16:38.397'),(1162,98,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.403','2025-03-15 16:16:38.403'),(1163,98,'team_id','团队id','int','int','TeamId','teamId','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.408','2025-03-15 16:16:38.408'),(1164,98,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.415','2025-03-15 16:16:38.415'),(1165,98,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.421','2025-03-15 16:16:38.421'),(1166,98,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.427','2025-03-15 16:16:38.427'),(1167,98,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.434','2025-03-15 16:16:38.434'),(1168,98,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','1','','EQ','datetime','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.439','2025-03-15 16:16:38.439'),(1169,99,'id','主键编码','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.457','2025-03-15 16:16:38.457'),(1170,99,'email','邮箱地址','varchar(255)','string','Email','email','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.462','2025-03-15 16:16:38.462'),(1171,99,'code','验证码','varchar(6)','string','Code','code','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.468','2025-03-15 16:16:38.468'),(1172,99,'type','类型','varchar(6)','string','Type','type','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.473','2025-03-15 16:16:38.473'),(1173,99,'status','状态','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.479','2025-03-15 16:16:38.479'),(1174,99,'use_status','使用状态','tinyint','int8','UseStatus','useStatus','0','','0','1','1','1','1','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.486','2025-03-15 16:16:38.486'),(1175,99,'created_at','创建时间','int unsigned','uint','CreatedAt','createdAt','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.492','2025-03-15 16:16:38.492'),(1176,99,'updated_at','最后更新时间','int unsigned','uint','UpdatedAt','updatedAt','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.497','2025-03-15 16:16:38.497'),(1177,100,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.515','2025-03-15 16:16:38.515'),(1178,100,'job_name','名称','varchar(255)','string','JobName','jobName','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.521','2025-03-15 16:16:38.521'),(1179,100,'job_group','组','varchar(255)','string','JobGroup','jobGroup','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.525','2025-03-15 16:16:38.525'),(1180,100,'job_type','类型','tinyint','int8','JobType','jobType','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.532','2025-03-15 16:16:38.532'),(1181,100,'cron_expression','表达式','varchar(255)','string','CronExpression','cronExpression','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.537','2025-03-15 16:16:38.537'),(1182,100,'invoke_target','调用目标','varchar(255)','string','InvokeTarget','invokeTarget','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.542','2025-03-15 16:16:38.542'),(1183,100,'args','参数','varchar(255)','string','Args','args','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.547','2025-03-15 16:16:38.547'),(1184,100,'misfire_policy','策略','bigint','int','MisfirePolicy','misfirePolicy','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.552','2025-03-15 16:16:38.552'),(1185,100,'concurrent','并发','tinyint','int8','Concurrent','concurrent','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.559','2025-03-15 16:16:38.559'),(1186,100,'status','状态','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.564','2025-03-15 16:16:38.564'),(1187,100,'entry_id','任务id','smallint','int16','EntryId','entryId','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.570','2025-03-15 16:16:38.570'),(1188,100,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.576','2025-03-15 16:16:38.576'),(1189,100,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.583','2025-03-15 16:16:38.583'),(1190,100,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','1','','EQ','datetime','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.589','2025-03-15 16:16:38.589'),(1191,100,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.595','2025-03-15 16:16:38.595'),(1192,100,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.601','2025-03-15 16:16:38.601'),(1193,101,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.615','2025-03-15 16:16:38.615'),(1194,101,'team_id','团队id','int','int','TeamId','teamId','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.623','2025-03-15 16:16:38.623'),(1195,101,'user_id','用户id','int unsigned','uint','UserId','userId','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.627','2025-03-15 16:16:38.627'),(1196,101,'nickname','昵称','varchar(128)','string','Nickname','nickname','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.633','2025-03-15 16:16:38.633'),(1197,101,'name','姓名','varchar(64)','string','Name','name','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.638','2025-03-15 16:16:38.638'),(1198,101,'py','姓名拼音','varchar(32)','string','Py','py','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.645','2025-03-15 16:16:38.645'),(1199,101,'phone','电话','varchar(11)','string','Phone','phone','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.650','2025-03-15 16:16:38.650'),(1200,101,'dept_path','部门路径','varchar(255)','string','DeptPath','deptPath','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.656','2025-03-15 16:16:38.656'),(1201,101,'dept_id','部门id','int unsigned','uint','DeptId','deptId','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.662','2025-03-15 16:16:38.662'),(1202,101,'roles','角色id','varchar(255)','string','Roles','roles','0','','0','1','1','1','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.668','2025-03-15 16:16:38.668'),(1203,101,'post_id','职位 1系统超管 2 团队拥有者 4主管 8副主管 16员工','tinyint unsigned','int8','PostId','postId','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.675','2025-03-15 16:16:38.675'),(1204,101,'entry_time','入职时间','datetime(3)','time.Time','EntryTime','entryTime','0','','0','1','','1','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.682','2025-03-15 16:16:38.682'),(1205,101,'retire_time','离职时间','datetime(3)','time.Time','RetireTime','retireTime','0','','0','1','','1','','EQ','datetime','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.687','2025-03-15 16:16:38.687'),(1206,101,'gender','性别 1男 2女 3未知','tinyint(1)','int8','Gender','gender','0','','0','1','1','1','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.693','2025-03-15 16:16:38.693'),(1207,101,'birthday','生日 格式 yyyy-MM-dd','date','time.Time','Birthday','birthday','0','','0','1','','1','','EQ','datetime','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.698','2025-03-15 16:16:38.698'),(1208,101,'status','状态 1正常 ','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.704','2025-03-15 16:16:38.704'),(1209,101,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.709','2025-03-15 16:16:38.709'),(1210,101,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.716','2025-03-15 16:16:38.716'),(1211,101,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',19,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.722','2025-03-15 16:16:38.722'),(1212,101,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',20,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.727','2025-03-15 16:16:38.727'),(1213,102,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.743','2025-03-15 16:16:38.743'),(1214,102,'menu_name','菜单名','varchar(128)','string','MenuName','menuName','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.748','2025-03-15 16:16:38.748'),(1215,102,'title','显示名称','varchar(128)','string','Title','title','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.755','2025-03-15 16:16:38.755'),(1216,102,'icon','图标','varchar(128)','string','Icon','icon','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.761','2025-03-15 16:16:38.761'),(1217,102,'path','路径','varchar(128)','string','Path','path','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.768','2025-03-15 16:16:38.768'),(1218,102,'platform_type','平台类型 1 平台管理 2团队管理','bigint','int','PlatformType','platformType','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.777','2025-03-15 16:16:38.777'),(1219,102,'menu_type','菜单类型 1 分类 2菜单 3方法按钮','tinyint','int8','MenuType','menuType','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.783','2025-03-15 16:16:38.783'),(1220,102,'permission','权限','varchar(255)','string','Permission','permission','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.789','2025-03-15 16:16:38.789'),(1221,102,'parent_id','菜单父id','int unsigned','uint','ParentId','parentId','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.795','2025-03-15 16:16:38.795'),(1222,102,'no_cache','是否缓存','tinyint(1)','int8','NoCache','noCache','0','','0','1','1','1','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.802','2025-03-15 16:16:38.802'),(1223,102,'component','前端组件路径','varchar(255)','string','Component','component','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.808','2025-03-15 16:16:38.808'),(1224,102,'sort','排序倒叙','tinyint','int8','Sort','sort','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.814','2025-03-15 16:16:38.814'),(1225,102,'hidden','是否隐藏','tinyint(1)','int8','Hidden','hidden','0','','0','1','1','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.821','2025-03-15 16:16:38.821'),(1226,102,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.826','2025-03-15 16:16:38.826'),(1227,102,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.833','2025-03-15 16:16:38.833'),(1228,102,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.838','2025-03-15 16:16:38.838'),(1229,102,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.844','2025-03-15 16:16:38.844'),(1230,102,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','1','','EQ','datetime','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.850','2025-03-15 16:16:38.850'),(1231,103,'sys_menu_id','主键','int unsigned','uint','SysMenuId','sysMenuId','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.867','2025-03-15 16:16:38.867'),(1232,103,'sys_api_id','主键编码','int unsigned','uint','SysApiId','sysApiId','1','','1','1','1','1','','EQ','input','',2,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.873','2025-03-15 16:16:38.873'),(1233,104,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.934','2025-03-15 16:16:38.934'),(1234,104,'title','操作模块','varchar(255)','string','Title','title','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.944','2025-03-15 16:16:38.944'),(1235,104,'business_type','操作类型','varchar(128)','string','BusinessType','businessType','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.949','2025-03-15 16:16:38.949'),(1236,104,'business_types','BusinessTypes','varchar(128)','string','BusinessTypes','businessTypes','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.954','2025-03-15 16:16:38.954'),(1237,104,'method','函数','varchar(128)','string','Method','method','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.962','2025-03-15 16:16:38.962'),(1238,104,'request_method','请求方式 GET POST PUT DELETE','varchar(128)','string','RequestMethod','requestMethod','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.967','2025-03-15 16:16:38.967'),(1239,104,'operator_type','操作类型','varchar(128)','string','OperatorType','operatorType','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.972','2025-03-15 16:16:38.972'),(1240,104,'oper_name','操作者','varchar(128)','string','OperName','operName','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.978','2025-03-15 16:16:38.978'),(1241,104,'dept_name','部门名称','varchar(128)','string','DeptName','deptName','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.984','2025-03-15 16:16:38.984'),(1242,104,'oper_url','访问地址','varchar(255)','string','OperUrl','operUrl','0','','0','1','1','1','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.989','2025-03-15 16:16:38.989'),(1243,104,'oper_ip','客户端ip','varchar(128)','string','OperIp','operIp','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:38.997','2025-03-15 16:16:38.997'),(1244,104,'oper_location','访问位置','varchar(128)','string','OperLocation','operLocation','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.003','2025-03-15 16:16:39.003'),(1245,104,'oper_param','请求参数','longtext','string','OperParam','operParam','0','','0','1','1','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.008','2025-03-15 16:16:39.008'),(1246,104,'status','操作状态 1:成功 2:失败','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.015','2025-03-15 16:16:39.015'),(1247,104,'oper_time','操作时间','datetime(3)','time.Time','OperTime','operTime','0','','0','1','','1','','EQ','datetime','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.022','2025-03-15 16:16:39.022'),(1248,104,'json_result','返回数据','varchar(255)','string','JsonResult','jsonResult','0','','0','1','1','1','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.026','2025-03-15 16:16:39.026'),(1249,104,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.032','2025-03-15 16:16:39.032'),(1250,104,'latency_time','耗时','varchar(128)','string','LatencyTime','latencyTime','0','','0','1','1','1','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.037','2025-03-15 16:16:39.037'),(1251,104,'user_agent','ua','varchar(255)','string','UserAgent','userAgent','0','','0','1','1','1','','EQ','input','',19,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.043','2025-03-15 16:16:39.043'),(1252,104,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',20,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.049','2025-03-15 16:16:39.049'),(1253,104,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',21,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.056','2025-03-15 16:16:39.056'),(1254,104,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',22,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.062','2025-03-15 16:16:39.062'),(1255,104,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',23,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.068','2025-03-15 16:16:39.068'),(1256,105,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.084','2025-03-15 16:16:39.084'),(1257,105,'name','角色名称','varchar(128)','string','Name','name','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.089','2025-03-15 16:16:39.089'),(1258,105,'role_key','角色代码','varchar(128)','string','RoleKey','roleKey','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.095','2025-03-15 16:16:39.095'),(1259,105,'role_sort','排序','int unsigned','uint','RoleSort','roleSort','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.103','2025-03-15 16:16:39.103'),(1260,105,'status','状态','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.109','2025-03-15 16:16:39.109'),(1261,105,'team_id','团队','tinyint(1)','int8','TeamId','teamId','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.117','2025-03-15 16:16:39.117'),(1262,105,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.123','2025-03-15 16:16:39.123'),(1263,105,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.128','2025-03-15 16:16:39.128'),(1264,105,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.135','2025-03-15 16:16:39.135'),(1265,105,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.141','2025-03-15 16:16:39.141'),(1266,105,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.146','2025-03-15 16:16:39.146'),(1267,105,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','1','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.152','2025-03-15 16:16:39.152'),(1268,106,'role_id','主键','int unsigned','uint','RoleId','roleId','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.169','2025-03-15 16:16:39.169'),(1269,106,'menu_id','主键','int unsigned','uint','MenuId','menuId','1','','1','1','1','1','','EQ','input','',2,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.175','2025-03-15 16:16:39.175'),(1270,107,'id','主键编码','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.192','2025-03-15 16:16:39.192'),(1271,107,'phone','手机号','varchar(16)','string','Phone','phone','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.197','2025-03-15 16:16:39.197'),(1272,107,'code','验证码','varchar(6)','string','Code','code','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.203','2025-03-15 16:16:39.203'),(1273,107,'type','类型','varchar(6)','string','Type','type','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.208','2025-03-15 16:16:39.208'),(1274,107,'status','状态','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.214','2025-03-15 16:16:39.214'),(1275,107,'use_status','使用状态','tinyint','int8','UseStatus','useStatus','0','','0','1','1','1','1','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.221','2025-03-15 16:16:39.221'),(1276,107,'created_at','创建时间','int unsigned','uint','CreatedAt','createdAt','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.226','2025-03-15 16:16:39.226'),(1277,107,'updated_at','最后更新时间','int unsigned','uint','UpdatedAt','updatedAt','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.232','2025-03-15 16:16:39.232'),(1278,108,'id','主键','int','int','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.246','2025-03-15 16:16:39.246'),(1279,108,'name','团队名','varchar(32)','string','Name','name','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.252','2025-03-15 16:16:39.252'),(1280,108,'owner','团队拥有者','int unsigned','uint','Owner','owner','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.257','2025-03-15 16:16:39.257'),(1281,108,'status','状态','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.263','2025-03-15 16:16:39.263'),(1282,108,'created_at','创建时间','datetime','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.270','2025-03-15 16:16:39.270'),(1283,108,'updated_at','更新时间','datetime','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.275','2025-03-15 16:16:39.275'),(1284,109,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.289','2025-03-15 16:16:39.289'),(1285,109,'username','用户名','varchar(32)','string','Username','username','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.294','2025-03-15 16:16:39.294'),(1286,109,'phone','手机号','varchar(11)','string','Phone','phone','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.301','2025-03-15 16:16:39.301'),(1287,109,'email','邮箱','varchar(128)','string','Email','email','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.306','2025-03-15 16:16:39.306'),(1288,109,'password','密码','varchar(128)','string','Password','password','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.312','2025-03-15 16:16:39.312'),(1289,109,'nickname','昵称','varchar(128)','string','Nickname','nickname','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.318','2025-03-15 16:16:39.318'),(1290,109,'name','姓名','varchar(64)','string','Name','name','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.324','2025-03-15 16:16:39.324'),(1291,109,'avatar','头像','varchar(255)','string','Avatar','avatar','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.331','2025-03-15 16:16:39.331'),(1292,109,'bio','签名','varchar(255)','string','Bio','bio','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.336','2025-03-15 16:16:39.336'),(1293,109,'birthday','生日 格式 yyyy-MM-dd','date','time.Time','Birthday','birthday','0','','0','1','','1','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.341','2025-03-15 16:16:39.341'),(1294,109,'gender','性别 1男 2女 3未知','tinyint(1)','int8','Gender','gender','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.346','2025-03-15 16:16:39.346'),(1295,109,'platform_role_id','平台角色ID 大于0为平台账户,0为团队账户','mediumint','int','PlatformRoleId','platformRoleId','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.350','2025-03-15 16:16:39.350'),(1296,109,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.356','2025-03-15 16:16:39.356'),(1297,109,'lock_time','锁定结束时间','datetime(3)','time.Time','LockTime','lockTime','0','','0','1','','1','','EQ','datetime','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.362','2025-03-15 16:16:39.362'),(1298,109,'status','状态 1正常 ','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.368','2025-03-15 16:16:39.368'),(1299,109,'src_id','来源','varchar(32)','string','SrcId','srcId','0','','0','1','1','1','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.374','2025-03-15 16:16:39.374'),(1300,109,'inviter','邀请人','int unsigned','uint','Inviter','inviter','0','','0','1','1','1','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.380','2025-03-15 16:16:39.380'),(1301,109,'invite_code','邀请码','varchar(32)','string','InviteCode','inviteCode','0','','0','1','1','1','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.386','2025-03-15 16:16:39.386'),(1302,109,'client_id','客户端id','varchar(64)','string','ClientId','clientId','0','','0','1','1','1','','EQ','input','',19,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.392','2025-03-15 16:16:39.392'),(1303,109,'reg_ip','注册ip','varchar(42)','string','RegIp','regIp','0','','0','1','1','1','','EQ','input','',20,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.397','2025-03-15 16:16:39.397'),(1304,109,'ip_location','','varchar(100)','string','IpLocation','ipLocation','0','','0','1','1','1','','EQ','input','',21,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.404','2025-03-15 16:16:39.404'),(1305,109,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',22,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.411','2025-03-15 16:16:39.411'),(1306,109,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',23,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.416','2025-03-15 16:16:39.416'),(1307,109,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',24,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.422','2025-03-15 16:16:39.422'),(1308,110,'id','id','bigint unsigned','int','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.437','2025-03-15 16:16:39.437'),(1309,110,'uid','用户id','int unsigned','uint','Uid','uid','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.443','2025-03-15 16:16:39.443'),(1310,110,'method','登录方式','tinyint','int8','Method','method','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.447','2025-03-15 16:16:39.447'),(1311,110,'ip','登录ip','varchar(42)','string','Ip','ip','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.454','2025-03-15 16:16:39.454'),(1312,110,'region','地域','varchar(100)','string','Region','region','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.460','2025-03-15 16:16:39.460'),(1313,110,'client_id','客户端','varchar(64)','string','ClientId','clientId','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.465','2025-03-15 16:16:39.465'),(1314,110,'client_ver','操作系统','varchar(32)','string','ClientVer','clientVer','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.471','2025-03-15 16:16:39.471'),(1315,110,'os','操作系统','varchar(32)','string','Os','os','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.476','2025-03-15 16:16:39.476'),(1316,110,'os_ver','操作系统版本','varchar(64)','string','OsVer','osVer','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.482','2025-03-15 16:16:39.482'),(1317,110,'updated_at','更新时间','datetime','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:39.487','2025-03-15 16:16:39.487'),(1318,111,'id','主键编码','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.754','2025-03-15 16:16:53.754'),(1319,111,'user_id','用户id','int unsigned','uint','UserId','userId','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.758','2025-03-15 16:16:53.758'),(1320,111,'platform','平台 1 微信 2 钉钉','tinyint unsigned','int8','Platform','platform','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.763','2025-03-15 16:16:53.763'),(1321,111,'open_id','第三方open_id','varchar(128)','string','OpenId','openId','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.770','2025-03-15 16:16:53.770'),(1322,111,'union_id','第三方union_id','varchar(128)','string','UnionId','unionId','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.778','2025-03-15 16:16:53.778'),(1323,111,'third_data','第三方返回数据','json','string','ThirdData','thirdData','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.782','2025-03-15 16:16:53.782'),(1324,111,'created_at','创建时间','int unsigned','uint','CreatedAt','createdAt','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.788','2025-03-15 16:16:53.788'),(1325,111,'updated_at','最后更新时间','int unsigned','uint','UpdatedAt','updatedAt','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-15 16:16:53.796','2025-03-15 16:16:53.796'),(1326,112,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.168','2025-03-17 09:22:27.168'),(1327,112,'username','用户名','varchar(32)','string','Username','username','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.179','2025-03-17 09:22:27.179'),(1328,112,'phone','手机号','varchar(11)','string','Phone','phone','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.190','2025-03-17 09:22:27.190'),(1329,112,'email','邮箱','varchar(128)','string','Email','email','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.202','2025-03-17 09:22:27.202'),(1330,112,'password','密码','varchar(128)','string','Password','password','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.213','2025-03-17 09:22:27.213'),(1331,112,'nickname','昵称','varchar(128)','string','Nickname','nickname','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.224','2025-03-17 09:22:27.224'),(1332,112,'name','姓名','varchar(64)','string','Name','name','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.240','2025-03-17 09:22:27.240'),(1333,112,'avatar','头像','varchar(255)','string','Avatar','avatar','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.250','2025-03-17 09:22:27.250'),(1334,112,'bio','签名','varchar(255)','string','Bio','bio','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.261','2025-03-17 09:22:27.261'),(1335,112,'birthday','生日 格式 yyyy-MM-dd','date','time.Time','Birthday','birthday','0','','0','1','','1','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.273','2025-03-17 09:22:27.273'),(1336,112,'gender','性别 1男 2女 3未知','tinyint(1)','int8','Gender','gender','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.286','2025-03-17 09:22:27.286'),(1337,112,'role_id','团队id','int','int','RoleId','roleId','0','','0','1','1','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.297','2025-03-17 09:22:27.297'),(1338,112,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.311','2025-03-17 09:22:27.311'),(1339,112,'lock_time','锁定结束时间','datetime(3)','time.Time','LockTime','lockTime','0','','0','1','','1','','EQ','datetime','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.323','2025-03-17 09:22:27.323'),(1340,112,'status','状态 1正常 ','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.333','2025-03-17 09:22:27.333'),(1341,112,'client_id','客户端id','varchar(64)','string','ClientId','clientId','0','','0','1','1','1','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.346','2025-03-17 09:22:27.346'),(1342,112,'reg_ip','注册ip','varchar(42)','string','RegIp','regIp','0','','0','1','1','1','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.358','2025-03-17 09:22:27.358'),(1343,112,'ip_location','','varchar(100)','string','IpLocation','ipLocation','0','','0','1','1','1','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.371','2025-03-17 09:22:27.371'),(1344,112,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',19,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.383','2025-03-17 09:22:27.383'),(1345,112,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',20,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.395','2025-03-17 09:22:27.395'),(1346,112,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',21,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.407','2025-03-17 09:22:27.407'),(1347,113,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.436','2025-03-17 09:22:27.436'),(1348,113,'parent_id','父id','int unsigned','uint','ParentId','parentId','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.448','2025-03-17 09:22:27.448'),(1349,113,'dept_path','部门路径','varchar(255)','string','DeptPath','deptPath','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.461','2025-03-17 09:22:27.461'),(1350,113,'name','部门名','varchar(128)','string','Name','name','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.473','2025-03-17 09:22:27.473'),(1351,113,'type','类型 1分公司 2部门','tinyint','int8','Type','type','0','','0','1','1','1','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.485','2025-03-17 09:22:27.485'),(1352,113,'principal','部门领导','varchar(128)','string','Principal','principal','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.498','2025-03-17 09:22:27.498'),(1353,113,'phone','手机号','varchar(11)','string','Phone','phone','0','','0','1','1','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.512','2025-03-17 09:22:27.512'),(1354,113,'email','邮箱','varchar(128)','string','Email','email','0','','0','1','1','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.523','2025-03-17 09:22:27.523'),(1355,113,'sort','排序','tinyint','int8','Sort','sort','0','','0','1','1','1','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.535','2025-03-17 09:22:27.535'),(1356,113,'status','状态 1正常 2关闭','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.545','2025-03-17 09:22:27.545'),(1357,113,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.557','2025-03-17 09:22:27.557'),(1358,113,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.569','2025-03-17 09:22:27.569'),(1359,113,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.583','2025-03-17 09:22:27.583'),(1360,113,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',14,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.601','2025-03-17 09:22:27.601'),(1361,113,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',15,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.613','2025-03-17 09:22:27.613'),(1362,113,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','1','','EQ','datetime','',16,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.624','2025-03-17 09:22:27.624'),(1363,114,'id','主键','int unsigned','uint','Id','id','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.653','2025-03-17 09:22:27.653'),(1364,114,'name','角色名称','varchar(128)','string','Name','name','0','','0','1','1','1','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.668','2025-03-17 09:22:27.668'),(1365,114,'role_key','角色代码','varchar(128)','string','RoleKey','roleKey','0','','0','1','1','1','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.686','2025-03-17 09:22:27.686'),(1366,114,'role_sort','排序','int unsigned','uint','RoleSort','roleSort','0','','0','1','1','1','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.696','2025-03-17 09:22:27.696'),(1367,114,'status','状态','tinyint','int8','Status','status','0','','0','1','1','1','1','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.711','2025-03-17 09:22:27.711'),(1368,114,'remark','备注','varchar(255)','string','Remark','remark','0','','0','1','1','1','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.724','2025-03-17 09:22:27.724'),(1369,114,'create_by','创建者','int unsigned','uint','CreateBy','createBy','0','','0','1','','1','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.737','2025-03-17 09:22:27.737'),(1370,114,'update_by','更新者','int unsigned','uint','UpdateBy','updateBy','0','','0','1','','1','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.749','2025-03-17 09:22:27.749'),(1371,114,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','1','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.760','2025-03-17 09:22:27.760'),(1372,114,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','1','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.771','2025-03-17 09:22:27.771'),(1373,114,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','1','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.783','2025-03-17 09:22:27.783'),(1374,115,'role_id','主键','int unsigned','uint','RoleId','roleId','1','','1','1','1','1','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.815','2025-03-17 09:22:27.815'),(1375,115,'menu_id','主键','int unsigned','uint','MenuId','menuId','1','','1','1','1','1','','EQ','input','',2,'',1,1,0,0,0,1,0,0,'','','','','','',0,0,'2025-03-17 09:22:27.826','2025-03-17 09:22:27.826');
/*!40000 ALTER TABLE `gen_columns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gen_tables`
--

DROP TABLE IF EXISTS `gen_tables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gen_tables` (
  `table_id` bigint NOT NULL AUTO_INCREMENT,
  `db_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `table_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `table_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `class_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tpl_category` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `package_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `module_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `module_front_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '前端文件名',
  `business_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `function_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `function_author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `pk_column` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `pk_go_field` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `pk_json_field` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `options` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tree_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tree_parent_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tree_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tree` tinyint(1) DEFAULT '0',
  `crud` tinyint(1) DEFAULT '1',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_data_scope` tinyint DEFAULT NULL,
  `is_actions` tinyint DEFAULT NULL,
  `is_auth` tinyint DEFAULT NULL,
  `is_logical_delete` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `logical_delete` tinyint(1) DEFAULT NULL,
  `logical_delete_column` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`table_id`) USING BTREE,
  KEY `idx_gen_tables_create_by` (`create_by`) USING BTREE,
  KEY `idx_gen_tables_update_by` (`update_by`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='生成表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gen_tables`
--

LOCK TABLES `gen_tables` WRITE;
/*!40000 ALTER TABLE `gen_tables` DISABLE KEYS */;
INSERT INTO `gen_tables` VALUES (94,'dilu-plus','gen_columns','生成列','GenColumns','crud','sys','gen-columns','','genColumns','生成列','baowk','column_id','ColumnId','columnId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:37.724','2025-03-15 16:16:37.724',0,0),(95,'dilu-plus','gen_tables','生成表','GenTables','crud','sys','gen-tables','','genTables','生成表','baowk','table_id','TableId','tableId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.003','2025-03-15 16:16:38.003',0,0),(96,'dilu-plus','sys_api','接口','SysApi','crud','sys','sys-api','','sysApi','接口','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.213','2025-03-15 16:16:38.213',0,0),(97,'dilu-plus','sys_cfg','系统配置','SysCfg','crud','sys','sys-cfg','','sysCfg','系统配置','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.272','2025-03-15 16:16:38.272',0,0),(98,'dilu-plus','sys_dept','部门','SysDept','crud','sys','sys-dept','','sysDept','部门','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.335','2025-03-15 16:16:38.335',0,0),(99,'dilu-plus','sys_email','邮件','SysEmail','crud','sys','sys-email','','sysEmail','邮件','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.450','2025-03-15 16:16:38.450',0,0),(100,'dilu-plus','sys_job','定时任务','SysJob','crud','sys','sys-job','','sysJob','定时任务','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.508','2025-03-15 16:16:38.508',0,0),(101,'dilu-plus','sys_member','团队成员','SysMember','crud','sys','sys-member','','sysMember','团队成员','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.611','2025-03-15 16:16:38.611',0,0),(102,'dilu-plus','sys_menu','菜单','SysMenu','crud','sys','sys-menu','','sysMenu','菜单','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.737','2025-03-15 16:16:38.737',0,0),(103,'dilu-plus','sys_menu_api_rule','菜单关联表','SysMenuApiRule','crud','sys','sys-menu-api-rule','','sysMenuApiRule','菜单关联表','baowk','sys_api_id','SysApiId','sysApiId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.860','2025-03-15 16:16:38.860',0,0),(104,'dilu-plus','sys_opera_log','操作日志','SysOperaLog','crud','sys','sys-opera-log','','sysOperaLog','操作日志','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:38.923','2025-03-15 16:16:38.923',0,0),(105,'dilu-plus','sys_role','角色','SysRole','crud','sys','sys-role','','sysRole','角色','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:39.079','2025-03-15 16:16:39.079',0,0),(106,'dilu-plus','sys_role_menu','角色菜单','SysRoleMenu','crud','sys','sys-role-menu','','sysRoleMenu','角色菜单','baowk','menu_id','MenuId','menuId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:39.163','2025-03-15 16:16:39.163',0,0),(107,'dilu-plus','sys_sms','短信','SysSms','crud','sys','sys-sms','','sysSms','短信','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:39.186','2025-03-15 16:16:39.186',0,0),(108,'dilu-plus','sys_team','团队','SysTeam','crud','sys','sys-team','','sysTeam','团队','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:39.240','2025-03-15 16:16:39.240',0,0),(109,'dilu-plus','sys_user','用户','SysUser','crud','sys','sys-user','','sysUser','用户','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:39.284','2025-03-15 16:16:39.284',0,0),(110,'dilu-plus','sys_user_login_log','用户登录记录','SysUserLoginLog','crud','sys','sys-user-login-log','','sysUserLoginLog','用户登录记录','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:39.431','2025-03-15 16:16:39.431',0,0),(111,'dilu-plus','sys_user_third_login','三方登录','SysUserThirdLogin','crud','sys','sys-user-third-login','','sysUserThirdLogin','三方登录','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-15 16:16:53.747','2025-03-15 16:16:53.747',0,0),(112,'dilu-plus','admin','用户','Admin','crud','sys','admin','','admin','用户','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-17 09:22:27.156','2025-03-17 09:22:27.156',0,0),(113,'dilu-plus','admin_dept','部门','AdminDept','crud','sys','admin-dept','','adminDept','部门','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-17 09:22:27.425','2025-03-17 09:22:27.425',0,0),(114,'dilu-plus','admin_role','角色','AdminRole','crud','sys','admin-role','','adminRole','角色','baowk','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-17 09:22:27.641','2025-03-17 09:22:27.641',0,0),(115,'dilu-plus','admin_role_menu','角色菜单','AdminRoleMenu','crud','sys','admin-role-menu','','adminRoleMenu','角色菜单','baowk','menu_id','MenuId','menuId','','','','',0,1,'',1,2,1,'1',1,'is_del','2025-03-17 09:22:27.802','2025-03-17 09:22:27.802',0,0);
/*!40000 ALTER TABLE `gen_tables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_api`
--

DROP TABLE IF EXISTS `sys_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '标题',
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '请求类型',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '请求地址',
  `perm_type` bigint DEFAULT NULL COMMENT '权限类型（1：无需认证 2:须token 3：须鉴权）',
  `status` tinyint DEFAULT NULL COMMENT '状态 3 DEF 2 OK 1 del',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_method_path` (`method`,`path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=377 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='接口';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_api`
--

LOCK TABLES `sys_api` WRITE;
/*!40000 ALTER TABLE `sys_api` DISABLE KEYS */;
INSERT INTO `sys_api` VALUES (1,'分页获取用户','POST','/api/sys/user/page',3,3,0,'2023-09-26 13:46:59.488'),(2,'根据id获取用户','POST','/api/sys/user/get',3,3,0,'2023-09-26 13:46:59.518'),(3,'创建用户','POST','/api/sys/user/create',3,3,0,'2023-09-26 13:46:59.532'),(4,'修改用户','POST','/api/sys/user/update',3,3,0,'2023-09-26 13:46:59.539'),(5,'删除用户','POST','/api/sys/user/del',3,3,0,'2023-09-26 13:46:59.550'),(6,'分页获取菜单','POST','/api/sys/menu/all',3,3,0,'2023-09-26 13:47:37.020'),(7,'根据id获取菜单','POST','/api/sys/menu/get',3,3,0,'2023-09-26 13:47:37.038'),(8,'创建菜单','POST','/api/sys/menu/create',3,3,0,'2023-09-26 13:47:37.064'),(9,'修改菜单','POST','/api/sys/menu/update',3,3,0,'2023-09-26 13:47:37.073'),(10,'删除菜单','POST','/api/sys/menu/del',3,3,0,'2023-09-26 13:47:37.082'),(11,'分页获取角色','POST','/api/sys/role/page',3,3,0,'2023-09-26 13:47:39.685'),(12,'根据id获取角色','POST','/api/sys/role/get',3,3,0,'2023-09-26 13:47:39.701'),(13,'创建角色','POST','/api/sys/role/create',3,3,0,'2023-09-26 13:47:39.734'),(14,'修改角色','POST','/api/sys/role/update',3,3,0,'2023-09-26 13:47:39.762'),(15,'删除角色','POST','/api/sys/role/del',3,3,0,'2023-09-26 13:47:39.773'),(16,'分页获取部门','POST','/api/sys/dept/all',3,3,0,'2023-09-26 13:47:42.136'),(17,'根据id获取部门','POST','/api/sys/dept/get',3,3,0,'2023-09-26 13:47:42.152'),(18,'创建部门','POST','/api/sys/dept/create',3,3,0,'2023-09-26 13:47:42.172'),(19,'修改部门','POST','/api/sys/dept/update',3,3,0,'2023-09-26 13:47:42.186'),(20,'删除部门','POST','/api/sys/dept/del',3,3,0,'2023-09-26 13:47:42.197'),(26,'分页获取成员','POST','/api/sys/member/page',3,3,0,'2023-09-26 14:09:40.443'),(27,'根据id获取成员','POST','/api/sys/member/get',3,3,0,'2023-09-26 14:09:40.459'),(28,'创建成员','POST','/api/sys/member/create',3,3,0,'2023-09-26 14:09:40.483'),(29,'修改成员','POST','/api/sys/member/update',3,3,0,'2023-09-26 14:09:40.493'),(30,'删除成员','POST','/api/sys/member/del',3,3,0,'2023-09-26 14:09:40.502'),(31,'分页获取团队','POST','/api/sys/team/page',3,3,0,'2023-09-29 08:40:00.840'),(32,'根据id获取团队','POST','/api/sys/team/get',3,3,0,'2023-09-29 08:40:00.855'),(33,'创建团队','POST','/api/sys/team/create',3,3,0,'2023-09-29 08:40:00.867'),(34,'修改团队','POST','/api/sys/team/update',3,3,0,'2023-09-29 08:40:00.879'),(35,'删除团队','POST','/api/sys/team/del',3,3,0,'2023-09-29 08:40:00.891'),(64,'分页获取GenTables','POST','/api/v1/sys/gen-tables/page',3,3,0,'2023-10-28 16:31:47.687'),(65,'根据id获取GenTables','POST','/api/v1/sys/gen-tables/get',3,3,0,'2023-10-28 16:31:47.701'),(66,'创建GenTables','POST','/api/v1/sys/gen-tables/create',3,3,0,'2023-10-28 16:31:47.710'),(67,'修改GenTables','POST','/api/v1/sys/gen-tables/update',3,3,0,'2023-10-28 16:31:47.720'),(68,'删除GenTables','POST','/api/v1/sys/gen-tables/del',3,3,0,'2023-10-28 16:31:47.728'),(69,'分页获取GenColumns','POST','/api/v1/sys/gen-columns/page',3,3,0,'2023-10-28 16:35:01.328'),(70,'根据id获取GenColumns','POST','/api/v1/sys/gen-columns/get',3,3,0,'2023-10-28 16:35:01.350'),(71,'创建GenColumns','POST','/api/v1/sys/gen-columns/create',3,3,0,'2023-10-28 16:35:01.360'),(72,'修改GenColumns','POST','/api/v1/sys/gen-columns/update',3,3,0,'2023-10-28 16:35:01.368'),(73,'删除GenColumns','POST','/api/v1/sys/gen-columns/del',3,3,0,'2023-10-28 16:35:01.388'),(74,'分页获取接口','POST','/api/v1/sys/sys-api/page',3,3,0,'2023-10-28 16:51:44.921'),(75,'根据id获取接口','POST','/api/v1/sys/sys-api/get',3,3,0,'2023-10-28 16:51:44.933'),(76,'创建接口','POST','/api/sys/api/create',3,3,0,'2023-10-28 16:51:44.943'),(77,'修改接口','POST','/api/sys/api/update',3,3,0,'2023-10-28 16:51:44.952'),(78,'删除接口','POST','/api/sys/api/del',3,3,0,'2023-10-28 16:51:44.960'),(80,'所有角色','POST','/api/sys/role/list',3,3,0,'2023-11-04 11:51:22.000'),(81,'分页获取配置','POST','/api/sys/cfg/page',3,3,0,'2023-11-08 20:34:50.872'),(82,'根据id获取配置','POST','/api/sys/cfg/get',3,3,0,'2023-11-08 20:34:50.885'),(83,'创建配置','POST','/api/sys/cfg/create',3,3,0,'2023-11-08 20:34:50.898'),(84,'修改配置','POST','/api/sys/cfg/update',3,3,0,'2023-11-08 20:34:50.912'),(85,'删除配置','POST','/api/sys/cfg/del',3,3,0,'2023-11-08 20:34:50.925'),(86,'分页获取操作日志','POST','/api/sys/opera-log/page',3,3,0,'2023-11-09 22:07:17.252'),(87,'根据id获取操作日志','POST','/api/sys/opera-log/get',3,3,0,'2023-11-09 22:07:17.268'),(88,'创建操作日志','POST','/api/sys/opera-log/create',3,3,0,'2023-11-09 22:07:17.284'),(89,'修改操作日志','POST','/api/sys/opera-log/update',3,3,0,'2023-11-09 22:07:17.297'),(90,'删除操作日志','POST','/api/sys/opera-log/del',3,3,0,'2023-11-09 22:07:17.310'),(92,'修改我的企业信息','POST','/api/team/change',3,3,0,'2023-11-18 15:45:42.000'),(93,'监控','POST','/api/v1/tools/monitor',3,3,0,'2023-11-20 19:42:50.000'),(94,'分页获取用户通知','POST','/api/notice/user-notice/page',3,3,0,'2023-11-22 21:45:33.859'),(95,'根据id获取用户通知','POST','/api/notice/user-notice/get',3,3,0,'2023-11-22 21:45:33.885'),(96,'创建用户通知','POST','/api/notice/user-notice/create',3,3,0,'2023-11-22 21:45:33.902'),(97,'修改用户通知','POST','/api/notice/user-notice/update',3,3,0,'2023-11-22 21:45:33.918'),(98,'删除用户通知','POST','/api/notice/user-notice/del',3,3,0,'2023-11-22 21:45:33.936'),(99,'分页获取公用通知','POST','/api/notice/pub-notice/page',3,3,0,'2023-11-22 21:45:36.388'),(100,'根据id获取公用通知','POST','/api/notice/pub-notice/get',3,3,0,'2023-11-22 21:45:36.400'),(101,'创建公用通知','POST','/api/notice/pub-notice/create',3,3,0,'2023-11-22 21:45:36.411'),(102,'修改公用通知','POST','/api/notice/pub-notice/update',3,3,0,'2023-11-22 21:45:36.422'),(103,'删除公用通知','POST','/api/notice/pub-notice/del',3,3,0,'2023-11-22 21:45:36.433'),(104,'分页获取Task','POST','/api/notice/task/page',3,3,0,'2023-11-24 15:39:24.097'),(105,'根据id获取Task','POST','/api/notice/task/get',3,3,0,'2023-11-24 15:39:24.113'),(106,'创建Task','POST','/api/notice/task/create',3,3,0,'2023-11-24 15:39:24.125'),(107,'修改Task','POST','/api/notice/task/update',3,3,0,'2023-11-24 15:39:24.139'),(108,'删除Task','POST','/api/notice/task/del',3,3,0,'2023-11-24 15:39:24.150'),(302,'分页获取管理员用户','POST','/api/sys/admin/page',3,3,0,'2025-03-15 16:17:48.373'),(303,'根据id获取管理员用户','POST','/api/sys/admin/get',3,3,0,'2025-03-15 16:17:48.384'),(304,'创建管理员用户','POST','/api/sys/admin/create',3,3,0,'2025-03-15 16:17:48.399'),(305,'修改管理员用户','POST','/api/sys/admin/update',3,3,0,'2025-03-15 16:17:48.412'),(306,'删除管理员用户','POST','/api/sys/admin/del',3,3,0,'2025-03-15 16:17:48.424'),(317,'分页获取用户登录记录','POST','/api/sys/user-login-log/page',3,3,0,'2025-03-15 16:43:48.106'),(318,'根据id获取用户登录记录','POST','/api/sys/user-login-log/get',3,3,0,'2025-03-15 16:43:48.124'),(319,'创建用户登录记录','POST','/api/sys/user-login-log/create',3,3,0,'2025-03-15 16:43:48.140'),(320,'修改用户登录记录','POST','/api/sys/user-login-log/update',3,3,0,'2025-03-15 16:43:48.161'),(321,'删除用户登录记录','POST','/api/sys/user-login-log/del',3,3,0,'2025-03-15 16:43:48.178'),(322,'分页获取菜单','POST','/api/team/menu/all',3,3,0,'2023-09-26 13:47:37.020'),(323,'根据id获取菜单','POST','/api/team/menu/get',3,3,0,'2023-09-26 13:47:37.038'),(324,'创建菜单','POST','/api/team/menu/create',3,3,0,'2023-09-26 13:47:37.064'),(325,'修改菜单','POST','/api/team/menu/update',3,3,0,'2023-09-26 13:47:37.073'),(326,'删除菜单','POST','/api/team/menu/del',3,3,0,'2023-09-26 13:47:37.082'),(327,'分页获取角色','POST','/api/team/role/page',3,3,0,'2023-09-26 13:47:39.685'),(328,'根据id获取角色','POST','/api/team/role/get',3,3,0,'2023-09-26 13:47:39.701'),(329,'创建角色','POST','/api/team/role/create',3,3,0,'2023-09-26 13:47:39.734'),(330,'修改角色','POST','/api/team/role/update',3,3,0,'2023-09-26 13:47:39.762'),(331,'删除角色','POST','/api/team/role/del',3,3,0,'2023-09-26 13:47:39.773'),(332,'分页获取部门','POST','/api/team/dept/all',3,3,0,'2023-09-26 13:47:42.136'),(333,'根据id获取部门','POST','/api/team/dept/get',3,3,0,'2023-09-26 13:47:42.152'),(334,'创建部门','POST','/api/team/dept/create',3,3,0,'2023-09-26 13:47:42.172'),(335,'修改部门','POST','/api/team/dept/update',3,3,0,'2023-09-26 13:47:42.186'),(336,'删除部门','POST','/api/team/dept/del',3,3,0,'2023-09-26 13:47:42.197'),(337,'分页获取成员','POST','/api/team/member/page',3,3,0,'2023-09-26 14:09:40.443'),(338,'根据id获取成员','POST','/api/team/member/get',3,3,0,'2023-09-26 14:09:40.459'),(339,'创建成员','POST','/api/team/member/create',3,3,0,'2023-09-26 14:09:40.483'),(340,'修改成员','POST','/api/team/member/update',3,3,0,'2023-09-26 14:09:40.493'),(341,'删除成员','POST','/api/team/member/del',3,3,0,'2023-09-26 14:09:40.502'),(342,'分页获取团队','POST','/api/team/team/page',3,3,0,'2023-09-29 08:40:00.840'),(343,'根据id获取团队','POST','/api/team/team/get',3,3,0,'2023-09-29 08:40:00.855'),(344,'创建团队','POST','/api/team/team/create',3,3,0,'2023-09-29 08:40:00.867'),(345,'修改团队','POST','/api/team/team/update',3,3,0,'2023-09-29 08:40:00.879'),(346,'删除团队','POST','/api/team/team/del',3,3,0,'2023-09-29 08:40:00.891'),(347,'团队用户菜单','POST','/api/team/user/menus',3,3,0,'2023-09-29 08:40:00.891'),(348,'已授权菜单','POST','/api/team/menu/grant',3,3,0,'2023-09-29 08:40:00.891'),(349,'分页获取部门','POST','/api/sys/admin-dept/page',3,3,0,'2025-03-17 20:44:16.549'),(350,'根据id获取部门','POST','/api/sys/admin-dept/get',3,3,0,'2025-03-17 20:44:16.569'),(351,'创建部门','POST','/api/sys/admin-dept/create',3,3,0,'2025-03-17 20:44:16.587'),(352,'修改部门','POST','/api/sys/admin-dept/update',3,3,0,'2025-03-17 20:44:16.604'),(353,'删除部门','POST','/api/sys/admin-dept/del',3,3,0,'2025-03-17 20:44:16.621'),(354,'分页获取角色','POST','/api/sys/admin-role/page',3,3,0,'2025-03-17 20:44:28.393'),(355,'根据id获取角色','POST','/api/sys/admin-role/get',3,3,0,'2025-03-17 20:44:28.411'),(356,'创建角色','POST','/api/sys/admin-role/create',3,3,0,'2025-03-17 20:44:28.431'),(357,'修改角色','POST','/api/sys/admin-role/update',3,3,0,'2025-03-17 20:44:28.450'),(358,'删除角色','POST','/api/sys/admin-role/del',3,3,0,'2025-03-17 20:44:28.470'),(359,'所有角色','POST','/api/sys/admin-role/list',3,3,0,'2025-03-17 22:28:59.000'),(360,'所有部门','POST','/api/sys/admin-dept/all',3,3,0,'2025-03-17 22:29:02.000'),(362,'删除Task','POST','/api/team/task/del',3,3,0,'2023-11-24 15:39:24.150'),(363,'分页获取公用通知','POST','/api/team/pub-notice/page',3,3,0,'2023-11-22 21:45:36.388'),(364,'根据id获取公用通知','POST','/api/team/pub-notice/get',3,3,0,'2023-11-22 21:45:36.400'),(365,'创建公用通知','POST','/api/team/pub-notice/create',3,3,0,'2023-11-22 21:45:36.411'),(366,'修改公用通知','POST','/api/team/pub-notice/update',3,3,0,'2023-11-22 21:45:36.422'),(367,'删除公用通知','POST','/api/team/pub-notice/del',3,3,0,'2023-11-22 21:45:36.433'),(368,'分页获取用户通知','POST','/api/team/user-notice/page',3,3,0,'2023-11-22 21:45:33.859'),(369,'根据id获取用户通知','POST','/api/team/user-notice/get',3,3,0,'2023-11-22 21:45:33.885'),(370,'创建用户通知','POST','/api/team/user-notice/create',3,3,0,'2023-11-22 21:45:33.902'),(371,'修改用户通知','POST','/api/team/user-notice/update',3,3,0,'2023-11-22 21:45:33.918'),(372,'删除用户通知','POST','/api/team/user-notice/del',3,3,0,'2023-11-22 21:45:33.936'),(373,'分页获取Task','POST','/api/team/task/page',3,3,0,'2023-11-24 15:39:24.097'),(374,'根据id获取Task','POST','/api/team/task/get',3,3,0,'2023-11-24 15:39:24.113'),(375,'创建Task','POST','/api/team/task/create',3,3,0,'2023-11-24 15:39:24.125'),(376,'修改Task','POST','/api/team/task/update',3,3,0,'2023-11-24 15:39:24.139');
/*!40000 ALTER TABLE `sys_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_cfg`
--

DROP TABLE IF EXISTS `sys_cfg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_cfg` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '名字',
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'Type',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'key',
  `val` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '值',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'Remark',
  `status` tinyint DEFAULT NULL COMMENT 'Status',
  `team_id` int DEFAULT NULL COMMENT '团队id -1为系统',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `sys_cfg_unique` (`team_id`,`type`,`key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='系统配置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_cfg`
--

LOCK TABLES `sys_cfg` WRITE;
/*!40000 ALTER TABLE `sys_cfg` DISABLE KEYS */;
INSERT INTO `sys_cfg` VALUES (1,NULL,'1111','222','2222','法大师傅',1,-1,1,'2025-03-21 17:23:35.936');
/*!40000 ALTER TABLE `sys_cfg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_email`
--

DROP TABLE IF EXISTS `sys_email`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_email` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱地址',
  `code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '验证码',
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '类型',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `use_status` tinyint DEFAULT NULL COMMENT '使用状态',
  `created_at` int unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` int unsigned DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sys_email_email_IDX` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='邮件';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_email`
--

LOCK TABLES `sys_email` WRITE;
/*!40000 ALTER TABLE `sys_email` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_email` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_job`
--

DROP TABLE IF EXISTS `sys_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_job` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '名称',
  `job_group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '组',
  `job_type` tinyint DEFAULT NULL COMMENT '类型',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '表达式',
  `invoke_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '调用目标',
  `args` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '参数',
  `misfire_policy` bigint DEFAULT NULL COMMENT '策略',
  `concurrent` tinyint DEFAULT NULL COMMENT '并发',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `entry_id` smallint DEFAULT NULL COMMENT '任务id',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_job_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='定时任务';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_job`
--

LOCK TABLES `sys_job` WRITE;
/*!40000 ALTER TABLE `sys_job` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `menu_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '菜单名',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '显示名称',
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '路径',
  `platform_type` bigint DEFAULT NULL COMMENT '平台类型 1 平台管理 2团队管理',
  `menu_type` tinyint DEFAULT NULL COMMENT '菜单类型 1 分类 2菜单 3方法按钮',
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '权限',
  `parent_id` int unsigned DEFAULT NULL COMMENT '菜单父id',
  `no_cache` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '前端组件路径',
  `sort` tinyint DEFAULT NULL COMMENT '排序倒叙',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_menu_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_menu_create_by` (`create_by`) USING BTREE,
  KEY `idx_sys_menu_update_by` (`update_by`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=453 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (1,'','系统管理','setting','/sys',1,1,'',0,0,'Layout',50,0,1,0,'2023-09-26 13:46:59.480','2024-03-13 20:49:05.368',NULL),(2,'SysUserManage','用户管理','flUser','/sys/sys-user',1,2,'sys:sysUser:list',394,0,'/sys/sys-user/index',0,0,1,1,'2023-09-26 13:46:59.493','2025-03-15 16:41:45.297',NULL),(3,'','用户详情','','sys_user_detail',1,3,'sys:sysUser:query',2,0,'',0,0,1,1,'2023-09-26 13:46:59.522','2023-09-26 13:46:59.526',NULL),(4,'','用户创建','','sys_user_create',1,3,'sys:sysUser:add',2,0,'',0,0,1,1,'2023-09-26 13:46:59.534','2023-09-26 13:46:59.536',NULL),(5,'','用户修改','','sys_user_update',1,3,'sys:sysUser:edit',2,0,'',0,0,1,1,'2023-09-26 13:46:59.543','2023-09-26 13:46:59.546',NULL),(6,'','用户删除','','sys_user_del',1,3,'sys:sysUser:remove',2,0,'',0,0,1,1,'2023-09-26 13:46:59.553','2023-09-26 13:46:59.555',NULL),(7,'SysMenuManage','菜单管理','menu','/sys/sys-menu',1,2,'sys:sysMenu:list',91,0,'/sys/sys-menu/index',0,0,1,1,'2023-09-26 13:47:37.025','2024-05-24 22:47:05.611',NULL),(8,'','菜单详情','','sys_menu_detail',1,3,'sys:sysMenu:query',7,0,'',0,0,1,1,'2023-09-26 13:47:37.042','2023-09-26 13:47:37.055',NULL),(9,'','菜单创建','','sys_menu_create',1,3,'sys:sysMenu:add',7,0,'',0,0,1,1,'2023-09-26 13:47:37.067','2023-09-26 13:47:37.069',NULL),(10,'','菜单修改','','sys_menu_update',1,3,'sys:sysMenu:edit',7,0,'',0,0,1,1,'2023-09-26 13:47:37.076','2023-09-26 13:47:37.079',NULL),(11,'','菜单删除','','sys_menu_del',1,3,'sys:sysMenu:remove',7,0,'',0,0,1,1,'2023-09-26 13:47:37.084','2023-09-26 13:47:37.086',NULL),(12,'SysRoleManage','角色管理','role','/sys/sys-role',1,2,'sys:sysRole:list',394,0,'/sys/sys-role/index',14,1,1,1,'2023-09-26 13:47:39.688','2025-03-17 16:18:52.560',NULL),(13,'','角色详情','','sys_role_detail',1,3,'sys:sysRole:query',12,0,'',0,0,1,1,'2023-09-26 13:47:39.714','2023-09-26 13:47:39.717',NULL),(14,'','角色创建','','sys_role_create',1,3,'sys:sysRole:add',12,0,'',0,0,1,1,'2023-09-26 13:47:39.747','2023-09-26 13:47:39.752',NULL),(15,'','角色修改','','sys_role_update',1,3,'sys:sysRole:edit',12,0,'',0,0,1,1,'2023-09-26 13:47:39.765','2023-09-26 13:47:39.769',NULL),(16,'','角色删除','','sys_role_del',1,3,'sys:sysRole:remove',12,0,'',0,0,1,1,'2023-09-26 13:47:39.776','2023-09-26 13:47:39.778',NULL),(17,'SysDeptManage','部门管理','dept','/sys/sys-dept',1,2,'sys:sysDept:list',394,0,'/sys/sys-dept/index',12,1,1,1,'2023-09-26 13:47:42.140','2025-03-17 16:18:44.940',NULL),(18,'','部门详情','','sys_dept_detail',1,3,'sys:sysDept:query',17,0,'',0,0,1,1,'2023-09-26 13:47:42.155','2023-09-26 13:47:42.167',NULL),(19,'','部门创建','','sys_dept_create',1,3,'sys:sysDept:add',17,0,'',0,0,1,1,'2023-09-26 13:47:42.180','2023-09-26 13:47:42.182',NULL),(20,'','部门修改','','sys_dept_update',1,3,'sys:sysDept:edit',17,0,'',0,0,1,1,'2023-09-26 13:47:42.188','2023-09-26 13:47:42.190',NULL),(21,'','部门删除','','sys_dept_del',1,3,'sys:sysDept:remove',17,0,'',0,0,1,1,'2023-09-26 13:47:42.200','2023-09-26 13:47:42.202',NULL),(52,'SysTeamManage','团队管理','team','/sys/sys-team',1,2,'sys:sysTeam:list',394,0,'/sys/sys-team/index',10,0,1,1,'2023-09-29 08:44:07.932','2025-03-15 16:41:55.374',NULL),(53,'','团队详情','','sys_team_detail',1,3,'sys:sysTeam:query',52,0,'',0,0,1,1,'2023-09-29 08:44:07.941','2023-09-29 08:44:07.945',NULL),(54,'','团队创建','','sys_team_create',1,3,'sys:sysTeam:add',52,0,'',0,0,1,1,'2023-09-29 08:44:07.951','2023-09-29 08:44:07.954',NULL),(55,'','团队修改','','sys_team_update',1,3,'sys:sysTeam:edit',52,0,'',0,0,1,1,'2023-09-29 08:44:07.958','2023-09-29 08:44:07.961',NULL),(56,'','团队删除','','sys_team_del',1,3,'sys:sysTeam:remove',52,0,'',0,0,1,1,'2023-09-29 08:44:07.966','2023-09-29 08:44:07.971',NULL),(57,'SysMemberManage','成员管理','openArm','/sys/sys-member',1,2,'sys:sysMember:list',394,0,'/sys/sys-member/index',16,1,1,1,'2023-09-29 08:44:33.484','2025-03-17 16:18:58.556',NULL),(58,'','成员详情','','sys_member_detail',1,3,'sys:sysMember:query',57,0,'',0,0,1,1,'2023-09-29 08:44:33.517','2023-09-29 08:44:33.520',NULL),(59,'','成员创建','','sys_member_create',1,3,'sys:sysMember:add',57,0,'',0,0,1,1,'2023-09-29 08:44:33.526','2023-09-29 08:44:33.529',NULL),(60,'','成员修改','','sys_member_update',1,3,'sys:sysMember:edit',57,0,'',0,0,1,1,'2023-09-29 08:44:33.534','2023-09-29 08:44:33.538',NULL),(61,'','会员删除','','sys_member_del',1,3,'sys:sysMember:remove',57,0,'',0,0,1,1,'2023-09-29 08:44:33.543','2023-09-29 08:44:33.546',NULL),(91,'','工具管理','setUp','/tools',1,1,NULL,0,0,'Layout',111,0,1,1,'2023-10-28 09:41:26.000','2024-05-24 22:47:52.057',NULL),(93,'GenTablesManage','代码生成','table','/sys/gen-tables',1,2,'sys:genTables:list',91,0,'/tool/gen-tables/index',80,0,1,1,'2023-10-28 16:31:47.693','2025-03-15 16:36:31.947',NULL),(94,'','GenTables详情','','gen_tables_detail',1,3,'sys:genTables:query',93,0,'',0,0,1,1,'2023-10-28 16:31:47.704','2023-10-28 16:31:47.704',NULL),(95,'','GenTables创建','','gen_tables_create',1,3,'sys:genTables:add',93,0,'',0,0,1,1,'2023-10-28 16:31:47.714','2023-10-28 16:31:47.714',NULL),(96,'','GenTables修改','','gen_tables_update',1,3,'sys:genTables:edit',93,0,'',0,0,1,1,'2023-10-28 16:31:47.723','2023-10-28 16:31:47.723',NULL),(97,'','GenTables删除','','gen_tables_del',1,3,'sys:genTables:remove',93,0,'',0,0,1,1,'2023-10-28 16:31:47.731','2023-10-28 16:31:47.731',NULL),(103,'SysApiManage','接口管理','swapLine','/sys/sys-api',1,2,'sys:sysApi:list',91,0,'/tool/sys-api/index',90,0,1,1,'2023-10-28 16:51:44.924','2025-03-15 16:36:24.788',NULL),(104,'','接口详情','','sys_api_detail',1,3,'sys:sysApi:query',103,0,'',0,0,1,1,'2023-10-28 16:51:44.938','2023-10-28 16:51:44.938',NULL),(105,'','接口创建','','sys_api_create',1,3,'sys:sysApi:add',103,0,'',0,0,1,1,'2023-10-28 16:51:44.947','2023-10-28 16:51:44.947',NULL),(106,'','接口修改','','sys_api_update',1,3,'sys:sysApi:edit',103,0,'',0,0,1,1,'2023-10-28 16:51:44.955','2023-10-28 16:51:44.955',NULL),(107,'','接口删除','','sys_api_del',1,3,'sys:sysApi:remove',103,0,'',0,0,1,1,'2023-10-28 16:51:44.964','2023-10-28 16:51:44.964',NULL),(111,'SysCfgManage','配置管理','operation','/sys/sys-cfg',1,2,'sys:sysCfg:list',1,0,'/sys/sys-cfg/index',21,0,1,1,'2023-11-08 20:34:50.875','2025-03-15 16:31:48.457',NULL),(112,'','配置详情','','sys_cfg_detail',1,3,'sys:sysCfg:query',111,0,'',0,0,1,1,'2023-11-08 20:34:50.889','2023-11-08 20:34:50.889',NULL),(113,'','配置创建','','sys_cfg_create',1,3,'sys:sysCfg:add',111,0,'',0,0,1,1,'2023-11-08 20:34:50.903','2023-11-08 20:34:50.903',NULL),(114,'','配置修改','','sys_cfg_update',1,3,'sys:sysCfg:edit',111,0,'',0,0,1,1,'2023-11-08 20:34:50.916','2023-11-08 20:34:50.916',NULL),(115,'','配置删除','','sys_cfg_del',1,3,'sys:sysCfg:remove',111,0,'',0,0,1,1,'2023-11-08 20:34:50.929','2023-11-08 20:34:50.929',NULL),(117,'SysOperaLogManage','操作日志管理','generate','/sys/sys-opera-log',1,2,'sys:sysOperaLog:list',394,0,'/sys/sys-opera-log/index',19,0,1,1,'2023-11-09 22:07:17.257','2025-03-15 16:42:31.824',NULL),(118,'','操作日志详情','','sys_opera_log_detail',1,3,'sys:sysOperaLog:query',117,0,'',0,0,1,1,'2023-11-09 22:07:17.272','2023-11-09 22:07:17.272',NULL),(119,'','操作日志创建','','sys_opera_log_create',1,3,'sys:sysOperaLog:add',117,0,'',0,0,1,1,'2023-11-09 22:07:17.288','2023-11-09 22:07:17.288',NULL),(120,'','操作日志修改','','sys_opera_log_update',1,3,'sys:sysOperaLog:edit',117,0,'',0,0,1,1,'2023-11-09 22:07:17.301','2023-11-09 22:07:17.301',NULL),(121,'','操作日志删除','','sys_opera_log_del',1,3,'sys:sysOperaLog:remove',117,0,'',0,0,1,1,'2023-11-09 22:07:17.313','2023-11-09 22:07:17.313',NULL),(122,'My','我的',NULL,'/my',1,1,'',0,0,'Layout',123,1,1,1,'2023-11-18 15:25:44.000','2024-05-24 23:19:32.049',NULL),(123,NULL,'修改企业信息',NULL,'my_team',1,3,'my:change:team',122,0,NULL,0,0,1,1,'2023-11-18 15:33:28.000','2023-11-18 15:33:31.000',NULL),(124,'Monitor','监控','monitor','/tools/monitor',1,2,'sys:monitor',91,0,'/tool/monitor',100,0,1,1,'2023-11-20 19:46:04.000','2025-03-15 16:36:13.924',NULL),(126,'UserNoticeManage','用户通知管理','bell','/notice/user-notice',1,2,'notice:userNotice:list',394,0,'/notice/user-notice/index',20,0,1,1,'2023-11-22 21:45:33.866','2025-03-15 16:42:50.590',NULL),(127,'','用户通知详情','','user_notice_detail',1,3,'notice:userNotice:query',126,0,'',0,0,1,1,'2023-11-22 21:45:33.889','2023-11-22 21:45:33.889',NULL),(128,'','用户通知创建','','user_notice_create',1,3,'notice:userNotice:add',126,0,'',0,0,1,1,'2023-11-22 21:45:33.907','2023-11-22 21:45:33.907',NULL),(129,'','用户通知修改','','user_notice_update',1,3,'notice:userNotice:edit',126,0,'',0,0,1,1,'2023-11-22 21:45:33.925','2023-11-22 21:45:33.925',NULL),(130,'','用户通知删除','','user_notice_del',1,3,'notice:userNotice:remove',126,0,'',0,0,1,1,'2023-11-22 21:45:33.941','2023-11-22 21:45:33.941',NULL),(132,'PubNoticeManage','公用通知管理','bellFill','/notice/pub-notice',1,2,'notice:pubNotice:list',1,0,'/notice/pub-notice/index',28,0,1,1,'2023-11-22 21:45:36.391','2025-03-15 16:32:22.933',NULL),(133,'','公用通知详情','','pub_notice_detail',1,3,'notice:pubNotice:query',132,0,'',0,0,1,1,'2023-11-22 21:45:36.405','2023-11-22 21:45:36.405',NULL),(134,'','公用通知创建','','pub_notice_create',1,3,'notice:pubNotice:add',132,0,'',0,0,1,1,'2023-11-22 21:45:36.416','2023-11-22 21:45:36.416',NULL),(135,'','公用通知修改','','pub_notice_update',1,3,'notice:pubNotice:edit',132,0,'',0,0,1,1,'2023-11-22 21:45:36.427','2023-11-22 21:45:36.427',NULL),(136,'','公用通知删除','','pub_notice_del',1,3,'notice:pubNotice:remove',132,0,'',0,0,1,1,'2023-11-22 21:45:36.439','2023-11-22 21:45:36.439',NULL),(138,'TaskManage','Task管理','task','/notice/task',1,2,'notice:task:list',1,0,'/notice/task/index',25,0,1,1,'2023-11-24 15:39:24.103','2025-03-15 16:32:14.006',NULL),(139,'','Task详情','','task_detail',1,3,'notice:task:query',138,0,'',0,0,1,1,'2023-11-24 15:39:24.118','2023-11-24 15:39:24.118',NULL),(140,'','Task创建','','task_create',1,3,'notice:task:add',138,0,'',0,0,1,1,'2023-11-24 15:39:24.131','2023-11-24 15:39:24.131',NULL),(141,'','Task修改','','task_update',1,3,'notice:task:edit',138,0,'',0,0,1,1,'2023-11-24 15:39:24.142','2023-11-24 15:39:24.142',NULL),(142,'','Task删除','','task_del',1,3,'notice:task:remove',138,0,'',0,0,1,1,'2023-11-24 15:39:24.155','2023-11-24 15:39:24.155',NULL),(394,'','租户管理','team','/tenant',1,1,'',0,0,'Layout',30,0,1,1,'2023-09-26 13:46:59.480','2025-03-15 16:41:33.128',NULL),(395,'','用户登录记录','log','/sys',1,1,'',0,0,'Layout',0,0,1,0,'2025-03-15 16:43:48.097','2025-03-15 16:43:48.097','2025-03-15 16:44:16.727'),(396,'SysUserLoginLogManage','用户登录记录管理','log','/sys/sys-user-login-log',1,2,'sys:sysUserLoginLog:list',394,0,'/sys/sys-user-login-log/index',50,0,1,1,'2025-03-15 16:43:48.112','2025-03-15 16:48:04.165',NULL),(397,'','用户登录记录详情','','sys_user_login_log_detail',1,3,'sys:sysUserLoginLog:query',396,0,'',0,0,1,1,'2025-03-15 16:43:48.128','2025-03-15 16:43:48.128',NULL),(398,'','用户登录记录创建','','sys_user_login_log_create',1,3,'sys:sysUserLoginLog:add',396,0,'',0,0,1,1,'2025-03-15 16:43:48.150','2025-03-15 16:43:48.150',NULL),(399,'','用户登录记录修改','','sys_user_login_log_update',1,3,'sys:sysUserLoginLog:edit',396,0,'',0,0,1,1,'2025-03-15 16:43:48.169','2025-03-15 16:43:48.169',NULL),(400,'','用户登录记录删除','','sys_user_login_log_del',1,3,'sys:sysUserLoginLog:remove',396,0,'',0,0,1,1,'2025-03-15 16:43:48.183','2025-03-15 16:43:48.183',NULL),(401,'','内部管理','userSetting','/sys/internal',1,1,'',0,0,'Layout',30,0,1,1,'2025-03-17 09:25:50.030','2025-03-17 20:46:42.129',NULL),(402,'AdminManage','超管管理','userSetting','/sys/admin',1,2,'sys:admin:list',401,0,'/sys/admin/index',0,0,1,1,'2025-03-17 09:25:50.050','2025-03-17 10:36:27.593',NULL),(403,'','用户详情','','admin_detail',1,3,'sys:admin:query',402,0,'',0,0,1,1,'2025-03-17 09:25:50.068','2025-03-17 09:25:50.068',NULL),(404,'','用户创建','','admin_create',1,3,'sys:admin:add',402,0,'',0,0,1,1,'2025-03-17 09:25:50.088','2025-03-17 09:25:50.088',NULL),(405,'','用户修改','','admin_update',1,3,'sys:admin:edit',402,0,'',0,0,1,1,'2025-03-17 09:25:50.099','2025-03-17 09:25:50.099',NULL),(406,'','用户删除','','admin_del',1,3,'sys:admin:remove',402,0,'',0,0,1,1,'2025-03-17 09:25:50.117','2025-03-17 09:25:50.117',NULL),(407,'','系统管理','setting','/sys',3,1,'',0,0,'Layout',50,0,1,0,'2023-09-26 13:46:59.480','2024-03-13 20:49:05.368',NULL),(408,'SysRoleManage','角色管理','role','/sys/sys-role',3,2,'sys:sysRole:list',407,0,'/sys/sys-role/index',14,0,1,1,'2023-09-26 13:47:39.688','2025-03-15 16:42:09.754',NULL),(409,'','角色详情','','sys_role_detail',3,3,'sys:sysRole:query',408,0,'',0,0,1,1,'2023-09-26 13:47:39.714','2023-09-26 13:47:39.717',NULL),(410,'','角色创建','','sys_role_create',3,3,'sys:sysRole:add',408,0,'',0,0,1,1,'2023-09-26 13:47:39.747','2023-09-26 13:47:39.752',NULL),(411,'','角色修改','','sys_role_update',3,3,'sys:sysRole:edit',408,0,'',0,0,1,1,'2023-09-26 13:47:39.765','2023-09-26 13:47:39.769',NULL),(412,'','角色删除','','sys_role_del',3,3,'sys:sysRole:remove',408,0,'',0,0,1,1,'2023-09-26 13:47:39.776','2023-09-26 13:47:39.778',NULL),(413,'SysDeptManage','部门管理','dept','/sys/sys-dept',3,2,'sys:sysDept:list',407,0,'/sys/sys-dept/index',12,0,1,1,'2023-09-26 13:47:42.140','2025-03-15 16:42:02.161',NULL),(414,'','部门详情','','sys_dept_detail',3,3,'sys:sysDept:query',413,0,'',0,0,1,1,'2023-09-26 13:47:42.155','2023-09-26 13:47:42.167',NULL),(415,'','部门创建','','sys_dept_create',3,3,'sys:sysDept:add',413,0,'',0,0,1,1,'2023-09-26 13:47:42.180','2023-09-26 13:47:42.182',NULL),(416,'','部门修改','','sys_dept_update',3,3,'sys:sysDept:edit',413,0,'',0,0,1,1,'2023-09-26 13:47:42.188','2023-09-26 13:47:42.190',NULL),(417,'','部门删除','','sys_dept_del',3,3,'sys:sysDept:remove',413,0,'',0,0,1,1,'2023-09-26 13:47:42.200','2023-09-26 13:47:42.202',NULL),(418,'','团队详情','','sys_team_detail',3,3,'sys:sysTeam:query',413,0,'',0,0,1,1,'2023-09-29 08:44:07.941','2023-09-29 08:44:07.945',NULL),(419,'','团队修改','','sys_team_update',3,3,'sys:sysTeam:edit',413,0,'',0,0,1,1,'2023-09-29 08:44:07.958','2023-09-29 08:44:07.961',NULL),(420,'SysMemberManage','成员管理','openArm','/sys/sys-member',3,2,'sys:sysMember:list',407,0,'/sys/sys-member/index',16,0,1,1,'2023-09-29 08:44:33.484','2025-03-15 16:42:17.377',NULL),(421,'','成员详情','','sys_member_detail',3,3,'sys:sysMember:query',420,0,'',0,0,1,1,'2023-09-29 08:44:33.517','2023-09-29 08:44:33.520',NULL),(422,'','成员创建','','sys_member_create',3,3,'sys:sysMember:add',420,0,'',0,0,1,1,'2023-09-29 08:44:33.526','2023-09-29 08:44:33.529',NULL),(423,'','成员修改','','sys_member_update',3,3,'sys:sysMember:edit',420,0,'',0,0,1,1,'2023-09-29 08:44:33.534','2023-09-29 08:44:33.538',NULL),(424,'','成员删除','','sys_member_del',3,3,'sys:sysMember:remove',420,0,'',0,0,1,1,'2023-09-29 08:44:33.543','2023-09-29 08:44:33.546',NULL),(425,'','部门','dept','/sys',1,1,'',0,0,'Layout',0,0,1,0,'2025-03-17 20:44:16.535','2025-03-17 20:44:16.535','2025-03-17 20:45:33.142'),(426,'AdminDeptManage','部门管理','dept','/sys/admin-dept',1,2,'sys:adminDept:list',401,0,'/sys/admin-dept/index',0,0,1,1,'2025-03-17 20:44:16.556','2025-03-17 20:45:13.260',NULL),(427,'','部门详情','','admin_dept_detail',1,3,'sys:adminDept:query',426,0,'',0,0,1,1,'2025-03-17 20:44:16.577','2025-03-17 20:44:16.577',NULL),(428,'','部门创建','','admin_dept_create',1,3,'sys:adminDept:add',426,0,'',0,0,1,1,'2025-03-17 20:44:16.594','2025-03-17 20:44:16.594',NULL),(429,'','部门修改','','admin_dept_update',1,3,'sys:adminDept:edit',426,0,'',0,0,1,1,'2025-03-17 20:44:16.611','2025-03-17 20:44:16.611',NULL),(430,'','部门删除','','admin_dept_del',1,3,'sys:adminDept:remove',426,0,'',0,0,1,1,'2025-03-17 20:44:16.628','2025-03-17 20:44:16.628',NULL),(431,'','角色','role','/sys',1,1,'',0,0,'Layout',0,0,1,0,'2025-03-17 20:44:28.381','2025-03-17 20:44:28.381','2025-03-17 20:45:30.165'),(432,'AdminRoleManage','角色管理','role','/sys/admin-role',1,2,'sys:adminRole:list',401,0,'/sys/admin-role/index',0,0,1,1,'2025-03-17 20:44:28.400','2025-03-17 20:45:26.529',NULL),(433,'','角色详情','','admin_role_detail',1,3,'sys:adminRole:query',432,0,'',0,0,1,1,'2025-03-17 20:44:28.419','2025-03-17 20:44:28.419',NULL),(434,'','角色创建','','admin_role_create',1,3,'sys:adminRole:add',432,0,'',0,0,1,1,'2025-03-17 20:44:28.439','2025-03-17 20:44:28.439',NULL),(435,'','角色修改','','admin_role_update',1,3,'sys:adminRole:edit',432,0,'',0,0,1,1,'2025-03-17 20:44:28.459','2025-03-17 20:44:28.459',NULL),(436,'','角色删除','','admin_role_del',1,3,'sys:adminRole:remove',432,0,'',0,0,1,1,'2025-03-17 20:44:28.477','2025-03-17 20:44:28.477',NULL),(437,'teamDeptAll','部门管理','','/team/dept',3,3,'team:dept:list',413,0,'',0,0,1,0,'2025-03-21 09:21:51.880','2025-03-21 09:21:51.880','2025-03-21 09:41:23.104'),(438,'UserNoticeManage','团队用户通知管理','bell','/notice/user-notice',3,2,'notice:userNotice:list',407,0,'/notice/user-notice/index',20,0,1,1,'2023-11-22 21:45:33.866','2025-03-15 16:42:50.590','2025-03-21 09:41:23.104'),(439,'','团队用户通知详情','','user_notice_detail',3,3,'notice:userNotice:query',438,0,'',0,0,1,1,'2023-11-22 21:45:33.889','2023-11-22 21:45:33.889','2025-03-21 09:41:23.104'),(440,'','团队用户通知创建','','user_notice_create',3,3,'notice:userNotice:add',438,0,'',0,0,1,1,'2023-11-22 21:45:33.907','2023-11-22 21:45:33.907','2025-03-21 09:41:23.104'),(441,'','团队用户通知修改','','user_notice_update',3,3,'notice:userNotice:edit',438,0,'',0,0,1,1,'2023-11-22 21:45:33.925','2023-11-22 21:45:33.925','2025-03-21 09:41:23.104'),(442,'','团队用户通知删除','','user_notice_del',3,3,'notice:userNotice:remove',438,0,'',0,0,1,1,'2023-11-22 21:45:33.941','2023-11-22 21:45:33.941','2025-03-21 09:41:23.104'),(443,'PubNoticeManage','团队公用通知管理','bellFill','/notice/pub-notice',3,2,'notice:pubNotice:list',407,0,'/notice/pub-notice/index',28,0,1,1,'2023-11-22 21:45:36.391','2025-03-15 16:32:22.933','2025-03-21 09:41:23.104'),(444,'','团队公用通知详情','','pub_notice_detail',3,3,'notice:pubNotice:query',443,0,'',0,0,1,1,'2023-11-22 21:45:36.405','2023-11-22 21:45:36.405','2025-03-21 09:41:23.104'),(445,'','团队公用通知创建','','pub_notice_create',3,3,'notice:pubNotice:add',443,0,'',0,0,1,1,'2023-11-22 21:45:36.416','2023-11-22 21:45:36.416','2025-03-21 09:41:23.104'),(446,'','团队公用通知修改','','pub_notice_update',3,3,'notice:pubNotice:edit',443,0,'',0,0,1,1,'2023-11-22 21:45:36.427','2023-11-22 21:45:36.427','2025-03-21 09:41:23.104'),(447,'','团队公用通知删除','','pub_notice_del',3,3,'notice:pubNotice:remove',443,0,'',0,0,1,1,'2023-11-22 21:45:36.439','2023-11-22 21:45:36.439','2025-03-21 09:41:23.104'),(448,'TaskManage','团队Task管理','task','/notice/task',3,2,'notice:task:list',407,0,'/notice/task/index',25,0,1,1,'2023-11-24 15:39:24.103','2025-03-15 16:32:14.006','2025-03-21 09:41:23.104'),(449,'','团队Task详情','','task_detail',3,3,'notice:task:query',448,0,'',0,0,1,1,'2023-11-24 15:39:24.118','2023-11-24 15:39:24.118','2025-03-21 09:41:23.104'),(450,'','团队Task创建','','task_create',3,3,'notice:task:add',448,0,'',0,0,1,1,'2023-11-24 15:39:24.131','2023-11-24 15:39:24.131','2025-03-21 09:41:23.104'),(451,'','团队Task修改','','task_update',3,3,'notice:task:edit',448,0,'',0,0,1,1,'2023-11-24 15:39:24.142','2023-11-24 15:39:24.142','2025-03-21 09:41:23.104'),(452,'','团队Task删除','','task_del',3,3,'notice:task:remove',448,0,'',0,0,1,1,'2023-11-24 15:39:24.155','2023-11-24 15:39:24.155','2025-03-21 09:41:23.104');
/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menu_api_rule`
--

DROP TABLE IF EXISTS `sys_menu_api_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_menu_api_rule` (
  `sys_menu_id` int unsigned NOT NULL COMMENT '主键',
  `sys_api_id` int unsigned NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_id`,`sys_api_id`) USING BTREE,
  KEY `fk_sys_menu_api_rule_sys_api` (`sys_api_id`) USING BTREE,
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='菜单关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu_api_rule`
--

LOCK TABLES `sys_menu_api_rule` WRITE;
/*!40000 ALTER TABLE `sys_menu_api_rule` DISABLE KEYS */;
INSERT INTO `sys_menu_api_rule` VALUES (2,1),(3,2),(4,3),(5,4),(6,5),(7,6),(8,7),(9,8),(10,9),(11,10),(12,11),(13,12),(14,13),(15,14),(16,15),(17,16),(18,17),(19,18),(20,19),(21,20),(57,26),(58,27),(59,28),(60,29),(61,30),(52,31),(53,32),(54,33),(55,34),(56,35),(93,64),(94,65),(95,66),(96,67),(97,68),(103,74),(104,75),(105,76),(106,77),(107,78),(12,80),(111,81),(112,82),(113,83),(114,84),(115,85),(117,86),(118,87),(119,88),(120,89),(121,90),(123,92),(124,93),(126,94),(127,95),(128,96),(129,97),(130,98),(132,99),(133,100),(134,101),(135,102),(136,103),(138,104),(139,105),(140,106),(141,107),(142,108),(402,302),(403,303),(404,304),(405,305),(406,306),(396,317),(397,318),(398,319),(399,320),(400,321),(413,332),(437,332),(416,335),(420,337),(421,338),(422,339),(423,340),(424,341),(418,343),(419,345),(426,349),(414,350),(427,350),(415,351),(428,351),(429,352),(417,353),(430,353),(432,354),(409,355),(433,355),(410,356),(434,356),(411,357),(435,357),(412,358),(436,358),(408,359),(452,362),(443,363),(444,364),(445,365),(446,366),(447,367),(438,368),(439,369),(440,370),(441,371),(442,372),(448,373),(449,374),(450,375),(451,376);
/*!40000 ALTER TABLE `sys_menu_api_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_opera_log`
--

DROP TABLE IF EXISTS `sys_opera_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_opera_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作模块',
  `business_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作类型',
  `business_types` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'BusinessTypes',
  `method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '函数',
  `request_method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '请求方式 GET POST PUT DELETE',
  `operator_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作类型',
  `oper_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作者',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '访问位置',
  `oper_param` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '请求参数',
  `status` tinyint DEFAULT NULL COMMENT '操作状态 1:成功 2:失败',
  `oper_time` datetime(3) DEFAULT NULL COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'ua',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='操作日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_opera_log`
--

LOCK TABLES `sys_opera_log` WRITE;
/*!40000 ALTER TABLE `sys_opera_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_opera_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_sms`
--

DROP TABLE IF EXISTS `sys_sms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_sms` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号',
  `code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '验证码',
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '类型',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `use_status` tinyint DEFAULT NULL COMMENT '使用状态',
  `created_at` int unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` int unsigned DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sys_sms_phone_IDX` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='短信';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_sms`
--

LOCK TABLES `sys_sms` WRITE;
/*!40000 ALTER TABLE `sys_sms` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_sms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '昵称',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '姓名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '头像',
  `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '签名',
  `birthday` date DEFAULT NULL COMMENT '生日 格式 yyyy-MM-dd',
  `gender` tinyint(1) DEFAULT '2' COMMENT '性别 1男 2女 3未知',
  `team_id` int DEFAULT NULL COMMENT '团队id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `lock_time` datetime(3) DEFAULT NULL COMMENT '锁定结束时间',
  `status` tinyint DEFAULT NULL COMMENT '状态 1正常 ',
  `src_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '来源',
  `inviter` int unsigned DEFAULT NULL COMMENT '邀请人',
  `invite_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邀请码',
  `client_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '客户端id',
  `reg_ip` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '注册ip',
  `ip_location` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_user_update_by` (`update_by`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (2,'tangtang','13800138001',NULL,'$2a$10$hpISaioe30s6d3LYPdFaw.23rm29Yw9a44t8hl9G.tCNYKYoA72iW','糖糖','小唐',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-11-18 20:35:24.140'),(3,'zcm',NULL,NULL,'$2a$10$8ZbmdrNu22DMnRFDjfBugOwvK0fVOixMX.AAeLFbgOkjS9frXHjVW','小梅','小梅',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,'2023-11-12 19:59:39.849'),(4,'liyanlei',NULL,NULL,'$2a$10$sPJOyDdh/J4OFVa27z4LCuz2xtu3KwhyNegkp27hL.G9D3xMMYoKW','阿雷','阿雷',NULL,NULL,NULL,1,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,1,NULL,'2023-11-12 20:03:10.441'),(5,NULL,NULL,NULL,NULL,'小丽','小丽',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(6,NULL,NULL,NULL,NULL,'小珊','小珊',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(7,NULL,NULL,NULL,NULL,'大雁','大雁',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(8,'dajin','','','$2a$10$05iwlWEYT0eN62N1w.zD7ONb9sExNKNLgcsd9UvbxIFEcwq9wElNS','大金','大金','',NULL,NULL,2,0,'',NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,1,'2023-11-30 22:54:28.914','2025-03-20 21:17:04.934'),(9,'xiaoqi',NULL,NULL,'$2a$10$05iwlWEYT0eN62N1w.zD7ONb9sExNKNLgcsd9UvbxIFEcwq9wElNS','小琪','小琪',NULL,NULL,NULL,2,NULL,NULL,NULL,1,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_login_log`
--

DROP TABLE IF EXISTS `sys_user_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` int unsigned DEFAULT NULL COMMENT '用户id',
  `method` tinyint DEFAULT NULL COMMENT '登录方式',
  `ip` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '登录ip',
  `region` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '地域',
  `client_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '客户端',
  `ver` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作系统',
  `os` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作系统',
  `os_ver` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作系统版本',
  `team_id` int DEFAULT NULL COMMENT '团队id',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sys_user_login_log_uid_IDX` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户登录记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_login_log`
--

LOCK TABLES `sys_user_login_log` WRITE;
/*!40000 ALTER TABLE `sys_user_login_log` DISABLE KEYS */;
INSERT INTO `sys_user_login_log` VALUES (33,1,3,'127.0.0.1','','','','','',NULL,'2025-03-15 16:53:02'),(34,1,3,'127.0.0.1','','','','','',NULL,'2025-03-15 16:53:15'),(35,1,3,'127.0.0.1','','','','','',NULL,'2025-03-15 21:07:19'),(36,1,3,'127.0.0.1','','','','','',NULL,'2025-03-16 14:55:45'),(37,2,3,'127.0.0.1','','','','','',NULL,'2025-03-16 15:16:05'),(38,2,3,'127.0.0.1','','','','','',NULL,'2025-03-16 15:17:01'),(39,2,3,'127.0.0.1','','','','','',NULL,'2025-03-16 15:21:54'),(40,2,3,'127.0.0.1','','','','','',NULL,'2025-03-16 15:23:18'),(41,1,3,'127.0.0.1','','','','','',NULL,'2025-03-16 23:44:57'),(42,1,3,'127.0.0.1','','','','','',NULL,'2025-03-16 23:46:52'),(43,1,3,'127.0.0.1','','','','','',NULL,'2025-03-17 09:07:17'),(44,1,3,'127.0.0.1','','','','','',NULL,'2025-03-17 09:07:59'),(45,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 10:50:49'),(46,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 10:51:01'),(47,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 11:00:44'),(48,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 11:02:52'),(49,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 11:05:16'),(50,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 11:05:25'),(51,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 11:12:11'),(52,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 11:15:19'),(53,2,3,'127.0.0.1','','','','','',NULL,'2025-03-17 20:31:44'),(54,2,3,'127.0.0.1','','','','','',NULL,'2025-03-18 14:42:14'),(55,2,3,'127.0.0.1','','','','','',NULL,'2025-03-19 17:34:43'),(56,2,3,'127.0.0.1','','','','','',NULL,'2025-03-20 11:31:21'),(57,8,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:17:42'),(58,8,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:17:51'),(59,8,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:48:06'),(60,9,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:54:35'),(61,9,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:57:06'),(62,8,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:57:43'),(63,9,3,'127.0.0.1','','','','','',NULL,'2025-03-20 21:58:19'),(64,8,3,'127.0.0.1','','','','','',NULL,'2025-03-20 22:10:33'),(65,9,3,'127.0.0.1','','','','','',NULL,'2025-03-21 09:01:41'),(66,8,3,'127.0.0.1','','','','','',NULL,'2025-03-22 09:31:56'),(67,9,3,'127.0.0.1','','','','','',NULL,'2025-03-22 10:10:59'),(68,8,3,'127.0.0.1','','','','','',NULL,'2025-03-22 10:21:53'),(69,8,3,'127.0.0.1','','','','','',NULL,'2025-03-22 10:27:35'),(70,8,3,'127.0.0.1','','','','','',NULL,'2025-03-22 10:30:30'),(71,8,3,'127.0.0.1','','','','','',NULL,'2025-03-22 10:32:22'),(72,8,3,'127.0.0.1','','','','','',NULL,'2025-03-22 10:32:29'),(73,8,3,'127.0.0.1','','','','','',NULL,'2025-03-23 17:43:10'),(74,8,3,'127.0.0.1','','','','','',NULL,'2025-03-23 17:44:35'),(75,8,3,'127.0.0.1','','','','','',NULL,'2025-03-23 17:48:23'),(76,8,3,'127.0.0.1','','','','','',NULL,'2025-03-23 17:51:16'),(77,8,3,'127.0.0.1','','','','','',NULL,'2025-03-24 09:52:08'),(78,8,3,'127.0.0.1','','','','','',NULL,'2025-03-24 13:42:26'),(79,8,3,'127.0.0.1','','','','','',NULL,'2025-03-24 15:56:43'),(80,8,3,'127.0.0.1','','','','','',0,'2025-03-26 09:53:13');
/*!40000 ALTER TABLE `sys_user_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_third_login`
--

DROP TABLE IF EXISTS `sys_user_third_login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_third_login` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `user_id` int unsigned DEFAULT NULL COMMENT '用户id',
  `platform` tinyint unsigned DEFAULT NULL COMMENT '平台 1 微信 2 钉钉',
  `open_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '第三方open_id',
  `union_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '第三方union_id',
  `third_data` json DEFAULT NULL COMMENT '第三方返回数据',
  `created_at` int unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` int unsigned DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sys_user_third_login_open_id_IDX` (`open_id`) USING BTREE,
  KEY `sys_user_third_login_union_id_IDX` (`union_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='三方登录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_third_login`
--

LOCK TABLES `sys_user_third_login` WRITE;
/*!40000 ALTER TABLE `sys_user_third_login` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_third_login` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team`
--

DROP TABLE IF EXISTS `team`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '团队名',
  `owner` int unsigned DEFAULT NULL COMMENT '团队拥有者',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `update_by` int DEFAULT NULL COMMENT '更新人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='团队';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team`
--

LOCK TABLES `team` WRITE;
/*!40000 ALTER TABLE `team` DISABLE KEYS */;
INSERT INTO `team` VALUES (1,'销售管理系统',2,3,NULL,'2023-09-30 00:22:43','2025-03-21 10:40:46'),(4,'大金 Team',8,3,0,'2025-03-20 21:48:06','2025-03-21 10:40:42'),(6,'小琪 Team',9,3,0,'2025-03-22 10:10:59','2025-03-22 10:10:59');
/*!40000 ALTER TABLE `team` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team_dept`
--

DROP TABLE IF EXISTS `team_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team_dept` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int unsigned DEFAULT NULL COMMENT '父id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门路径',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门名',
  `type` tinyint DEFAULT NULL COMMENT '类型 1分公司 2部门',
  `principal` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门领导',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `sort` tinyint DEFAULT NULL COMMENT '排序',
  `status` tinyint DEFAULT NULL COMMENT '状态 1正常 2关闭',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `team_id` int DEFAULT NULL COMMENT '团队id',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dept_deleted_at` (`deleted_at`) USING BTREE,
  KEY `team_dept_team_id_IDX` (`team_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='部门';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team_dept`
--

LOCK TABLES `team_dept` WRITE;
/*!40000 ALTER TABLE `team_dept` DISABLE KEYS */;
INSERT INTO `team_dept` VALUES (1,0,'/0/1/','销售一部',2,'糖糖','13800138000',NULL,1,1,NULL,1,NULL,2,NULL,'2023-10-26 20:57:48.321',NULL),(2,1,'/0/1/2/','糖糖组',2,'','',NULL,NULL,1,NULL,1,NULL,2,NULL,'2023-10-26 21:03:59.644',NULL),(9,10,'/0/10/9/','哈哈组',2,'哈哈','','',0,1,'1',1,2,1,'2023-10-26 20:50:45.790','2025-03-17 16:40:29.951',NULL),(10,0,'/0/10/','销售二部',0,'','','',0,1,'',1,2,0,'2023-10-26 20:51:11.338','2023-10-26 20:51:11.342',NULL),(11,0,'/0/11/','销售三部',1,'','','',0,1,'1',1,2,1,'2023-10-26 21:00:05.460','2025-03-17 16:40:16.720',NULL),(12,11,'/0/11/12/','呵呵组',2,'','','',0,-1,'1',1,2,1,'2023-10-26 21:00:22.635','2025-03-20 20:49:10.278',NULL),(13,0,'/0/13/','总裁办',2,'','','',0,1,'',4,8,0,'2025-03-20 21:49:57.945','2025-03-20 21:49:57.954',NULL),(14,0,'/0/14/','销售部',2,'','','',0,1,'',4,8,0,'2025-03-20 21:50:09.882','2025-03-20 21:50:09.895',NULL),(15,14,'/0/14/15/','销售一部',2,'','','',0,1,'',4,8,0,'2025-03-20 21:50:26.527','2025-03-20 21:50:26.537',NULL),(16,14,'/0/14/16/','销售二部',2,'','','',0,1,'',4,8,0,'2025-03-20 21:50:42.669','2025-03-20 21:50:42.679',NULL);
/*!40000 ALTER TABLE `team_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team_member`
--

DROP TABLE IF EXISTS `team_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team_member` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int DEFAULT NULL COMMENT '团队id',
  `user_id` int unsigned DEFAULT NULL COMMENT '用户id',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '昵称',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '姓名',
  `py` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '姓名拼音',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '电话',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '部门路径',
  `dept_id` int unsigned DEFAULT NULL COMMENT '部门id',
  `role_id` int DEFAULT NULL COMMENT '角色id',
  `post_id` tinyint unsigned DEFAULT NULL COMMENT '职位 1系统超管 2 团队拥有者 4主管 8副主管 16员工',
  `entry_time` datetime(3) DEFAULT NULL COMMENT '入职时间',
  `retire_time` datetime(3) DEFAULT NULL COMMENT '离职时间',
  `gender` tinyint(1) DEFAULT '2' COMMENT '性别 1男 2女 3未知',
  `birthday` date DEFAULT NULL COMMENT '生日 格式 yyyy-MM-dd',
  `status` tinyint DEFAULT NULL COMMENT '状态 1正常 ',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `team_member_team_id_IDX` (`team_id`,`user_id`) USING BTREE,
  KEY `team_member_dept_id_IDX` (`dept_id`) USING BTREE,
  KEY `team_member_dept_path_IDX` (`dept_path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='团队成员';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team_member`
--

LOCK TABLES `team_member` WRITE;
/*!40000 ALTER TABLE `team_member` DISABLE KEYS */;
INSERT INTO `team_member` VALUES (1,1,2,'糖糖','小唐','xiao-tang','','/0/1/2/',2,1,1,'2023-02-20 00:00:00.000',NULL,2,'1991-02-01',1,NULL,NULL,NULL,'2025-03-21 11:23:17.249'),(2,1,3,'梅梅','小梅','xiao-mei','13800138000','/0/1/',1,8,6,'2021-01-13 00:00:00.000',NULL,2,NULL,-1,NULL,1,NULL,'2025-03-21 11:23:10.910'),(12,1,4,'阿雷','阿雷','a-lei','','',0,0,6,NULL,NULL,1,NULL,1,0,0,'2025-03-20 15:18:10.954','2025-03-21 11:22:54.038'),(13,4,8,'大金','大金','da-jin','','',0,1,1,NULL,NULL,2,NULL,1,0,0,'2025-03-20 21:48:06.090','2025-03-20 21:48:06.090'),(14,4,9,'小琪','小琪','xiao-qi','','/0/14/16/',16,10,6,NULL,NULL,2,NULL,1,0,0,'2025-03-20 21:53:58.189','2025-03-20 21:58:07.568'),(16,6,9,'小琪','小琪','xiao-qi','','',0,1,1,NULL,NULL,2,NULL,1,0,0,'2025-03-22 10:10:59.469','2025-03-22 10:10:59.469');
/*!40000 ALTER TABLE `team_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team_role`
--

DROP TABLE IF EXISTS `team_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '角色名称',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '角色代码',
  `role_sort` int unsigned DEFAULT NULL COMMENT '排序',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `team_id` tinyint(1) DEFAULT NULL COMMENT '团队',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `create_by` int unsigned DEFAULT NULL COMMENT '创建者',
  `update_by` int unsigned DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_role_deleted_at` (`deleted_at`) USING BTREE,
  KEY `team_role_team_id_IDX` (`team_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team_role`
--

LOCK TABLES `team_role` WRITE;
/*!40000 ALTER TABLE `team_role` DISABLE KEYS */;
INSERT INTO `team_role` VALUES (1,'系统管理','superadmin',NULL,NULL,-1,NULL,NULL,NULL,NULL,NULL,NULL),(5,'test1','test1',0,0,1,'string',0,2,'2023-11-15 11:25:21.346','2023-11-15 16:31:21.931',NULL),(6,'test2','test2',0,0,1,'string',2,2,'2023-11-15 13:38:10.292','2023-11-15 16:31:07.397',NULL),(7,'test3','test3',0,1,1,'string',2,2,'2023-11-15 14:06:39.197','2025-03-21 11:40:11.521',NULL),(8,'角色1','test4',3,1,1,'',2,1,'2023-11-15 16:33:01.325','2025-03-17 16:37:28.283',NULL),(9,'总经理','',0,1,4,'',8,1,'2025-03-20 21:50:56.905','2025-03-21 11:21:00.180',NULL),(10,'销售','',0,1,4,'1',8,1,'2025-03-20 21:51:09.757','2025-03-21 11:15:26.008',NULL);
/*!40000 ALTER TABLE `team_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team_role_menu`
--

DROP TABLE IF EXISTS `team_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team_role_menu` (
  `role_id` int unsigned NOT NULL COMMENT '主键',
  `menu_id` int unsigned NOT NULL COMMENT '主键',
  PRIMARY KEY (`role_id`,`menu_id`) USING BTREE,
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`) USING BTREE,
  KEY `team_role_menu_role_id_IDX` (`role_id`) USING BTREE,
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `team_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='角色菜单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team_role_menu`
--

LOCK TABLES `team_role_menu` WRITE;
/*!40000 ALTER TABLE `team_role_menu` DISABLE KEYS */;
INSERT INTO `team_role_menu` VALUES (1,1),(5,1),(6,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8),(5,12),(6,12),(5,13),(6,13),(5,14),(5,15),(5,16),(5,17),(5,18),(5,19),(9,407),(10,407),(9,408),(10,408),(9,409),(10,409),(8,410),(9,410),(10,410),(8,411),(9,411),(10,411),(8,412),(9,412),(10,412),(7,413),(8,413),(9,413),(10,413),(7,414),(8,414),(9,414),(10,414),(7,415),(8,415),(9,415),(10,415),(7,416),(8,416),(9,416),(10,416),(7,417),(8,417),(9,417),(10,417),(7,418),(8,418),(9,418),(10,418),(7,419),(8,419),(9,419),(10,419),(9,420),(10,420),(9,421),(10,421),(9,422),(10,422),(9,423),(10,423),(9,424),(10,424);
/*!40000 ALTER TABLE `team_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'dilu-plus'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-26  9:54:18
