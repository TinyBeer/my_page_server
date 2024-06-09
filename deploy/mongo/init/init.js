db = db.getSiblingDB('admin');

db.createUser({
  user: 'my-page',
  pwd: '123456',
  roles: [{ role: 'readWrite', db: 'my-page' }],
});
