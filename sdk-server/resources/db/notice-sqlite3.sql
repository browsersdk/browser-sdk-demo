-- SQLite3 compatible SQL dump
-- Converted from MySQL dump for notice-db database

PRAGMA foreign_keys = OFF;

-- Table structure for table pub_notice
DROP TABLE IF EXISTS pub_notice;
CREATE TABLE pub_notice (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  team_id INTEGER DEFAULT 0,
  title TEXT,
  content TEXT,
  notice_type INTEGER,
  op INTEGER,
  op_id INTEGER,
  status INTEGER,
  create_by INTEGER,
  update_by INTEGER,
  expired DATETIME,
  created_at DATETIME,
  updated_at DATETIME
);

-- Dumping data for table pub_notice
INSERT INTO pub_notice VALUES (1,0,'test','testcontent',NULL,NULL,NULL,1,1,1,'2023-12-07 14:40:06','2023-11-24 14:40:13','2023-11-24 14:40:17');
INSERT INTO pub_notice VALUES (2,1,'test2','content',NULL,NULL,NULL,1,1,1,'2023-11-25 14:53:05','2023-11-24 14:53:14','2023-11-24 14:53:18');
INSERT INTO pub_notice VALUES (3,0,'地方大师傅','的发射点',0,0,0,1,0,0,'2023-11-28 05:02:02','2023-11-26 09:12:03','2023-11-26 13:01:20');
INSERT INTO pub_notice VALUES (4,0,'活动啊','活动通知一下',0,0,0,1,0,0,'2023-11-30 00:00:00','2023-11-26 13:01:11','2023-11-26 13:01:11');

-- Table structure for table task
DROP TABLE IF EXISTS task;
CREATE TABLE task (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  team_id INTEGER,
  user_id INTEGER,
  title TEXT,
  content TEXT,
  task_type INTEGER,
  op INTEGER,
  op_id INTEGER,
  begin_at DATETIME,
  end_at DATETIME,
  reminder_time DATETIME,
  status INTEGER,
  reminder_status INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table task
INSERT INTO task VALUES (1,1,2,'嘻嘻嘻','顶顶顶顶顶',0,0,0,'2023-11-26 15:04:02','2023-11-27 00:00:00','2023-11-26 13:02:03',1,1,'2023-11-26 13:02:08','2023-11-26 13:02:08',NULL);
INSERT INTO task VALUES (2,-1,1,'这是一个任务','任务开始了',0,0,0,'2023-11-26 19:52:36','2023-11-26 19:52:38','2023-11-26 19:52:40',1,1,'2023-11-26 19:52:47','2023-11-26 19:52:47',NULL);

-- Table structure for table user_notice
DROP TABLE IF EXISTS user_notice;
CREATE TABLE user_notice (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  team_id INTEGER,
  user_id INTEGER,
  title TEXT,
  content TEXT,
  notice_type INTEGER,
  op INTEGER,
  op_id INTEGER,
  status INTEGER,
  pub_id INTEGER,
  create_by INTEGER,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME
);

-- Dumping data for table user_notice
INSERT INTO user_notice VALUES (2,1,2,'test2','content',0,0,0,2,2,1,'2023-11-24 14:53:14','2025-03-21 14:49:48',NULL);
INSERT INTO user_notice VALUES (3,1,2,'test','testcontent',0,0,0,2,1,1,'2023-11-24 14:40:13','2025-03-21 14:49:49',NULL);
INSERT INTO user_notice VALUES (4,1,2,'test3','33333',0,0,0,2,NULL,1,'2023-11-24 14:55:59','2025-03-21 14:49:46',NULL);
INSERT INTO user_notice VALUES (5,-1,1,'活动啊','活动通知一下',0,0,0,2,4,0,'2023-11-26 13:01:11','2025-03-21 15:04:15',NULL);
INSERT INTO user_notice VALUES (6,-1,1,'地方大师傅','的发射点',0,0,0,2,3,0,'2023-11-26 09:12:03','2025-03-21 15:04:16',NULL);
INSERT INTO user_notice VALUES (7,-1,1,'test','testcontent',0,0,0,2,1,1,'2023-11-24 14:40:13','2025-03-21 15:04:17',NULL);

PRAGMA foreign_keys = ON;