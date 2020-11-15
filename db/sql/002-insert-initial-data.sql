---- user ----
INSERT INTO user (name, token, created_at) VALUES ("testuser1", "testtoken1", NOW());
INSERT INTO user (name, token, created_at) VALUES ("testuser2", "testtoken2", NOW());
INSERT INTO user (name, token, created_at) VALUES ("testuser3", "testtoken3", NOW());
INSERT INTO user (name, token, created_at) VALUES ("testuser4", "testtoken4", NOW());
INSERT INTO user (name, token, created_at) VALUES ("testuser5", "testtoken5", NOW());

---- chara ----
INSERT INTO chara (name, created_at) VALUES ("testchara1", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara2", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara3", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara4", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara5", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara6", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara7", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara8", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara9", NOW());
INSERT INTO chara (name, created_at) VALUES ("testchara10", NOW());

---- gacha_type ----
INSERT INTO gacha_type (name, created_at) VALUES ("platinum", NOW());
INSERT INTO gacha_type (name, created_at) VALUES ("pickup", NOW());

---- gacha_probability ----
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★1", 1, 1, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★1", 2, 2, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★1", 3, 3, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★1", 4, 4, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★1", 5, 5, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★2", 1, 6, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★2", 2, 7, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★2", 3, 8, NOW());

INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★3", 1, 9, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★3", 2, 10, NOW());

INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★3PU", 1, 9, NOW());
INSERT INTO gacha_probability (group_id, number, chara_id, created_at) VALUES ("★3NOPU", 1, 10, NOW());

---- gacha_probability_group ----
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (1, "★1", 795, NOW());
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (1, "★2", 180, NOW());
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (1, "★3", 25, NOW());

---- gacha_probability_group ----
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (2, "★1", 795, NOW());
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (2, "★2", 180, NOW());
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (2, "★3PU", 7, NOW());
INSERT INTO gacha_probability_group (gacha_type_id, gacha_probability_group_id, rate, created_at) VALUES (2, "★3NOPU", 18, NOW());
