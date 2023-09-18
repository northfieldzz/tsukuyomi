import {GuildScheduledEventStatus} from "discord-api-types/v10";
import {GrantPointDefinitionType} from "../../lib/prisma";
import {ClientEvents, Events, GuildScheduledEvent} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";
import {notify} from "../bot";
import {handlePoint} from "../../lib/prisma/Point";

export class GuildScheduledEventUpdate implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.GuildScheduledEventUpdate
    point: number = GrantPointDefinitionType.HOLD_EVENT

    async run(client: TsukuyomiClient, oldEvent: GuildScheduledEvent | null, newEvent: GuildScheduledEvent) {
        switch (newEvent.status) {
            case GuildScheduledEventStatus.Scheduled:
                break
            case GuildScheduledEventStatus.Active:
                break
            case GuildScheduledEventStatus.Completed:
                await handlePoint(newEvent.creator!, newEvent.guild!, GrantPointDefinitionType.HOLD_EVENT, false)
                await notify(newEvent.guildId, `${newEvent.creator!.globalName}がイベントを終了したので${this.point}を付与しました`)
                break
            case GuildScheduledEventStatus.Canceled:
                console.info('canceled')
                break
            default:
                break
        }
        console.log(oldEvent, newEvent);
    }
}