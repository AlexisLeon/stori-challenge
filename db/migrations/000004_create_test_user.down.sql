DELETE FROM {{ index .Options "Namespace" }}.users WHERE id = '00000000-0000-0000-0000-000000000000';

DELETE FROM {{ index .Options "Namespace" }}.accounts WHERE id = '00000000-0000-0000-0000-000000000000';
