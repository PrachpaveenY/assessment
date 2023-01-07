-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS news_tables_id_seq;

-- Table Definition
CREATE TABLE "news_tables" (
    "id" int5 NOT NULL DEFAULT nextval('news_tables_id_seq'::regclass),
    "title" text,
    "amount" FLOAT,
	"note" TEXT,
	"tags" TEXT[],
    PRIMARY KEY ("id")
);

INSERT INTO "news_tables" ("id", "title", "amount", "note", "tags") VALUES (1, 'test-title', 'test-amount', 'test-note', "tags");

-- ====================

