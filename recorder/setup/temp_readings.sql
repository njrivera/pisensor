BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `readings` (
	`serial`	TEXT NOT NULL,
	`model`	TEXT NOT NULL,
	`temp`	REAL NOT NULL,
	`unit`	TEXT NOT NULL,
	`timestamp`	INTEGER NOT NULL,
	PRIMARY KEY(`serial`,`timestamp`)
);
COMMIT;
