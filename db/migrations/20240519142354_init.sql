-- Create "email" table
CREATE TABLE "email" ("recipient" character varying(255) NOT NULL, "subject" character varying(255) NOT NULL, "send_date" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "uuid" uuid NOT NULL DEFAULT gen_random_uuid(), PRIMARY KEY ("uuid"));
-- Create "tracker" table
CREATE TABLE "tracker" ("tracker_id" serial NOT NULL, "open_date" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "email_uuid" uuid NOT NULL, "ip_address" character varying(255) NULL, PRIMARY KEY ("tracker_id"), CONSTRAINT "tracker_email_uuid_fkey" FOREIGN KEY ("email_uuid") REFERENCES "email" ("uuid") ON UPDATE NO ACTION ON DELETE NO ACTION);
