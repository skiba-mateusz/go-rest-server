db = db.getSiblingDB("go-rest-server");
db.createCollection("records");
db.records.insertMany([
  {
    key: "DFGwGa9Wr9dduz49A",
    createdAt: ISODate("2020-07-26T04:01:25Z"),
    count: [100, 200, 300, 350, 150],
  },
  {
    key: "DFGwGa9Wr9dduz49A",
    createdAt: ISODate("2022-06-15T09:59:30Z"),
    count: [400, 25, 700, 300, 50],
  },
  {
    key: "DFGwGa9Wr9dduz49A",
    createdAt: ISODate("2024-05-14T19:35:23Z"),
    count: [150, 250, 100, 550, 250],
  },
  {
    key: "DFGwGa9Wr9dduz49A",
    createdAt: ISODate("2020-11-01T12:59:43Z"),
    count: [50, 200, 400, 50, 125],
  },
]);
