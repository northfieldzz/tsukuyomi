import {ChannelType} from "discord-api-types/v10";
import {GrantPointDefinitionType} from "../../lib/prisma";
import {ClientEvents, Events, MessageReaction, PartialMessageReaction, PartialUser, User} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";

export class MessageReactionRemove implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.MessageReactionRemove
    point: number = -GrantPointDefinitionType.REACTION_THREAD

    async run(client: TsukuyomiClient, reaction: MessageReaction | PartialMessageReaction, user: User | PartialUser) {
        const channel = await reaction.message.channel.fetch()
        if (channel.type === ChannelType.PublicThread) {
            await handlePoint(user, reaction.message.guild!, GrantPointDefinitionType.REACTION_THREAD, true)
            await notify(reaction.message.guildId!, `${user.globalName}がリアクションを削除したので${this.point}を付与しました`)
        }
    }
}