// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1078
	`ALTER`: {
		//line sql.y: 1079
		Category: hGroup,
		//line sql.y: 1080
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1094
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1095
		Category: hDDL,
		//line sql.y: 1096
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

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1120
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1132
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1133
		Category: hDDL,
		//line sql.y: 1134
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1136
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1143
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1144
		Category: hDDL,
		//line sql.y: 1145
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
	//line sql.y: 1168
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1169
		Category: hPriv,
		//line sql.y: 1170
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1172
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1177
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1178
		Category: hDDL,
		//line sql.y: 1179
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1181
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1192
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1193
		Category: hDDL,
		//line sql.y: 1194
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1202
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1538
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1539
		Category: hCCL,
		//line sql.y: 1540
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
		//line sql.y: 1557
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1565
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1566
		Category: hCCL,
		//line sql.y: 1567
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
		//line sql.y: 1583
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1614
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1615
		Category: hCCL,
		//line sql.y: 1616
		Text: `
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

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1637
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1650
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1651
		Category: hCCL,
		//line sql.y: 1652
		Text: `
EXPORT <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1661
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1748
	`CANCEL`: {
		//line sql.y: 1749
		Category: hGroup,
		//line sql.y: 1750
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1757
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1758
		Category: hMisc,
		//line sql.y: 1759
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1762
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1780
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1781
		Category: hMisc,
		//line sql.y: 1782
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1785
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1816
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1817
		Category: hMisc,
		//line sql.y: 1818
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1821
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1868
	`CREATE`: {
		//line sql.y: 1869
		Category: hGroup,
		//line sql.y: 1870
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1892
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 1893
		Category: hMisc,
		//line sql.y: 1894
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 1924
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1925
		Category: hDML,
		//line sql.y: 1926
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1930
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1945
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1946
		Category: hCfg,
		//line sql.y: 1947
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1959
	`DROP`: {
		//line sql.y: 1960
		Category: hGroup,
		//line sql.y: 1961
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 1977
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1978
		Category: hDDL,
		//line sql.y: 1979
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1980
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1992
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 1993
		Category: hDDL,
		//line sql.y: 1994
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1995
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2007
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2008
		Category: hDDL,
		//line sql.y: 2009
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2010
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2022
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2023
		Category: hDDL,
		//line sql.y: 2024
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2025
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2045
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2046
		Category: hDDL,
		//line sql.y: 2047
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2048
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2068
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2069
		Category: hPriv,
		//line sql.y: 2070
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2071
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2083
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2084
		Category: hPriv,
		//line sql.y: 2085
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2086
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2108
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2109
		Category: hMisc,
		//line sql.y: 2110
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, VERBOSE

`,
		//line sql.y: 2122
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2187
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2188
		Category: hMisc,
		//line sql.y: 2189
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2190
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2212
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2213
		Category: hMisc,
		//line sql.y: 2214
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2215
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2238
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2239
		Category: hMisc,
		//line sql.y: 2240
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2241
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2261
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2262
		Category: hPriv,
		//line sql.y: 2263
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
		//line sql.y: 2276
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2292
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2293
		Category: hPriv,
		//line sql.y: 2294
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
		//line sql.y: 2307
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2362
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2363
		Category: hCfg,
		//line sql.y: 2364
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2365
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2377
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2378
		Category: hCfg,
		//line sql.y: 2379
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2380
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2389
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2390
		Category: hCfg,
		//line sql.y: 2391
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2394
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2411
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2412
		Category: hExperimental,
		//line sql.y: 2413
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2421
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2427
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2428
		Category: hExperimental,
		//line sql.y: 2429
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2437
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2445
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2446
		Category: hExperimental,
		//line sql.y: 2447
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
		//line sql.y: 2458
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2513
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2514
		Category: hCfg,
		//line sql.y: 2515
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2516
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2537
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2538
		Category: hCfg,
		//line sql.y: 2539
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { off | cluster | on | kv | local } [,...]

`,
		//line sql.y: 2545
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2562
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2563
		Category: hTxn,
		//line sql.y: 2564
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2571
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2746
	`SHOW`: {
		//line sql.y: 2747
		Category: hGroup,
		//line sql.y: 2748
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW CREATE SEQUENCE, SHOW USERS,
SHOW TRANSACTION, SHOW BACKUP, SHOW JOBS, SHOW QUERIES, SHOW ROLES, SHOW SESSIONS, SHOW SYNTAX,
SHOW TRACE
`,
	},
	//line sql.y: 2782
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2783
		Category: hCfg,
		//line sql.y: 2784
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2785
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2806
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics`,
		//line sql.y: 2807
		Category: hMisc,
		//line sql.y: 2808
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2815
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2827
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram`,
		//line sql.y: 2828
		Category: hMisc,
		//line sql.y: 2829
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2833
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2846
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2847
		Category: hCCL,
		//line sql.y: 2848
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 2849
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2874
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2875
		Category: hCfg,
		//line sql.y: 2876
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2879
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2896
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2897
		Category: hDDL,
		//line sql.y: 2898
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2899
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2907
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2908
		Category: hDDL,
		//line sql.y: 2909
		Text: `SHOW DATABASES
`,
		//line sql.y: 2910
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2918
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2919
		Category: hPriv,
		//line sql.y: 2920
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 2926
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2939
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2940
		Category: hDDL,
		//line sql.y: 2941
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2942
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2960
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2961
		Category: hDDL,
		//line sql.y: 2962
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2963
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2976
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2977
		Category: hMisc,
		//line sql.y: 2978
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2979
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 2995
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2996
		Category: hMisc,
		//line sql.y: 2997
		Text: `SHOW JOBS
`,
		//line sql.y: 2998
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3006
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3007
		Category: hMisc,
		//line sql.y: 3008
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
SHOW [COMPACT] [KV] TRACE FOR <statement>
`,
		//line sql.y: 3011
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3041
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3042
		Category: hMisc,
		//line sql.y: 3043
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3044
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3060
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3061
		Category: hDDL,
		//line sql.y: 3062
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3063
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3089
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3090
		Category: hDDL,
		//line sql.y: 3091
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3103
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3104
		Category: hMisc,
		//line sql.y: 3105
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3114
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3115
		Category: hCfg,
		//line sql.y: 3116
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3117
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3136
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 3137
		Category: hDDL,
		//line sql.y: 3138
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 3139
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3147
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 3148
		Category: hDDL,
		//line sql.y: 3149
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 3150
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 3158
	`SHOW CREATE SEQUENCE`: {
		ShortDescription: `display the CREATE SEQUENCE statement for a sequence`,
		//line sql.y: 3159
		Category: hDDL,
		//line sql.y: 3160
		Text: `SHOW CREATE SEQUENCE <seqname>
`,
	},
	//line sql.y: 3168
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3169
		Category: hPriv,
		//line sql.y: 3170
		Text: `SHOW USERS
`,
		//line sql.y: 3171
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3179
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3180
		Category: hPriv,
		//line sql.y: 3181
		Text: `SHOW ROLES
`,
		//line sql.y: 3182
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3234
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3235
		Category: hMisc,
		//line sql.y: 3236
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3472
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3473
		Category: hMisc,
		//line sql.y: 3474
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3477
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3494
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3495
		Category: hDDL,
		//line sql.y: 3496
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
		//line sql.y: 3523
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4034
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4035
		Category: hDDL,
		//line sql.y: 4036
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]

`,
		//line sql.y: 4045
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4098
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4099
		Category: hDML,
		//line sql.y: 4100
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4101
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4109
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4110
		Category: hPriv,
		//line sql.y: 4111
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4112
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4134
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4135
		Category: hPriv,
		//line sql.y: 4136
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4137
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4149
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4150
		Category: hDDL,
		//line sql.y: 4151
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4152
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4166
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4167
		Category: hDDL,
		//line sql.y: 4168
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4176
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4374
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4375
		Category: hTxn,
		//line sql.y: 4376
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4377
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4385
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4386
		Category: hMisc,
		//line sql.y: 4387
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4390
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4407
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4408
		Category: hTxn,
		//line sql.y: 4409
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4410
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4425
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4426
		Category: hTxn,
		//line sql.y: 4427
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4435
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4448
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4449
		Category: hTxn,
		//line sql.y: 4450
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4453
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4477
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4478
		Category: hTxn,
		//line sql.y: 4479
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4480
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4593
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4594
		Category: hDDL,
		//line sql.y: 4595
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4596
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4665
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4666
		Category: hDML,
		//line sql.y: 4667
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4672
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4691
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4692
		Category: hDML,
		//line sql.y: 4693
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4697
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4802
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4803
		Category: hDML,
		//line sql.y: 4804
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4811
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 4981
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 4982
		Category: hDML,
		//line sql.y: 4983
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 4994
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 4995
		Category: hDML,
		//line sql.y: 4996
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
		//line sql.y: 5008
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5083
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5084
		Category: hDML,
		//line sql.y: 5085
		Text: `TABLE <tablename>
`,
		//line sql.y: 5086
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5352
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5353
		Category: hDML,
		//line sql.y: 5354
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5355
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5456
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5457
		Category: hDML,
		//line sql.y: 5458
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

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5476
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
