import {Client, Events, GuildScheduledEvent} from "discord.js";
import {GrantPointDefinitionType, handleInvite, handlePoint, prisma} from "../lib/prisma";
import {GuildScheduledEventStatus} from "discord-api-types/v10"

export function registerGuild(client: Client) {
    client.on(Events.GuildMemberAdd, async (member) => {
        const guild = await member.guild.fetch()
        const discordInvites = await guild.invites.fetch()
        const conditionInvites = []
        for (let discordInvite of discordInvites.map(invite => invite)) {
            conditionInvites.push({code: discordInvite.code})
        }
        // Bot側で把握している招待
        const tsukuyomiInvites = await prisma.invite.findMany({where: {OR: conditionInvites}})
        for (let discordInvite of discordInvites.map(invite => invite)) {
            const tsukuyomiInvite = tsukuyomiInvites.find(invite => discordInvite.code === invite.code)
            if (tsukuyomiInvite) {
                // Bot側で把握している招待である場合
                if (discordInvite.uses === tsukuyomiInvite.uses) {
                    // 招待数の増減がない場合，何もしない．
                    console.info(`No change detected: ${discordInvite.code}`)
                } else {
                    // 招待数が増えている場合,ポイントの付与を行う．
                    await handlePoint(discordInvite.inviter!, guild, GrantPointDefinitionType.INVITED, false)
                }
            } else {
                // Bot側で把握している招待はない場合
                if (discordInvite.uses) {
                    // 招待数が１以上であればポイントの付与を行う．
                    await handlePoint(discordInvite.inviter!, guild, GrantPointDefinitionType.INVITED, false)
                }
            }
            await handleInvite(discordInvite.code, discordInvite.inviterId, discordInvite.uses)
        }
    })

    // client.on(Events.GuildMemberRemove, async (member) => {
    //
    // })

    client.on(Events.GuildScheduledEventCreate, async (event: GuildScheduledEvent) => {
        await handlePoint(event.creator!, event.guild!, GrantPointDefinitionType.CREATE_EVENT, false)
    })

    client.on(Events.GuildScheduledEventUpdate, async (oldEvent, newEvent) => {
        switch (newEvent.status) {
            case GuildScheduledEventStatus.Scheduled:
                break
            case GuildScheduledEventStatus.Active:
                break
            case GuildScheduledEventStatus.Completed:
                await handlePoint(newEvent.creator!, newEvent.guild!, GrantPointDefinitionType.HOLD_EVENT, false)
                break
            case GuildScheduledEventStatus.Canceled:
                console.info('canceled')
                break
            default:
                break
        }

        console.log(oldEvent, newEvent);
    })

    client.on(Events.GuildScheduledEventDelete, async (event) => {
        await handlePoint(event.creator!, event.guild!, GrantPointDefinitionType.CREATE_EVENT, true)
    })
}
