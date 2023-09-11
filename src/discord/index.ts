import TsukuyomiClient from "./structures/Clients";
import {GatewayIntentBits, Partials} from "discord.js";

const client = new TsukuyomiClient({
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

export async function notify(guildId: string, message: string) {
    const channel = await client.channels.fetch('1148960300968726558')
    // @ts-ignore
    await channel!.send(message)
}

export default client
