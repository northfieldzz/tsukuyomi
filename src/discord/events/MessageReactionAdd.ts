import {ClientEvents, Events, MessageReaction, PartialMessageReaction, PartialUser, User} from "discord.js";
import {ChannelType} from "discord-api-types/v10"
import {GrantPointDefinitionType} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";

export class MessageReactionAdd implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.MessageReactionAdd
    point: number = GrantPointDefinitionType.REACTION_THREAD

    async run(client: TsukuyomiClient, reaction: MessageReaction | PartialMessageReaction, user: User | PartialUser) {
        const channel = await reaction.message.channel.fetch()
        if (channel.type === ChannelType.PublicThread) {
            await handlePoint(user, reaction.message.guild!, this.point, false)
            await notify(reaction.message.guildId!, `${user.globalName}がリアクションをしたので${this.point}を付与しました`)
        }
    }
}