import {ClientEvents, Events, Message} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../types/Event";
import {ChannelType} from "discord-api-types/v10"
import {GrantPointDefinitionType} from "../../lib/prisma";
import {handlePoint} from "../../lib/prisma/Point";
import {notify} from "../bot";

export class MessageCreate implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.MessageCreate

    async run(client: TsukuyomiClient, message: Message) {
        if (message.author.bot || !client.user || client.user.id === message.author.id) {
            return
        }
        switch (message.channel.type) {
            case ChannelType.PublicThread:
                if (message.guild) {
                    const messages = await message.channel.messages.fetch()
                    const isFirstTime = messages.find((m) => m.author.id === message.author.id)
                    if (isFirstTime) {
                        const point = GrantPointDefinitionType.REACTION_THREAD
                        await handlePoint(message.author, message.guild, point, false)
                        await notify(message.guild!.id, `${message.author.globalName}が${message.channel.name}に対してコメントしたので${point}を付与しました`)
                    } else {
                        console.info(`Already reaction`)
                    }
                } else {
                    console.info(`Direct Messageでは使用できません`)
                }
                console.log(message)
        }
        message.channel
    }
}