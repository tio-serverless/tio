CREATE SEQUENCE server_id
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS "tio"."server";
CREATE TABLE "tio"."server" (
  "id" int8 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "version" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "uid" int8 NOT NULL,
  "stype" int4 NOT NULL,
  "domain" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "Path" varchar(255) COLLATE "pg_catalog"."default",
  "tversion" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "timestamp" timestamp(0),
  "status" int4 NOT NULL,
  "raw" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "tio"."server" OWNER TO "tio";

-- ----------------------------
-- Primary Key structure for table server
-- ----------------------------
ALTER TABLE "tio"."server" ADD CONSTRAINT "server_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table server
-- ----------------------------
ALTER TABLE "tio"."server" ADD CONSTRAINT "user_id" FOREIGN KEY ("uid") REFERENCES "tio"."user" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "tio"."server" ALTER COLUMN  "id" set default nextval('server_id');