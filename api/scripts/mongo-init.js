print('Starting Database Init ###################################################################');

var dbName = 'forgeResponse'
var accessKeyCollection = "accessKeys"
var incidentCollection = "incidents"
var incidentCommentCollection = "incidentComments"
var investigationCollection = "investigations"
var tenantCollection = 'tenants'
var userCollection = 'users'
var webhookCollection = 'webhooks'

db = db.getSiblingDB(dbName)
db.createCollection(accessKeyCollection)
db.createCollection(incidentCollection)
db.createCollection(incidentCommentCollection)
db.createCollection(investigationCollection)
db.createCollection(tenantCollection)
db.createCollection(userCollection)
db.createCollection(webhookCollection)

db.users.createIndex({ "email": 1}, { unique: true })

print('Adding Access Keys ####################')
// API Key: forge_SQFo1b3.QjMLd36neDn5HQpRTREe97x2zc59dP9dSreUONGmlYvt
db.accessKeys.insert([
    {
        "id": "99fae8de-37f9-45de-95e8-446d09bad2fe",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "description": "Testing Key Tenant 1",
        "expiration": "2099-01-10T11:01:40.484+00:00",
        "keyHash": "b27a8cb054887f4353e3a0dca8933a399ce89e74d7f50702f43b9251524d9366",
        "keyPrefix": "forge_SQFo1b3",
        "scopes": [
            "delete:accesskeys", "modify:accesskeys", "read:accesskeys", "write:accesskeys",
            "read:events",
            "delete:incident", "modify:incident", "read:incident", "write:incident",
            "delete:incidentcomment", "modify:incidentcomment", "read:incidentcomment", "write:incidentcomment",
            "delete:investigations", "modify:investigations", "read:investigations", "write:investigations",
            "delete:investigationtemplates", "modify:investigationtemplates", "read:investigationtemplates", "write:investigationtemplates",
            "read:scopes",
            "delete:tenants", "modify:tenants", "read:tenants", "write:tenants",
            "delete:users", "modify:users", "read:users", "write:users",
            "delete:webhooks", "modify:webhooks", "read:webhooks", "write:webhooks"
        ],
        "version": 1
    },
    // API Key: forge_t4WEdrc.Jk6a9kYO2sWA3D9P3OoUqL58fukROgfAuqecxCQwZgPg
    {
        "id": "8d743f30-6a17-476e-8405-14678801e60c",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "description": "Testing Key Tenant 1",
        "expiration": "2099-01-10T11:01:40.484+00:00",
        "keyHash": "540eda7cf71593fe0b52c6c724bfc92a376bcfa8862d15fcf64476c542ffa6dd",
        "keyPrefix": "forge_t4WEdrc",
        "scopes": [
            "delete:accesskeys", "modify:accesskeys", "read:accesskeys", "write:accesskeys",
            "read:events",
            "delete:incident", "modify:incident", "read:incident", "write:incident",
            "delete:incidentcomment", "modify:incidentcomment", "read:incidentcomment", "write:incidentcomment",
            "delete:investigations", "modify:investigations", "read:investigations", "write:investigations",
            "delete:investigationtemplates", "modify:investigationtemplates", "read:investigationtemplates", "write:investigationtemplates",
            "read:scopes",
            "delete:tenants", "modify:tenants", "read:tenants", "write:tenants",
            "delete:users", "modify:users", "read:users", "write:users",
            "delete:webhooks", "modify:webhooks", "read:webhooks", "write:webhooks"
        ],
        "version": 1
    },
    // API Key: forge_qppbDPl.n07VqUhL5eFwyLfSOJ7LY7NYgGeeF5BqsWy8IEg4z7Q6
    {
        "id": "e52fb82a-fed6-40c1-ab00-11c1880d2a39",
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "description": "Testing Key Tenant 2",
        "expiration": "2099-01-10T11:01:40.484+00:00",
        "keyHash": "b667b8e213eb7b84fa45715f9f39f4328921caeda1ee5456868eb771fba0d65",
        "keyPrefix": "forge_qppbDPl",
        "scopes": [
            "delete:accesskeys", "modify:accesskeys", "read:accesskeys", "write:accesskeys",
            "read:events",
            "delete:incident", "modify:incident", "read:incident", "write:incident",
            "delete:incidentcomment", "modify:incidentcomment", "read:incidentcomment", "write:incidentcomment",
            "delete:investigations", "modify:investigations", "read:investigations", "write:investigations",
            "delete:investigationtemplates", "modify:investigationtemplates", "read:investigationtemplates", "write:investigationtemplates",
            "read:scopes",
            "delete:tenants", "modify:tenants", "read:tenants", "write:tenants",
            "delete:users", "modify:users", "read:users", "write:users",
            "delete:webhooks", "modify:webhooks", "read:webhooks", "write:webhooks"
        ],
        "version": 1
    },
    // API Key: forge_pNh4No9.e0jSldY8JxyWpYYoXIvYmEsyg1c9P8jmTU3pRBDA8yT4
    {
        "id": "dbd60e47-19c5-49d7-a186-eb24ff597510",
        "tenantId": "16167d1f-b2c7-4286-82e6-dde3d5eb449d",
        "description": "Testing Key Tenant 3",
        "expiration": "2099-01-10T11:01:40.484+00:00",
        "keyHash": "9c4aaecfb216dfb0a18c522c18018e10f567f209f301a497df1f4d448c1d07ef",
        "keyPrefix": "forge_pNh4No9",
        "scopes": [
            "delete:accesskeys", "modify:accesskeys", "read:accesskeys", "write:accesskeys",
            "read:events",
            "delete:incident", "modify:incident", "read:incident", "write:incident",
            "delete:incidentcomment", "modify:incidentcomment", "read:incidentcomment", "write:incidentcomment",
            "delete:investigations", "modify:investigations", "read:investigations", "write:investigations",
            "delete:investigationtemplates", "modify:investigationtemplates", "read:investigationtemplates", "write:investigationtemplates",
            "read:scopes",
            "delete:tenants", "modify:tenants", "read:tenants", "write:tenants",
            "delete:users", "modify:users", "read:users", "write:users",
            "delete:webhooks", "modify:webhooks", "read:webhooks", "write:webhooks"
        ],
        "version": 1
    }
])

