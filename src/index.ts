import client from "./discord"
import {prisma} from "./lib/prisma";

async function main() {
    await client.login(process.env.DISCORD_TOKEN)
}

main()
    .catch((e) => {
        throw e
    })
    .finally(async () => {
        await prisma.$disconnect()
    })
