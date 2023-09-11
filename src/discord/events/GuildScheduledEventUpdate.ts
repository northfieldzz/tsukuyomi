import {GuildScheduledEventStatus} from "discord-api-types/v10";
import {GrantPointDefinitionType, handlePoint} from "../../lib/prisma";
import {Events, GuildScheduledEvent} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

module.exports = new TsukuyomiEvent({
    name: Events.GuildScheduledEventUpdate,
    run: async (client: TsukuyomiClient, oldEvent: GuildScheduledEvent | null, newEvent: GuildScheduledEvent) => {
        switch (newEvent.status) {
            case GuildScheduledEventStatus.Scheduled:
                break
            case GuildScheduledEventStatus.Active:
                break
            case GuildScheduledEventStatus.Completed:
                await handlePoint(newEvent.creator!, newEvent.guild!, GrantPointDefinitionType.HOLD_EVENT, false)
                break
            case GuildScheduledEventStatus.Canceled:
                console.info('canceled')
                break
            default:
                break
        }
        console.log(oldEvent, newEvent);
    }
})