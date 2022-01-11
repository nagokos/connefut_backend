CREATE TABLE IF NOT EXISTS "users"(
  "id" varchar UNIQUE NOT NULL, 
  "name" varchar(50) NOT NULL, 
  "email" varchar(100) UNIQUE NOT NULL, 
  "role" varchar NOT NULL DEFAULT 'general', 
  "avatar" varchar NOT NULL DEFAULT 'https://abs.twimg.com/sticky/default_profile_images/default_profile.png', 
  "introduction" varchar(4000) NULL, 
  "email_verification_status" boolean NOT NULL DEFAULT false, 
  "email_verification_token" varchar UNIQUE NULL, 
  "email_verification_token_expires_at" timestamp with time zone NULL, 
  "password_digest" varchar NULL,
  "created_at" timestamp with time zone NOT NULL, 
  "updated_at" timestamp with time zone NOT NULL, 
  PRIMARY KEY("id")
);