


DROP TABLE IF EXISTS `mfw_auth_apps`;
CREATE TABLE "mfw_auth_apps" (
  "id" int(10) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" timestamp NULL DEFAULT NULL,
  "updated_at" timestamp NULL DEFAULT NULL,
  "name" varchar(255) NOT NULL,
  "status" int(11) DEFAULT '1',
  "desc" varchar(255) DEFAULT NULL,
  "content" varchar(4096) DEFAULT NULL,
  "resource_id" int(10) unsigned DEFAULT NULL,
  PRIMARY KEY ("id"),
  UNIQUE KEY "name" ("name")
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `mfw_auth_auths`;
CREATE TABLE "mfw_auth_auths" (
  "id" int(10) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" timestamp NULL DEFAULT NULL,
  "updated_at" timestamp NULL DEFAULT NULL,
  "resource_id" int(10) unsigned DEFAULT NULL,
  "operation_id" int(11) DEFAULT NULL,
  PRIMARY KEY ("id"),
  KEY "mfw_auth_auths_resource_id_mfw_auth_resources_id_foreign" ("resource_id"),
  CONSTRAINT "mfw_auth_auths_resource_id_mfw_auth_resources_id_foreign" FOREIGN KEY ("resource_id") REFERENCES "mfw_auth_resources" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `mfw_auth_operations`;
CREATE TABLE "mfw_auth_operations" (
  "id" int(11) NOT NULL AUTO_INCREMENT,
  "name" varchar(255) NOT NULL,
  "value" varchar(255) NOT NULL,
  "status" int(11) DEFAULT '0',
  PRIMARY KEY ("id"),
  UNIQUE KEY "name" ("name"),
  UNIQUE KEY "value" ("value")
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `mfw_auth_resources`;
CREATE TABLE "mfw_auth_resources" (
  "id" int(10) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" timestamp NULL DEFAULT NULL,
  "updated_at" timestamp NULL DEFAULT NULL,
  "app_id" int(10) unsigned DEFAULT NULL,
  "name" varchar(255) DEFAULT NULL,
  "path" varchar(255) NOT NULL,
  PRIMARY KEY ("id"),
  UNIQUE KEY "path" ("path")
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `mfw_auth_user_auths`;
CREATE TABLE "mfw_auth_user_auths" (
  "id" int(10) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" timestamp NULL DEFAULT NULL,
  "updated_at" timestamp NULL DEFAULT NULL,
  "uid" varchar(255) DEFAULT NULL,
  "auth_id" int(10) unsigned DEFAULT NULL,
  PRIMARY KEY ("id"),
  KEY "mfw_auth_user_auths_auth_id_mfw_auth_auths_id_foreign" ("auth_id"),
  CONSTRAINT "mfw_auth_user_auths_auth_id_mfw_auth_auths_id_foreign" FOREIGN KEY ("auth_id") REFERENCES "mfw_auth_auths" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;