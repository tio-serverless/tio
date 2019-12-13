CREATE SEQUENCE user_id
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;


-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "tio"."tio-user";
CREATE TABLE "tio"."tio-user" (
  "id" int8 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "passwd" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "tio"."tio-user" OWNER TO "tio";

-- ----------------------------
-- Uniques structure for table user
-- ----------------------------
ALTER TABLE "tio"."tio-user" ADD CONSTRAINT "user_pname" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "tio"."tio-user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
ALTER TABLE "tio"."tio-user" ALTER COLUMN  "id" set default nextval('user_id');

