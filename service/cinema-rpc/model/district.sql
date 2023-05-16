CREATE TABLE district(
     id BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键编号',
     city_name VARCHAR(100) NOT NULL DEFAULT '' COMMENT '市名',
     district_name VARCHAR(100) NOT NULL DEFAULT '' COMMENT '区名',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT ='地域信息表';

INSERT INTO district(id,city_name,district_name) VALUES(1,'北京市','朝阳区');
INSERT INTO district(id,city_name,district_name) VALUES(2,'北京市','海淀区');
INSERT INTO district(id,city_name,district_name) VALUES(3,'北京市','丰台区');
INSERT INTO district(id,city_name,district_name) VALUES(4,'北京市','大兴区');
INSERT INTO district(id,city_name,district_name) VALUES(5,'北京市','东城区');
INSERT INTO district(id,city_name,district_name) VALUES(6,'北京市','西城区');
INSERT INTO district(id,city_name,district_name) VALUES(7,'北京市','通州区');
INSERT INTO district(id,city_name,district_name) VALUES(8,'北京市','房山区');
INSERT INTO district(id,city_name,district_name) VALUES(9,'北京市','昌平区');
INSERT INTO district(id,city_name,district_name) VALUES(10,'北京市','顺义区');
INSERT INTO district(id,city_name,district_name) VALUES(11,'北京市','怀柔区');
INSERT INTO district(id,city_name,district_name) VALUES(12,'北京市','门头沟');
INSERT INTO district(id,city_name,district_name) VALUES(13,'北京市','石景山区');
INSERT INTO district(id,city_name,district_name) VALUES(14,'北京市','密云区');
INSERT INTO district(id,city_name,district_name) VALUES(15,'北京市','平谷区');
INSERT INTO district(id,city_name,district_name) VALUES(16,'北京市','延庆区');