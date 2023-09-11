import {PrismaClient} from '@prisma/client'
import {notify} from "../discord";
import {Guild, PartialUser, User} from "discord.js";

export const prisma = new PrismaClient()

/**
 GrantPointDefinition
 ポイント付与数定義
 **/
export const GrantPointDefinitionType = {
    /** INVITED 招待を行った  */
    INVITED: 100,
    /** CREATE_EVENT イベントを作成した **/
    CREATE_EVENT: 1,
    /** HOLD_EVENT イベントを実施した **/
    HOLD_EVENT: 20,
    /** ATTEND_EVENT イベントに参加した **/
    ATTEND_EVENT: 1,
    /** CREATE_THREAD 投稿を作成した **/
    CREATE_THREAD: 5,
    /** REACTION_THREAD 投稿にリアクションをした **/
    REACTION_THREAD: 1,
    /** WIN_MIN 勝負に勝利した最低付与数 **/
    WIN_MIN: 1,
    /** WIN_MAX 勝負に勝利した最高付与数 **/
    WIN_MAX: 5,
} as const

export type GrantPointDefinitionType = typeof GrantPointDefinitionType[keyof typeof GrantPointDefinitionType]


export async function handlePoint(user: User | PartialUser, guild: Guild, type: GrantPointDefinitionType, negative: boolean) {
    let grantPoint: number = type
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
    await notify(guild.id, `${user.globalName}に${grantPoint}を付与しました． ${user.globalName}の所持pigは${point.value}です．`)
}

export async function handleInvite(code: string, inviterId: string | null, uses: number | null) {
    const invite = await prisma.invite.upsert({
        where: {
            code: code
        },
        create: {
            code: code,
            inviterId: inviterId,
            uses: uses || 0,
            expires: null
        },
        update: {
            uses: uses || 0,
            expires: null
        }
    })
    console.info(`Register invite ${invite.code}`)
}