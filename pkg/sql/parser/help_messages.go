// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1058
	`ALTER`: {
		//line sql.y: 1059
		Category: hGroup,
		//line sql.y: 1060
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1074
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1075
		Category: hDDL,
		//line sql.y: 1076
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>
  ALTER PARTITION ... OF TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1112
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1126
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1127
		Category: hDDL,
		//line sql.y: 1128
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1130
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1137
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1138
		Category: hDDL,
		//line sql.y: 1139
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1172
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1173
		Category: hPriv,
		//line sql.y: 1174
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1176
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1181
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1182
		Category: hDDL,
		//line sql.y: 1183
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1185
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1193
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1194
		Category: hDDL,
		//line sql.y: 1195
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1207
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1212
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1213
		Category: hDDL,
		//line sql.y: 1214
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1222
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1649
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1650
		Category: hCCL,
		//line sql.y: 1651
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1668
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1676
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1677
		Category: hCCL,
		//line sql.y: 1678
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1694
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1712
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1713
		Category: hCCL,
		//line sql.y: 1714
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1742
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1792
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1793
		Category: hCCL,
		//line sql.y: 1794
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1803
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1900
	`CANCEL`: {
		//line sql.y: 1901
		Category: hGroup,
		//line sql.y: 1902
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1909
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1910
		Category: hMisc,
		//line sql.y: 1911
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1914
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1932
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1933
		Category: hMisc,
		//line sql.y: 1934
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1937
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1968
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1969
		Category: hMisc,
		//line sql.y: 1970
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1973
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2050
	`CREATE`: {
		//line sql.y: 2051
		Category: hGroup,
		//line sql.y: 2052
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2133
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic (experimental)`,
		//line sql.y: 2134
		Category: hExperimental,
		//line sql.y: 2135
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2221
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2222
		Category: hDML,
		//line sql.y: 2223
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2227
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2242
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2243
		Category: hCfg,
		//line sql.y: 2244
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2256
	`DROP`: {
		//line sql.y: 2257
		Category: hGroup,
		//line sql.y: 2258
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2275
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2276
		Category: hDDL,
		//line sql.y: 2277
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2278
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2290
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2291
		Category: hDDL,
		//line sql.y: 2292
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2293
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2305
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2306
		Category: hDDL,
		//line sql.y: 2307
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2308
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2320
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2321
		Category: hDDL,
		//line sql.y: 2322
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2323
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2343
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2344
		Category: hDDL,
		//line sql.y: 2345
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2346
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2366
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2367
		Category: hPriv,
		//line sql.y: 2368
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2369
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2381
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2382
		Category: hPriv,
		//line sql.y: 2383
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2384
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2416
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2417
		Category: hMisc,
		//line sql.y: 2418
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2431
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2491
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2492
		Category: hMisc,
		//line sql.y: 2493
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2494
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2516
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2517
		Category: hMisc,
		//line sql.y: 2518
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2519
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2540
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2541
		Category: hMisc,
		//line sql.y: 2542
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2543
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2563
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2564
		Category: hPriv,
		//line sql.y: 2565
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2578
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2594
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2595
		Category: hPriv,
		//line sql.y: 2596
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2609
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2664
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2665
		Category: hCfg,
		//line sql.y: 2666
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2667
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2679
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2680
		Category: hCfg,
		//line sql.y: 2681
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2682
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2691
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2692
		Category: hCfg,
		//line sql.y: 2693
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2696
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2717
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2718
		Category: hExperimental,
		//line sql.y: 2719
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2727
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2733
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2734
		Category: hExperimental,
		//line sql.y: 2735
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2743
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2751
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2752
		Category: hExperimental,
		//line sql.y: 2753
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 2764
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2824
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2825
		Category: hCfg,
		//line sql.y: 2826
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2827
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2848
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2849
		Category: hCfg,
		//line sql.y: 2850
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2856
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2873
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2874
		Category: hTxn,
		//line sql.y: 2875
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2882
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3065
	`SHOW`: {
		//line sql.y: 3066
		Category: hGroup,
		//line sql.y: 3067
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW JOBS,
SHOW QUERIES, SHOW ROLES, SHOW SESSION, SHOW SESSIONS, SHOW STATISTICS,
SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3099
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3100
		Category: hCfg,
		//line sql.y: 3101
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3102
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3123
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3124
		Category: hExperimental,
		//line sql.y: 3125
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3132
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3155
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3156
		Category: hExperimental,
		//line sql.y: 3157
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3161
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3175
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3176
		Category: hCCL,
		//line sql.y: 3177
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 3178
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3205
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3206
		Category: hCfg,
		//line sql.y: 3207
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 3210
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3227
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3228
		Category: hDDL,
		//line sql.y: 3229
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3230
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3243
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3244
		Category: hDDL,
		//line sql.y: 3245
		Text: `SHOW DATABASES
`,
		//line sql.y: 3246
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3254
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3255
		Category: hPriv,
		//line sql.y: 3256
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3262
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3275
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3276
		Category: hDDL,
		//line sql.y: 3277
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3278
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3311
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3312
		Category: hDDL,
		//line sql.y: 3313
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3314
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3337
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3338
		Category: hMisc,
		//line sql.y: 3339
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3340
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3356
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3357
		Category: hMisc,
		//line sql.y: 3358
		Text: `SHOW JOBS
`,
		//line sql.y: 3359
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3367
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3368
		Category: hMisc,
		//line sql.y: 3369
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3371
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3394
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3395
		Category: hMisc,
		//line sql.y: 3396
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3397
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3413
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3414
		Category: hDDL,
		//line sql.y: 3415
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3416
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3448
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3449
		Category: hDDL,
		//line sql.y: 3450
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3462
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3463
		Category: hMisc,
		//line sql.y: 3464
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3473
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3474
		Category: hCfg,
		//line sql.y: 3475
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3476
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3495
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3496
		Category: hDDL,
		//line sql.y: 3497
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3498
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3526
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3527
		Category: hPriv,
		//line sql.y: 3528
		Text: `SHOW USERS
`,
		//line sql.y: 3529
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3537
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3538
		Category: hPriv,
		//line sql.y: 3539
		Text: `SHOW ROLES
`,
		//line sql.y: 3540
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3595
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3596
		Category: hMisc,
		//line sql.y: 3597
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3843
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3844
		Category: hMisc,
		//line sql.y: 3845
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3848
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3865
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3866
		Category: hDDL,
		//line sql.y: 3867
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3894
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4510
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4511
		Category: hDDL,
		//line sql.y: 4512
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4522
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4577
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4578
		Category: hDML,
		//line sql.y: 4579
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4580
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4588
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4589
		Category: hPriv,
		//line sql.y: 4590
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4591
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4613
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4614
		Category: hPriv,
		//line sql.y: 4615
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4616
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4634
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4635
		Category: hDDL,
		//line sql.y: 4636
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4637
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4675
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4676
		Category: hDDL,
		//line sql.y: 4677
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4685
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4999
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5000
		Category: hTxn,
		//line sql.y: 5001
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5002
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5010
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5011
		Category: hMisc,
		//line sql.y: 5012
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5015
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5032
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5033
		Category: hTxn,
		//line sql.y: 5034
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5035
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5050
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5051
		Category: hTxn,
		//line sql.y: 5052
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5060
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5073
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5074
		Category: hTxn,
		//line sql.y: 5075
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5078
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5102
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5103
		Category: hTxn,
		//line sql.y: 5104
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5105
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5219
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5220
		Category: hDDL,
		//line sql.y: 5221
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5222
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5291
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5292
		Category: hDML,
		//line sql.y: 5293
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5298
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5317
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5318
		Category: hDML,
		//line sql.y: 5319
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5323
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5438
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5439
		Category: hDML,
		//line sql.y: 5440
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5447
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5621
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5622
		Category: hDML,
		//line sql.y: 5623
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5634
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5635
		Category: hDML,
		//line sql.y: 5636
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 5648
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5723
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5724
		Category: hDML,
		//line sql.y: 5725
		Text: `TABLE <tablename>
`,
		//line sql.y: 5726
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6014
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6015
		Category: hDML,
		//line sql.y: 6016
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6017
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6122
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6123
		Category: hDML,
		//line sql.y: 6124
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 6142
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
