import {ClientEvents, Events, GuildMember} from "discord.js";
import {GrantPointDefinitionType, prisma} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";
import {handleInvite} from "../../lib/prisma/Invite";

export class GuildMemberAdd implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.GuildMemberAdd
    point: number = GrantPointDefinitionType.INVITED

    async run(client: TsukuyomiClient, member: GuildMember) {
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
                    const point = await handlePoint(discordInvite.inviter!, guild, GrantPointDefinitionType.INVITED, false)
                    await notify(guild.id, `${discordInvite.inviter?.globalName}が"${member.user.globalName}"を招待したので${this.point}を付与しました`)
                }
            } else {
                // Bot側で把握している招待はない場合
                if (discordInvite.uses) {
                    // 招待数が１以上であればポイントの付与を行う．
                    const point = await handlePoint(discordInvite.inviter!, guild, GrantPointDefinitionType.INVITED, false)
                    await notify(guild.id, `${discordInvite.inviter?.globalName}が"${member.user.globalName}"を招待したので${this.point}を付与しました`)
                }
            }
            await handleInvite(discordInvite.code, discordInvite.inviterId, discordInvite.uses)
        }
    }
}