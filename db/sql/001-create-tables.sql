---- drop ----
DROP TABLE IF EXISTS `chara_user`;
DROP TABLE IF EXISTS `gacha_probability_group`;
DROP TABLE IF EXISTS `gacha_probability`;
DROP TABLE IF EXISTS `gacha_type`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `chara`;

---- create ----
create table IF not exists `user`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `token`            VARCHAR(255) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ユーザー一覧';

create UNIQUE INDEX token_index on user(token);

create table IF not exists `chara`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(127) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='キャラ一覧';

create table IF not exists `chara_user`
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `chara_id`         INT(20) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ユーザー所持キャラ';

ALTER TABLE chara_user add FOREIGN KEY (user_id) references user(id);
ALTER TABLE chara_user add FOREIGN KEY (chara_id) references chara(id);

-- ガチャ種類
CREATE TABLE `gacha_type` (
  `id`              INT(20) AUTO_INCREMENT,
  `name`            VARCHAR(255),
  `created_at`      Datetime NOT NULL,
  `updated_at`      Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャ種類';

-- gacha_probability_group: gacha_probability_group_id=normal: 1
-- gacha_probability_group: gacha_probability_group_id=rare: 1
-- gacha_probability_group: gacha_probability_group_id=s_rare: 1
-- gacha_probability_group: gacha_probability_group_id=ss_rare: 1
-- レアリティの排出確率
CREATE TABLE `gacha_probability_group` (
  `gacha_type_id`              INT(20),
  `gacha_probability_group_id` VARCHAR(255),
  `rate`                       INT(20) NOT NULL DEFAULT 1,
  `created_at`                 Datetime NOT NULL,
  `updated_at`                 Datetime DEFAULT NULL,
  PRIMARY KEY (`gacha_type_id`, `gacha_probability_group_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャ確率グループ';

-- gacha_probability: chara_id=hoge: 1
-- gacha_probability: chara_id=fuga: 1
-- レアリティごとの排出内容
CREATE TABLE `gacha_probability` (
  `group_id`    VARCHAR(255), -- normal
  `number`      INT(20), -- 1,2,3
  `chara_id`    INT(20) NOT NULL,
  `rate`        INT(20) NOT NULL DEFAULT 1,
  `created_at`  Datetime NOT NULL,
  `updated_at`  Datetime DEFAULT NULL,
  PRIMARY KEY (`group_id`, `number`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャ確率';

ALTER TABLE gacha_probability_group add FOREIGN KEY (gacha_probability_group_id) references gacha_probability(group_id);
ALTER TABLE gacha_probability_group add FOREIGN KEY (gacha_type_id) references gacha_type(id);
ALTER TABLE gacha_probability add FOREIGN KEY (chara_id) references chara(id);