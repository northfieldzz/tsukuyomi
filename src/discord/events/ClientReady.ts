
import {ClientEvents, Events, Routes} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event"
import {notify} from "../bot";
import {handleInvite} from "../../lib/prisma/Invite";

export class ClientReady implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.ClientReady

    async run(client: TsukuyomiClient) {
        console.info('Ready discord client')
        if (client.user) {
            console.info(`bot user tag: ${client.user.tag}`)
        }
        const guilds = client.guilds.cache.map(guild => guild)
        const commands = []
        for (const command of client.commands) {
            commands.push(command[1].builder.toJSON())
        }
        for (let guild of guilds) {
            let invites = await guild.invites.fetch()
            for (let invite of invites.map(invites => invites)) {
                await handleInvite(invite.code, invite.inviterId, invite.uses)
            }
            await client.rest.put(Routes.applicationCommands(process.env.DISCORD_APPLICATION_ID!), {body: commands})
            await notify(guild.id, 'つくよみ，いんです！')
        }
        console.info('Register invite complete')
    }
}
