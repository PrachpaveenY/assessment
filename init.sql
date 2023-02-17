-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS news-tables-id;

-- Table Definition
CREATE TABLE "news-tables" (
    "id" int5 NOT NULL DEFAULT nextval('news-tables-id'::regclass),
    "title" text,
    "amount" FLOAT,
	"note" TEXT,
	"tags" TEXT[],
    PRIMARY KEY ("id")
);

INSERT INTO "news-tables" ("id", "title", "amount", "note", "tags") VALUES (1, 'test-title', 'test-amount', 'test-note', "tags");

-- ====================