print('Adding Incidents ####################')
db.incidents.insert([
    {
        "id": "87e31788-2625-4bb2-ae21-004e07f76d45",
        "assignedTo": "36235e81-e789-49cd-b3d8-a7d164982a50",
        "attachments": null,
        "createdBy": "",
        "description": "Test Description",
        "severity": "critical",
        "status": "in-progress",
        "tags": [""],
        "tasks": null,
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "title": "Test Incident 1",
        "tlp": 3,
        "version": 1
    },
    {
        "id": "670dbf0d-fe67-4b25-8dc7-22c2557fbf06",
        "assignedTo": "36235e81-e789-49cd-b3d8-a7d164982a50",
        "attachments": null,
        "createdBy": "",
        "description": "Test Description",
        "severity": "critical",
        "status": "in-progress",
        "tags": [""],
        "tasks": null,
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "title": "Test Incident 2",
        "tlp": 3,
        "version": 1
    },
    {
        "id": "867dbf0d-fe67-4b25-8dc7-22c2667fbf39",
        "assignedTo": "36235e81-e789-49cd-b3d8-a7d164982a50",
        "attachments": null,
        "createdBy": "",
        "description": "Test Description",
        "severity": "critical",
        "status": "in-progress",
        "tags": [""],
        "tasks": null,
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "title": "Test Incident 1",
        "tlp": 3,
        "version": 1
    }
])

print('Adding Incident Comments ####################')
db.incidentComments.insert([
    {
        "id": "0509add2-0ed1-469c-8af3-6dd95d293eae",
        "incidentId": "87e31788-2625-4bb2-ae21-004e07f76d45",
        "comment": "Test comment",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "createdBy": "",
        "order": 0,
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d"
    },
    {
        "id": "5447dc2f-92fa-4a5d-b884-7b624a8ae67b",
        "incidentId": "87e31788-2625-4bb2-ae21-004e07f76d45",
        "comment": "Test comment",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "createdBy": "",
        "order": 1,
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d"
    },
    {
        "id": "a8f2987f-dabb-45bd-a029-2c2d0a1660ac",
        "incidentId": "670dbf0d-fe67-4b25-8dc7-22c2557fbf06",
        "comment": "Test comment",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "createdBy": "",
        "order": 0,
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d"
    },
    {
        "id": "7cd5e91c-49e8-42fb-a5e4-d2241b8a84db",
        "incidentId": "867dbf0d-fe67-4b25-8dc7-22c2667fbf39",
        "comment": "Test comment",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "createdBy": "",
        "order": 0,
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f"
    }
])

