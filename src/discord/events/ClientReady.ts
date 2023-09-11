import {handleInvite} from "../../lib/prisma";
import {notify} from "../index";
import {Events} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event"

module.exports = new TsukuyomiEvent({
    name: Events.ClientReady,
    run: async (client: TsukuyomiClient) => {
        console.info('Ready discord client')
        if (client.user) {
            console.info(`bot user tag: ${client.user.tag}`)
        }
        const guilds = client.guilds.cache.map(guild => guild)
        for (let guild of guilds) {
            let invites = await guild.invites.fetch()
            for (let invite of invites.map(invites => invites)) {
                await handleInvite(invite.code, invite.inviterId, invite.uses)
            }
        }
        await notify('', 'Tsukuyomi Ready!')
        console.info('Register invite complete')

        // await client.rest!.put(Routes.applicationCommands(clientId), {body: client.commands})
    }
})
