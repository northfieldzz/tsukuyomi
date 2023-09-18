import {AnyThreadChannel, ClientEvents, Events} from "discord.js";
import {GrantPointDefinitionType} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";

export class ThreadDelete implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.ThreadDelete
    point: number = -GrantPointDefinitionType.CREATE_THREAD

    async run(client: TsukuyomiClient, threadChannel: AnyThreadChannel) {
        const owner = await threadChannel.fetchOwner()
        await handlePoint(owner!.user!, threadChannel.guild, GrantPointDefinitionType.CREATE_THREAD, true)
        await notify(threadChannel.guildId, `${owner!.user!.globalName}が"${threadChannel.name}"を削除したので${this.point}を付与しました`)
    }
}