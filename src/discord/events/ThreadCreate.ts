import {GrantPointDefinitionType} from "../../lib/prisma";
import {AnyThreadChannel, ClientEvents, Events} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";

export class ThreadCreate implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.ThreadCreate
    point: number = GrantPointDefinitionType.CREATE_THREAD

    async run(client: TsukuyomiClient, threadChannel: AnyThreadChannel) {
        const owner = await threadChannel.fetchOwner()
        await handlePoint(owner!.user!, threadChannel.guild, this.point, false)
        await notify(threadChannel.guildId, `${owner!.user!.globalName}が"${threadChannel.name}"を作成したので${this.point}を付与しました`)
    }
}