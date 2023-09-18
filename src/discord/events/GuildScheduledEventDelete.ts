import {ClientEvents, Events, GuildScheduledEvent} from "discord.js";
import {GrantPointDefinitionType} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";

export class GuildScheduledEventDelete implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.GuildScheduledEventDelete
    point: number = -GrantPointDefinitionType.CREATE_EVENT

    async run(client: TsukuyomiClient, event: GuildScheduledEvent) {
        await handlePoint(event.creator!, event.guild!, this.point, true)
        await notify(event.guildId, `${event.creator!.globalName}がイベントを作成したので${this.point}を付与しました`)
    }
}