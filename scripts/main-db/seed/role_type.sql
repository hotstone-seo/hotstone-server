INSERT INTO role_type (id, "name", modules)
VALUES 
    (1, 'Admin','{"modules": [ {"name": "rule", "pattern": "rules*", "path": "rules" }, {"name": "datasources", "pattern": "datasources*", "path": "datasources" }, {"name": "mismatchrule", "pattern": "mismatch*", "path": "mismatch-rule" }, {"name": "analytic", "pattern": "analytic*", "path": "analytic" }, {"name": "simulation", "pattern": "simulation*", "path": "simulation" }, {"name": "audittrail", "pattern": "audit*", "path": "audit-trail" }, {"name": "user", "pattern": "users*", "path": "users" }, {"name": "clientkey", "pattern": "client-keys*", "path": "client-keys" }, {"name": "roletype", "path": "role-type", "pattern": "role-type*"}, {"name": "module", "path": "modules", "pattern": "module*"} ] }'),
    (2, 'Staff','{"modules": [ {"name": "rule", "pattern": "rules*", "path": "rules" }, {"name": "datasources", "pattern": "datasources*", "path": "datasources" }, {"name": "mismatchrule", "pattern": "mismatch*", "path": "mismatch-rule" }, {"name": "analytic", "pattern": "analytic*", "path": "analytic" }, {"name": "simulation", "pattern": "simulation*", "path": "simulation" } ] }');

SELECT setval('role_type_id_seq', (SELECT MAX(id) FROM role_type));