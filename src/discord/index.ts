import {Channel, Client, Events, GatewayIntentBits, Partials, SlashCommandBuilder} from "discord.js";
import {registerGuild} from "./guild"
import {handleInvite} from "../lib/prisma";
import {registerThread} from "./thread";
import {registerMessage} from "./message";

const client = new Client({
    intents: [
        GatewayIntentBits.Guilds,
        GatewayIntentBits.GuildMembers,
        GatewayIntentBits.GuildInvites,
        GatewayIntentBits.GuildMessages,
        GatewayIntentBits.GuildMessageReactions,
        GatewayIntentBits.MessageContent,
        GatewayIntentBits.GuildScheduledEvents
    ],
    partials: [
        Partials.Message,
        Partials.Channel,
        Partials.Reaction
    ]
})

client.once(Events.ClientReady, async (client) => {
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
})


registerGuild(client)
registerThread(client)
registerMessage(client)

export async function notify(guildId: string, message: string) {
    const channel = await client.channels.fetch('1148960300968726558')
    // @ts-ignore
    await channel!.send(message)
}

export default client
