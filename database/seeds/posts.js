const path = require("path");
const bytea = require('postgres-bytea')
const { Readable } = require('stream')
const streamToPromise = require('stream-to-promise')
const dotenv = require("dotenv");
const {Client} = require("pg");
const faker = require("faker");
const _ = require("lodash");

dotenv.config({
    path: path.normalize(path.join(__dirname, "../../config/.env")),
});
const POST_PARTS = 3;

async function populate_posts() {
    const n = Number(process.argv[2]);
    const client = new Client({
        user: process.env.BF_DATABASE_USER,
        password: process.env.BF_DATABASE_PASSWORD,
        host: process.env.BF_DATABASE_HOST,
        port: process.env.BF_DATABASE_PORT,
        database: process.env.BF_DATABASE_NAME,
    });
    await client.connect();

    const query_user = `INSERT INTO "user"(user_id, username, role, password_hash, created_on)
                        VALUES ($1, $2, $3, $4, $5)`;
    const query_post_published = `INSERT INTO "post"(post_id, author, title, slug, published, created_at, updated_at,
                                                     published_at)
                                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`;
    const query_post_not_published = `INSERT INTO "post"(post_id, author, title, slug, published, created_at, updated_at)
                                      VALUES ($1, $2, $3, $4, $5, $6, $7);`;
    const query_post_part = `INSERT INTO "post_part"(post_id, number, type, data)
                             VALUES ($1, $2, $3, $4);`;

    const admin_user_uuid = faker.datatype.uuid();
    // create admin user
    await client.query(query_user, [
        admin_user_uuid,
        "bogdanfloris",
        "admin",
        "password",
        faker.date.past(),
    ]);

    // Not published posts
    for (const _1 of _.range(n)) {
        const post_uuid = faker.datatype.uuid();
        // Create post
        await client.query(query_post_not_published, [
            post_uuid,
            admin_user_uuid,
            faker.lorem.word(),
            faker.lorem.slug(),
            false,
            faker.date.past(),
            faker.date.past(),
        ]);
        // Create 3 post parts
        for (const i of _.range(POST_PARTS)) {
            const fake_text = faker.lorem.paragraph();
            const encoded = (await streamToPromise(Readable.from(fake_text).pipe(new bytea.Encoder()))).toString();
            await client.query(query_post_part, [
                post_uuid,
                i,
                "text",
                encoded,
            ]);
        }
    }

    // Published posts
    for (const _1 of _.range(n)) {
        const post_uuid = faker.datatype.uuid();
        await client.query(query_post_published, [
            post_uuid,
            admin_user_uuid,
            faker.lorem.word(),
            faker.lorem.slug(),
            true,
            faker.date.past(),
            faker.date.past(),
            faker.date.recent(),
        ]);
        // Create 3 post parts
        for (const i of _.range(POST_PARTS)) {
            const fake_text = faker.lorem.paragraph();
            const encoded = (await streamToPromise(Readable.from(fake_text).pipe(new bytea.Encoder()))).toString();
            await client.query(query_post_part, [
                post_uuid,
                i,
                "text",
                encoded,
            ]);
        }
    }

    console.log(`Successfully seeded ${n * 2} posts with 3 post parts for author with uuid: ${admin_user_uuid}.`)
    await client.end();
}

populate_posts().then(_ => console.log("Ok"));