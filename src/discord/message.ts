import {Client, Events} from "discord.js";
import {ChannelType} from "discord-api-types/v10"
import {GrantPointDefinitionType, handlePoint} from "../lib/prisma";

export function registerMessage(client: Client) {
    client.on(Events.MessageCreate, async (message) => {
        if (message.author === client.user) {
            return
        }
        message.channel
    })

    client.on(Events.MessageReactionAdd, async (reaction, user) => {
        const channel = await reaction.message.channel.fetch()
        if (channel.type === ChannelType.PublicThread) {
            await handlePoint(user, reaction.message.guild!, GrantPointDefinitionType.REACTION_THREAD, false)
        }
    })

    client.on(Events.MessageReactionRemove, async (reaction, user) => {
        const channel = await reaction.message.channel.fetch()
        if (channel.type === ChannelType.PublicThread) {
            await handlePoint(user, reaction.message.guild!, GrantPointDefinitionType.REACTION_THREAD, true)
        }
    })
}
