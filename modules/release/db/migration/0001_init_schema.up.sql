CREATE TABLE "bounded_contexts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "manager" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bcTags" (
  "id" bigserial PRIMARY KEY,
  "bounded_context" bigint,
  "tag" bigint
);

CREATE TABLE "tags" (
  "id" bigserial PRIMARY KEY,
  "tag" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


ALTER TABLE "bcTags" ADD FOREIGN KEY ("bounded_context") REFERENCES "bounded_contexts" ("id");

ALTER TABLE "bcTags" ADD FOREIGN KEY ("tag") REFERENCES "tags" ("id");