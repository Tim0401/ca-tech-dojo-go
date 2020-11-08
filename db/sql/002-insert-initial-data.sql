---- user ----
INSERT INTO user (name, token, created_at) VALUES ("testuser1", "testtoken1", NOW());
INSERT INTO user (name, token, created_at) VALUES ("testuser2", "testtoken2", NOW());

---- chara ----
INSERT INTO chara (name, created_at) VALUES ("testchara1", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara2", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara3", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara4", NOW());


---- gacha_type ----
INSERT INTO gacha_type (name, created_at) VALUES ("プラチナ", NOW());
INSERT INTO gacha_type (name, created_at) VALUES ("ピックアップ", NOW());

---- rate_type ----
INSERT INTO rate_type (name, rate, created_at) VALUES ("★3",  25, NOW());
INSERT INTO rate_type (name, rate, created_at) VALUES ("★2", 180, NOW());
INSERT INTO rate_type (name, rate, created_at) VALUES ("★1", 795, NOW());

INSERT INTO rate_type (name, rate, created_at) VALUES ("★3",  18, NOW());
INSERT INTO rate_type (name, rate, created_at) VALUES ("★3PU", 7, NOW());
INSERT INTO rate_type (name, rate, created_at) VALUES ("★2", 180, NOW());
INSERT INTO rate_type (name, rate, created_at) VALUES ("★1", 795, NOW());


---- gacha_type_rate_type ----
INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (1, 1, NOW());
INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (1, 2, NOW());
INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (1, 3, NOW());

INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (2, 4, NOW());
INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (2, 5, NOW());
INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (2, 6, NOW());
INSERT INTO gacha_type_rate_type (gacha_type_id, rate_type_id, created_at) VALUES (2, 7, NOW());

---- gacha ----
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (1, 1, 1, NOW());
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (2, 1, 2, NOW());
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (3, 1, 3, NOW());
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (4, 1, 3, NOW());

INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (1, 2, 4, NOW());
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (2, 2, 5, NOW());
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (3, 2, 6, NOW());
INSERT INTO gacha (chara_id, gacha_type_id, rate_type_id, created_at) VALUES (4, 2, 7, NOW());