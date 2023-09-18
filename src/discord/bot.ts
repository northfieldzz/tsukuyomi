import TsukuyomiClient from "./structures/Clients";
import {GatewayIntentBits, Partials} from "discord.js";
import {prisma} from "../lib/prisma";


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
    try {
        const notifyChannel = await prisma.notificationChannel.findUnique({
            where: {
                guildId: guildId
            }
        })
        const channel = await client.channels.fetch(notifyChannel!.channelId)
        // @ts-ignore
        await channel!.send(message)
    } catch (error) {
        console.info('Not found notify channel')
        return
    }
}

export default client
