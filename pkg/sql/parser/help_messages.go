// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1090
	`ALTER`: {
		//line sql.y: 1091
		Category: hGroup,
		//line sql.y: 1092
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1106
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1107
		Category: hDDL,
		//line sql.y: 1108
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
		//line sql.y: 1132
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1145
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1146
		Category: hDDL,
		//line sql.y: 1147
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1149
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1156
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1157
		Category: hDDL,
		//line sql.y: 1158
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
	//line sql.y: 1181
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1182
		Category: hPriv,
		//line sql.y: 1183
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1185
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1190
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1191
		Category: hDDL,
		//line sql.y: 1192
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1194
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1205
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1206
		Category: hDDL,
		//line sql.y: 1207
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1215
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1568
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1569
		Category: hCCL,
		//line sql.y: 1570
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
		//line sql.y: 1587
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1595
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1596
		Category: hCCL,
		//line sql.y: 1597
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
		//line sql.y: 1613
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1631
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1632
		Category: hCCL,
		//line sql.y: 1633
		Text: `
IMPORT [ TABLE <tablename> FROM ]
       <format> ( <datafile> )
       [ WITH <option> [= <value>] [, ...] ]

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
		//line sql.y: 1659
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1688
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1689
		Category: hCCL,
		//line sql.y: 1690
		Text: `
EXPORT <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1699
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1786
	`CANCEL`: {
		//line sql.y: 1787
		Category: hGroup,
		//line sql.y: 1788
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1795
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1796
		Category: hMisc,
		//line sql.y: 1797
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1800
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1818
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1819
		Category: hMisc,
		//line sql.y: 1820
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1823
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1854
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1855
		Category: hMisc,
		//line sql.y: 1856
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1859
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1906
	`CREATE`: {
		//line sql.y: 1907
		Category: hGroup,
		//line sql.y: 1908
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1931
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 1932
		Category: hMisc,
		//line sql.y: 1933
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 1990
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1991
		Category: hDML,
		//line sql.y: 1992
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1996
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2011
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2012
		Category: hCfg,
		//line sql.y: 2013
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2025
	`DROP`: {
		//line sql.y: 2026
		Category: hGroup,
		//line sql.y: 2027
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2043
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2044
		Category: hDDL,
		//line sql.y: 2045
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2046
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2058
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2059
		Category: hDDL,
		//line sql.y: 2060
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2061
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2073
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2074
		Category: hDDL,
		//line sql.y: 2075
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2076
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2088
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2089
		Category: hDDL,
		//line sql.y: 2090
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2091
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2111
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2112
		Category: hDDL,
		//line sql.y: 2113
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2114
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2134
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2135
		Category: hPriv,
		//line sql.y: 2136
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2137
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2149
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2150
		Category: hPriv,
		//line sql.y: 2151
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2152
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2174
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2175
		Category: hMisc,
		//line sql.y: 2176
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
		//line sql.y: 2188
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2253
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2254
		Category: hMisc,
		//line sql.y: 2255
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2256
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2278
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2279
		Category: hMisc,
		//line sql.y: 2280
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2281
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2304
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2305
		Category: hMisc,
		//line sql.y: 2306
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2307
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2327
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2328
		Category: hPriv,
		//line sql.y: 2329
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
		//line sql.y: 2342
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2358
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2359
		Category: hPriv,
		//line sql.y: 2360
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
		//line sql.y: 2373
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2428
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2429
		Category: hCfg,
		//line sql.y: 2430
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2431
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2443
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2444
		Category: hCfg,
		//line sql.y: 2445
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2446
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2455
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2456
		Category: hCfg,
		//line sql.y: 2457
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2460
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2477
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2478
		Category: hExperimental,
		//line sql.y: 2479
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2487
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2493
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2494
		Category: hExperimental,
		//line sql.y: 2495
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2503
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2511
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2512
		Category: hExperimental,
		//line sql.y: 2513
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
		//line sql.y: 2524
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2579
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2580
		Category: hCfg,
		//line sql.y: 2581
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2582
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2603
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2604
		Category: hCfg,
		//line sql.y: 2605
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2611
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2628
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2629
		Category: hTxn,
		//line sql.y: 2630
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2637
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2814
	`SHOW`: {
		//line sql.y: 2815
		Category: hGroup,
		//line sql.y: 2816
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE, SHOW USERS,
SHOW TRANSACTION, SHOW BACKUP, SHOW JOBS, SHOW QUERIES, SHOW ROLES, SHOW SESSIONS, SHOW SYNTAX,
SHOW TRACE
`,
	},
	//line sql.y: 2848
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2849
		Category: hCfg,
		//line sql.y: 2850
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2851
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2872
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics`,
		//line sql.y: 2873
		Category: hMisc,
		//line sql.y: 2874
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2881
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2893
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram`,
		//line sql.y: 2894
		Category: hMisc,
		//line sql.y: 2895
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2899
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2912
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2913
		Category: hCCL,
		//line sql.y: 2914
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 2915
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2942
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2943
		Category: hCfg,
		//line sql.y: 2944
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2947
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2964
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2965
		Category: hDDL,
		//line sql.y: 2966
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2967
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2975
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2976
		Category: hDDL,
		//line sql.y: 2977
		Text: `SHOW DATABASES
`,
		//line sql.y: 2978
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2986
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2987
		Category: hPriv,
		//line sql.y: 2988
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 2994
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3007
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3008
		Category: hDDL,
		//line sql.y: 3009
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3010
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3028
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3029
		Category: hDDL,
		//line sql.y: 3030
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3031
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3044
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3045
		Category: hMisc,
		//line sql.y: 3046
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3047
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3063
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3064
		Category: hMisc,
		//line sql.y: 3065
		Text: `SHOW JOBS
`,
		//line sql.y: 3066
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3074
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3075
		Category: hMisc,
		//line sql.y: 3076
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3078
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3101
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3102
		Category: hMisc,
		//line sql.y: 3103
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3104
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3120
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3121
		Category: hDDL,
		//line sql.y: 3122
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3123
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3149
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3150
		Category: hDDL,
		//line sql.y: 3151
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3163
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3164
		Category: hMisc,
		//line sql.y: 3165
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3174
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3175
		Category: hCfg,
		//line sql.y: 3176
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3177
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3196
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3197
		Category: hDDL,
		//line sql.y: 3198
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3199
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3217
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3218
		Category: hPriv,
		//line sql.y: 3219
		Text: `SHOW USERS
`,
		//line sql.y: 3220
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3228
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3229
		Category: hPriv,
		//line sql.y: 3230
		Text: `SHOW ROLES
`,
		//line sql.y: 3231
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3283
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3284
		Category: hMisc,
		//line sql.y: 3285
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3521
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3522
		Category: hMisc,
		//line sql.y: 3523
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3526
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3543
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3544
		Category: hDDL,
		//line sql.y: 3545
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
		//line sql.y: 3572
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4085
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4086
		Category: hDDL,
		//line sql.y: 4087
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
		//line sql.y: 4097
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4151
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4152
		Category: hDML,
		//line sql.y: 4153
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4154
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4162
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4163
		Category: hPriv,
		//line sql.y: 4164
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4165
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4187
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4188
		Category: hPriv,
		//line sql.y: 4189
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4190
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4202
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4203
		Category: hDDL,
		//line sql.y: 4204
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4205
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4235
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4236
		Category: hDDL,
		//line sql.y: 4237
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4245
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4449
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4450
		Category: hTxn,
		//line sql.y: 4451
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4452
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4460
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4461
		Category: hMisc,
		//line sql.y: 4462
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4465
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4482
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4483
		Category: hTxn,
		//line sql.y: 4484
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4485
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4500
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4501
		Category: hTxn,
		//line sql.y: 4502
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4510
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4523
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4524
		Category: hTxn,
		//line sql.y: 4525
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4528
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4552
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4553
		Category: hTxn,
		//line sql.y: 4554
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4555
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4668
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4669
		Category: hDDL,
		//line sql.y: 4670
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4671
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4740
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4741
		Category: hDML,
		//line sql.y: 4742
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4747
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4766
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4767
		Category: hDML,
		//line sql.y: 4768
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4772
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4877
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4878
		Category: hDML,
		//line sql.y: 4879
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4886
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5056
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5057
		Category: hDML,
		//line sql.y: 5058
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5069
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5070
		Category: hDML,
		//line sql.y: 5071
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
		//line sql.y: 5083
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5158
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5159
		Category: hDML,
		//line sql.y: 5160
		Text: `TABLE <tablename>
`,
		//line sql.y: 5161
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5427
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5428
		Category: hDML,
		//line sql.y: 5429
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5430
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5520
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5521
		Category: hDML,
		//line sql.y: 5522
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
		//line sql.y: 5540
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
