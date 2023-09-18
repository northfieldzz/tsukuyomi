import {Invite} from "@prisma/client";
import {prisma} from "./index";

export async function handleInvite(code: string, inviterId: string | null, uses: number | null): Promise<Invite> {
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
    return invite
}
