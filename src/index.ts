import client from "./discord"
import {prisma} from "./lib/prisma";


client.start()
    .catch((e) => {
        throw e
    })
    .finally(async () => {
        await prisma.$disconnect()
    })
