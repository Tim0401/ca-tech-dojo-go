---- drop ----
DROP TABLE IF EXISTS `user`;

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