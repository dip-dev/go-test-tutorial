CREATE DATABASE IF NOT EXISTS TUTORIAL_USER;

CREATE USER IF NOT EXISTS 'TUTORIAL_USER'@'%' IDENTIFIED BY 'GO_TEST_TUTORIAL';
GRANT ALL ON *.* TO 'TUTORIAL_USER'@'%';

CREATE TABLE IF NOT EXISTS TUTORIAL_USER.m_prefecture(
  id        INT(11) PRIMARY KEY NOT NULL,
  `name`    VARCHAR(10) NOT NULL,
  area      VARCHAR(10) NOT NULL,
  land_area INT(11) NOT NULL,
  UNIQUE (`name`)
);

INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (1,'北海道','北海道',83424);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (2,'青森県','東北',9646);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (3,'岩手県','東北',15275);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (4,'宮城県','東北',7282);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (5,'秋田県','東北',11638);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (6,'山形県','東北',9323);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (7,'福島県','東北',13784);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (8,'茨城県','関東',6097);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (9,'栃木県','関東',6408);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (10,'群馬県','関東',6362);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (11,'埼玉県','関東',3798);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (12,'千葉県','関東',5158);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (13,'東京都','関東',2194);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (14,'神奈川県','関東',2416);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (15,'新潟県','中部',12584);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (16,'富山県','中部',4248);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (17,'石川県','中部',4186);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (18,'福井県','中部',4191);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (19,'山梨県','中部',4465);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (20,'長野県','中部',13562);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (21,'岐阜県','中部',10621);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (22,'静岡県','中部',7777);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (23,'愛知県','中部',5173);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (24,'三重県','近畿',5774);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (25,'滋賀県','近畿',4017);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (26,'京都府','近畿',4612);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (27,'大阪府','近畿',1905);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (28,'兵庫県','近畿',8401);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (29,'奈良県','近畿',3691);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (30,'和歌山県','近畿',4725);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (31,'鳥取県','中国',3507);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (32,'島根県','中国',6708);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (33,'岡山県','中国',7114);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (34,'広島県','中国',8480);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (35,'山口県','中国',6113);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (36,'徳島県','四国',4147);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (37,'香川県','四国',1877);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (38,'愛媛県','四国',5676);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (39,'高知県','四国',7104);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (40,'福岡県','九州',4987);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (41,'佐賀県','九州',2441);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (42,'長崎県','九州',4131);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (43,'熊本県','九州',7409);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (44,'大分県','九州',6341);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (45,'宮崎県','九州',7735);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (46,'鹿児島県','九州',9187);
INSERT INTO TUTORIAL_USER.m_prefecture (id,`name`,area,land_area) VALUES (47,'沖縄県','九州',2283);
