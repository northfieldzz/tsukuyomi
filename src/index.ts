import client from "./discord/bot";
import {prisma} from "./lib/prisma";


client.start()
    .catch((e) => {
        throw e
    })
    .finally(async () => {
        await prisma.$disconnect()
    })