print('Adding Investigations ####################')
db.investigations.insert([
    {
        "id": "5a85fe12-b4d6-43b8-acbb-736e94aaba35",
        "assignedTo": "3298224e-e54b-4515-b724-71921d963e6d",
        "attachments": null,
        "createdBy": "",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "comments": null,
        "description": "Test Investigation 1",
        "investigationId": "",
        "severity": "Critical",
        "status": "In Progress",
        "tags": [""],
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "title": "Investigation 1",
        "tlp": 1,
        "version": 1
    },
    {
        "id": "9108b666-34c1-445b-99fa-a263f0b247b3",
        "assignedTo": "3298224e-e54b-4515-b724-71921d963e6d",
        "attachments": null,
        "createdBy": "",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "comments": null,
        "description": "Test Investigation 2",
        "investigationId": "",
        "severity": "Critical",
        "status": "In Progress",
        "tags": [""],
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "title": "Investigation 2",
        "tlp": 1,
        "version": 1
    },
    {
        "id": "a58df6f9-bbd4-450c-8fb0-588411546666",
        "assignedTo": "a263321b-f7f7-4cfb-af86-f8628b6562df",
        "attachments": null,
        "createdBy": "",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "comments": null,
        "description": "Test Investigation 3",
        "investigationId": "",
        "severity": "Critical",
        "status": "In Progress",
        "tags": [""],
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "title": "Investigation 3",
        "tlp": 1,
        "version": 1
    }
])

print('Adding Investigation Templates ###################')
db.investigationsTemplates.insert([
    {
        "id": "0a7e3e90-6e96-4c0c-b522-d1531c4539ef",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "createdBy": "",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "description": "This is a test template",
        "titlePrefix": "Test Template 1",
        "severity": "Critical",
        "status": "New",
        "tags": [""],
        "tlp": 1
    },
    {
        "id": "5bde8616-b0b0-4541-8032-56a18f69b5ca",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "createdBy": "",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "description": "This is a test template",
        "titlePrefix": "Test Template 2",
        "severity": "Critical",
        "status": "New",
        "tags": [""],
        "tlp": 1
    },
    {
        "id": "e8adeff8-5827-46c7-adb8-9bc2809d2e0f",
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "createdBy": "",
        "createdAt": "2023-01-10T11:01:40.484+00:00",
        "description": "This is a test template",
        "titlePrefix": "Test Template 3",
        "severity": "Critical",
        "status": "New",
        "tags": [""],
        "tlp": 1
    }
])


print('Adding Tenants ####################')
db.tenants.insert([
    {
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "name": "Tenant 1"
    },
    {
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "name": "Tenant 2"
    },
    {
        "tenantId": "16167d1f-b2c7-4286-82e6-dde3d5eb449d",
        "name": "Tenant 3"
    }
])

print('Adding Users ####################')
db.users.insert([
    {
        "id": "3298224e-e54b-4515-b724-71921d963e6d",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "email": "test@test.com",
        "firstName": "Test",
        "lastName": "Test",
        "lastSignIn": "2023-01-10T11:01:40.484+00:00",
        "createdTime": "2023-01-10T11:01:40.484+00:00",
        "roles": null,
        "isActive": true
    },
    {
        "id": "40dde5ed-5127-4b63-a144-c0e815cb84b6",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "email": "test3@test.com",
        "firstName": "Test",
        "lastName": "Test",
        "lastSignIn": "2023-01-10T11:01:40.484+00:00",
        "createdTime": "2023-01-10T11:01:40.484+00:00",
        "roles": null,
        "isActive": true
    },
    {
        "id": "a263321b-f7f7-4cfb-af86-f8628b6562df",
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "email": "test2@test.com",
        "firstName": "Test 2",
        "lastName": "Test 2",
        "lastSignIn": "2023-01-10T11:01:40.484+00:00",
        "createdTime": "2023-01-10T11:01:40.484+00:00",
        "roles": null,
        "isActive": false
    }
])

