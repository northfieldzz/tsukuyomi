import {Guild, PartialUser, User} from "discord.js";
import {Point} from "@prisma/client";
import {prisma} from "./index";

export async function handlePoint(
    user: User | PartialUser,
    guild: Guild, type: number,
    negative: boolean
): Promise<Point> {
    let grantPoint = type
    if (negative) {
        grantPoint = -grantPoint
    }
    const point = await prisma.point.upsert({
        where: {
            userId_guildId: {
                userId: user.id,
                guildId: guild.id
            }
        },
        create: {
            userId: user.id,
            guildId: guild.id,
            value: grantPoint
        },
        update: {
            value: {
                increment: grantPoint
            }
        }
    })
    console.info(`Grant point: ${grantPoint} -> ${user.id}(${point.value})`)
    return point
}