const path = require("path");
const dotenv = require("dotenv");
const {Pool} = require("pg");
const faker = require("faker");
const _ = require("lodash");

dotenv.config({
    path: path.normalize(path.join(__dirname, "../../config/.env")),
});

async function populate_users() {
    const n = Number(process.argv[2]);
    const db_pool = new Pool({
        user: process.env.BF_DATABASE_USER,
        password: process.env.BF_DATABASE_PASSWORD,
        host: process.env.BF_DATABASE_HOST,
        port: process.env.BF_DATABASE_PORT,
        database: process.env.BF_DATABASE_NAME,
    });

    const query = `INSERT INTO "user"(user_id, username, role, password_hash, created_on)
                   VALUES ($1, $2, $3, $4, $5)`;
    for (const _1 of _.range(n)) {
        await db_pool.query(query, [
            faker.datatype.uuid(),
            faker.internet.userName(),
            "normal",
            "password",
            faker.date.past(),
        ]);
    }
    await db_pool.end();
}

populate_users().then(r => console.log(`Successfully seeded the users table.`));