print("Adding Webhooks ####################")
db.webhooks.insert([
    {
        "id": "304c73db-d06a-48d1-a490-ff1c69ab4a6d",
        "algorithm": "sha256",
        "description": "test webhook",
        "events": [
            "access_key_created"
        ],
        "secret": "0a2400674774e1d834fb4bdb1f257d8612961025e750360b1d5be5236270c26ac886d3d3f0681255007aed8208691c5faccd50fb410fb2800419dc30f8899d8ad291f0dfdf1518090e477a287125fb141bf6b07acb697edb560cfbef07fad5622c4f0aa5b859719dd9d21e75a881d9dccb7db570da8f057b3d9680241d",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "url": "0a2400674774e178c848f5709a33666474823aa1151cf9595e881db0f16e008079fcbc6eacd51242007aed8208ef4a50206124e461ae779e012f1b19913e20d1a30e6030183ec9daaabc94371778b46c99a52539209039049e6e75686f6dd54db2724e329d87ffea9ef7",
        "version": 1
    },
    {
        "id": "196d8f3b-39f9-4472-920d-62ce02163f0e",
        "algorithm": "sha256",
        "description": "test webhook",
        "events": [
            "access_key_created"
        ],
        "secret": "0a2400674774e1b30ea51224058be7c6b06aed410cb673bdfb416ab03c6b680f5f053715a03d1254007aed820801a44d33e363f43838311d998449cd0c4c70afa7233d66032e541c2ae3ca97fd4ac4a7f50a99aa56b12035056f1ba0256e4b1c2a5db569c17a44c27a8622d99325b519ccede575ae155147bc94e921",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "url": "0a2400674774e15274904c370e6598dcdcda8b942473bf66c475bd37f93b690516fef502bb7c1242007aed820838f183c8717304abb50e66d225f7975ef43eb993734596056f45fbc8c00e89ac804df2c72531ee3970959a3a110cd7f3777f46ebee71214966428543ec",
        "version": 1
    },
    {
        "id": "86ecacb3-94b2-4e47-aab1-9de2b973194f",
        "algorithm": "sha256",
        "description": "test webhook",
        "events": [
            "access_key_created"
        ],
        "secret": "0a2400674774e13dc542bb3fbe125dd5f6ca047ed2eb1947815c5ac7e2a202478216cec757a71255007aed8208cc9a9e4ec75b9087ea0705f8d24c35700a0148c2c0e19ccd0005bd7ae97b99d1318901702d5faf7b8fcad9dee8e3f889f185c32231729fbfef56f5c6ddfab697acade946a884139787e872fd6fb9739e",
        "tenantId": "0cb71832-042f-4f1f-aa6c-9ead59caa57d",
        "url": "0a2400674774e1899f902e5719138a5ff5ca53ca5f193ad63ec6c10790be086fb72c1f235bbd1242007aed8208b2d8cb793e8ff025848490f32dea110291a5caadd7111b06eb4b7ec9933a32357498e293ffd957a9d60cd8a4d5e32cdfa41478b74a92ddf52f25aa84ec",
        "version": 1
    },
    {
        "id": "c5535e1d-bb3c-4c1a-af10-0fa693a9e6d4",
        "algorithm": "sha256",
        "description": "test webhook",
        "events": [
            "access_key_created"
        ],
        "secret": "0a2400674774e15f75f75e34b15e0085df821f56ac27b5027bb24aa2dd300a9a00491c94e7681255007aed8208a0d9b406ee9131306729d71f1ebbccf39b0b82b876e9fec00eae154b95d38d69d79d0501d94f8b46aba6108d45224352d01c5733103e75313cb91f4328e45d84ac8d640275946e72a124c2d1bac5d1c7",
        "tenantId": "fcf825c5-faf2-4ae9-9ec9-caa4bb36779f",
        "url": "0a2400674774e1b8ee79e6b10ca984e0d3957cfd48e3a082924c67fbfcd88cb2a7429bc0af961242007aed8208f2d1ccf16f46fb814697861cbd95d3b7cddf4f815af7f379a09b81f6da07b4f0f766775ffc475c4473c9ac402277fe8a4660792e15929f1f4abbd1fa29",
        "version": 1
    }
])

print('End Database Init ###################################################################')