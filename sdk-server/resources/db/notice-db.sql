-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: notice-db
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
-- Table structure for table `ai_msg`
--

DROP TABLE IF EXISTS `ai_msg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ai_msg` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `platform` tinyint DEFAULT NULL COMMENT '平台',
  `model` tinyint DEFAULT NULL COMMENT '模型',
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `topic_id` int DEFAULT NULL COMMENT '话题',
  `is_user` tinyint DEFAULT NULL COMMENT '用户',
  `text` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文本',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ai_msg`
--

LOCK TABLES `ai_msg` WRITE;
/*!40000 ALTER TABLE `ai_msg` DISABLE KEYS */;
/*!40000 ALTER TABLE `ai_msg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ai_topic`
--

DROP TABLE IF EXISTS `ai_topic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ai_topic` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `platform` tinyint DEFAULT NULL COMMENT '平台',
  `model` tinyint DEFAULT NULL COMMENT '模型',
  `topic` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '话题',
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ai_topic`
--

LOCK TABLES `ai_topic` WRITE;
/*!40000 ALTER TABLE `ai_topic` DISABLE KEYS */;
/*!40000 ALTER TABLE `ai_topic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pub_notice`
--

DROP TABLE IF EXISTS `pub_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pub_notice` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int DEFAULT '0' COMMENT '针对组消息',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '内容',
  `notice_type` tinyint DEFAULT NULL COMMENT '消息类型',
  `op` tinyint DEFAULT NULL COMMENT '操作类型',
  `op_id` int DEFAULT NULL COMMENT '操作id',
  `status` tinyint DEFAULT NULL COMMENT '状态',
  `create_by` int DEFAULT NULL COMMENT '创建人',
  `update_by` int DEFAULT NULL COMMENT '更新人',
  `expired` datetime DEFAULT NULL COMMENT '到期时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `pub_notice_team_id_IDX` (`team_id`,`expired` DESC) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='公用通知';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pub_notice`
--

LOCK TABLES `pub_notice` WRITE;
/*!40000 ALTER TABLE `pub_notice` DISABLE KEYS */;
INSERT INTO `pub_notice` VALUES (1,0,'test','testcontent',NULL,NULL,NULL,1,1,1,'2023-12-07 14:40:06','2023-11-24 14:40:13','2023-11-24 14:40:17'),(2,1,'test2','content',NULL,NULL,NULL,1,1,1,'2023-11-25 14:53:05','2023-11-24 14:53:14','2023-11-24 14:53:18'),(3,0,'地方大师傅','的发射点',0,0,0,1,0,0,'2023-11-28 05:02:02','2023-11-26 09:12:03','2023-11-26 13:01:20'),(4,0,'活动啊','活动通知一下',0,0,0,1,0,0,'2023-11-30 00:00:00','2023-11-26 13:01:11','2023-11-26 13:01:11');
/*!40000 ALTER TABLE `pub_notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int DEFAULT NULL COMMENT '团队id',
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '任务标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '任务内容',
  `task_type` tinyint DEFAULT NULL COMMENT '任务类型',
  `op` int DEFAULT NULL COMMENT '操作类型',
  `op_id` int DEFAULT NULL COMMENT '操作id',
  `begin_at` datetime DEFAULT NULL COMMENT '开始时间',
  `end_at` datetime DEFAULT NULL COMMENT '结束时间',
  `reminder_time` datetime DEFAULT NULL COMMENT '提醒时间',
  `status` tinyint DEFAULT NULL COMMENT '状态1开启2关闭',
  `reminder_status` tinyint DEFAULT NULL COMMENT '提醒状态 1开启 2关闭',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task`
--

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
INSERT INTO `task` VALUES (1,1,2,'嘻嘻嘻','顶顶顶顶顶',0,0,0,'2023-11-26 15:04:02','2023-11-27 00:00:00','2023-11-26 13:02:03',1,1,'2023-11-26 13:02:08','2023-11-26 13:02:08',NULL),(2,-1,1,'这是一个任务','任务开始了',0,0,0,'2023-11-26 19:52:36','2023-11-26 19:52:38','2023-11-26 19:52:40',1,1,'2023-11-26 19:52:47','2023-11-26 19:52:47',NULL);
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_notice`
--

DROP TABLE IF EXISTS `user_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_notice` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int DEFAULT NULL COMMENT '团队id',
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '内容',
  `notice_type` tinyint DEFAULT NULL COMMENT '消息类型',
  `op` tinyint DEFAULT NULL COMMENT '操作类型',
  `op_id` int DEFAULT NULL COMMENT '操作对象id',
  `status` tinyint DEFAULT NULL COMMENT '状态 1未读 2已读 -1回收站',
  `pub_id` int DEFAULT NULL COMMENT '公共id',
  `create_by` int DEFAULT NULL COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_notice_team_id_IDX` (`team_id`,`user_id`,`pub_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户通知';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_notice`
--

LOCK TABLES `user_notice` WRITE;
/*!40000 ALTER TABLE `user_notice` DISABLE KEYS */;
INSERT INTO `user_notice` VALUES (2,1,2,'test2','content',0,0,0,2,2,1,'2023-11-24 14:53:14','2025-03-21 14:49:48',NULL),(3,1,2,'test','testcontent',0,0,0,2,1,1,'2023-11-24 14:40:13','2025-03-21 14:49:49',NULL),(4,1,2,'test3','33333',0,0,0,2,NULL,1,'2023-11-24 14:55:59','2025-03-21 14:49:46',NULL),(5,-1,1,'活动啊','活动通知一下',0,0,0,2,4,0,'2023-11-26 13:01:11','2025-03-21 15:04:15',NULL),(6,-1,1,'地方大师傅','的发射点',0,0,0,2,3,0,'2023-11-26 09:12:03','2025-03-21 15:04:16',NULL),(7,-1,1,'test','testcontent',0,0,0,2,1,1,'2023-11-24 14:40:13','2025-03-21 15:04:17',NULL);
/*!40000 ALTER TABLE `user_notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'notice-db'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-24 14:18:53
