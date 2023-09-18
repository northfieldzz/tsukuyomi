import {PrismaClient} from "@prisma/client";

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
