CREATE TABLE "images" (
  "id" uuid PRIMARY KEY NOT NULL,
  "post_id" uuid NOT NULL,
  "img_url" varchar(1024) NOT NULL
);

ALTER TABLE "images" ADD CONSTRAINT post_images FOREIGN KEY ("post_id") REFERENCES "post" ("id");
