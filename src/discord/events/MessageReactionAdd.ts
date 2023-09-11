import {Events, MessageReaction, PartialMessageReaction, PartialUser, User} from "discord.js";
import {ChannelType} from "discord-api-types/v10"
import {GrantPointDefinitionType, handlePoint} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

module.exports = new TsukuyomiEvent({
    name: Events.MessageReactionAdd,
    run: async (client: TsukuyomiClient, reaction: MessageReaction | PartialMessageReaction, user: User | PartialUser) => {
        const channel = await reaction.message.channel.fetch()
        if (channel.type === ChannelType.PublicThread) {
            await handlePoint(user, reaction.message.guild!, GrantPointDefinitionType.REACTION_THREAD, false)
        }
    }
})