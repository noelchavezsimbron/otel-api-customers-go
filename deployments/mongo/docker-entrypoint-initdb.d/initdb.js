const configuration = JSON.parse(cat('/etc/configs/config.json'));
const rootUser =configuration.admin.user
const rootPass = configuration.admin.pwd

print('user:',rootUser);

// auth against admin
const adminDb = db.getSiblingDB('admin');
adminDb.auth(rootUser, rootPass);
print('Successfully authenticated admin user');

// Create users
const dbName=configuration.databaseName;
const targetDb = db.getSiblingDB(dbName);


print('create users => start');

configuration.users.forEach(userCredential=>{
  print("configuration",JSON.stringify(userCredential));

  const roles = userCredential.roles.map(role => {
    return {role: role, db: dbName};
  });

  targetDb.createUser(
    {
      user: userCredential.user,
      pwd: userCredential.pwd,
      roles: roles
    }
  );


});

print('create users => complete');

// Create Collection with Data

const collectionName="customers"

targetDb.createCollection(collectionName);

targetDb[collectionName].insertMany(
    [
      {
        "document": "47506306",
        "name":     "Alexander",
        "role":     "Software Architect",
        "salary":   12000
      },
      {
        "document": "34806205",
        "name":     "Esteban",
        "role":     "Software Developer",
        "salary":   5000
      }
    ]
)
