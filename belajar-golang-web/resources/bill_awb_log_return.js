// bill_awb_log_return
db.createCollection('bill_awb_log_return');

db.bill_awb_log_return.createIndex(
    { awbId: 1 },
    { name: 'idx_awbId_1', unique: false }
);
// find({
//     filter: { awbId: 1234567 },
//     projection: options?.projection,
//     limit: options?.limit ?? 1000,
//     session: session,
// });

db.bill_awb_log_return.createIndex(
    { awbNumber: 1 },
    { name: 'idx_awbNumber_1', unique: false }
);
// find({
//     filter: { awbNumber: '000012345456' },
//     projection: options?.projection,
//     limit: options?.limit ?? 1000,
//     session: session,
// });

// retention
db.bill_awb_log_return.createIndex(
    {
        revisionDateUtc: 1,
    },
    {
        name: 'bill_awb_log_ttl_idx',
        expireAfterSeconds: 2592000,
    }
);
