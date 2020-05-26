INSERT INTO role_type (id, "name", modules)
VALUES 
    (1, 'Admin','{"modules": [{"name": "rule", "path": "rules", "label": "Rules", "pattern": "rules*"}, {"name": "datasources", "path": "datasources", "label": "Data Sources", "pattern": "datasources*"}, {"name": "mismatchrule", "path": "mismatch-rule", "label": "Mismatch Rule", "pattern": "mismatch*"}, {"name": "analytic", "path": "analytic", "label": "Analytic", "pattern": "analytic*"}, {"name": "simulation", "path": "simulation", "label": "Simulation Page", "pattern": "simulation*"}, {"name": "audittrail", "path": "audit-trail", "label": "Audit Trail", "pattern": "audit*"}, {"name": "user", "path": "users", "label": "Users", "pattern": "users*"}, {"name": "roletype", "path": "role-type", "label": "User Role", "pattern": "role-type*"}, {"name": "module", "path": "modules", "label": "Modules", "pattern": "modules*"}, {"name": "clientkey", "path": "client-keys", "label": "Client Keys", "pattern": "client-keys*"}]}'),
    (2, 'Staff','{"modules": [{"name": "rule", "path": "rules", "label": "Rules", "pattern": "rules*"}, {"name": "datasources", "path": "datasources", "label": "Data Sources", "pattern": "datasources*"}, {"name": "mismatchrule", "path": "mismatch-rule", "label": "Mismatch Rule", "pattern": "mismatch*"}, {"name": "analytic", "path": "analytic", "label": "Analytic", "pattern": "analytic*"}, {"name": "simulation", "path": "simulation", "label": "Simulation Page", "pattern": "simulation*"}]}');

SELECT setval('role_type_id_seq', (SELECT MAX(id) FROM role_type));