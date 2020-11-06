---- drop ----
DROP TABLE IF EXISTS `chara_user`;
DROP TABLE IF EXISTS `gacha`;
DROP TABLE IF EXISTS `gacha_type_rate_type`;
DROP TABLE IF EXISTS `rate_type`;
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

create UNIQUE INDEX chara_user_index on chara_user(user_id, chara_id);
ALTER TABLE chara_user add FOREIGN KEY (user_id) references user(id);
ALTER TABLE chara_user add FOREIGN KEY (chara_id) references chara(id);

create table IF not exists `gacha_type`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャタイプ';

create table IF not exists `rate_type`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `rate`             INT(20) NOT NULL comment 'タイプごとの排出率(重み)',
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャ排出率タイプ';

create table IF not exists `gacha_type_rate_type`
(
 `id`               INT(20) AUTO_INCREMENT,
 `gacha_type_id`    INT(20) NOT NULL,
 `rate_type_id`     INT(20) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャタイプで使用する排出率タイプ';

create UNIQUE INDEX gacha_type_rate_type_index on gacha_type_rate_type(gacha_type_id, rate_type_id);
ALTER TABLE gacha_type_rate_type add FOREIGN KEY (gacha_type_id) references gacha_type(id);
ALTER TABLE gacha_type_rate_type add FOREIGN KEY (rate_type_id) references rate_type(id);

create table IF not exists `gacha`
(
 `id`               INT(20) AUTO_INCREMENT,
 `chara_id`         INT(20) NOT NULL,
 `gacha_type_id`    INT(20) NOT NULL,
 `rate_type_id`     INT(20) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='ガチャ排出キャラ';

create UNIQUE INDEX gacha_index on gacha(chara_id, gacha_type_id);
ALTER TABLE gacha add FOREIGN KEY (chara_id) references chara(id);
ALTER TABLE gacha add FOREIGN KEY (gacha_type_id) references gacha_type(id);
ALTER TABLE gacha add FOREIGN KEY (rate_type_id) references rate_type(id);
ALTER TABLE gacha add FOREIGN KEY (gacha_type_id, rate_type_id) references gacha_type_rate_type(gacha_type_id, rate_type_id);