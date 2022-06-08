CREATE TABLE IF NOT EXISTS "tags"(
  "id" VARCHAR UNIQUE NOT NULL, 
  "name" VARCHAR UNIQUE NOT NULL, 
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL, 
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL, 
  PRIMARY KEY("id")
);
