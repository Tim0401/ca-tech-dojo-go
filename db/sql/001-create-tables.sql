---- drop ----
DROP TABLE IF EXISTS `chara_user`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `character`;

---- create ----
create table IF not exists `user`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `token`            VARCHAR(255) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create UNIQUE INDEX token_index on user(token);

create table IF not exists `chara`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(127) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

create table IF not exists `chara_user`
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `chara_id`     INT(20) NOT NULL,
 `created_at`       Datetime NOT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE chara_user add FOREIGN KEY (user_id) references user(id);
ALTER TABLE chara_user add FOREIGN KEY (chara_id) references chara(id);

create UNIQUE INDEX chara_user_index on chara_user(user_id, chara_id);