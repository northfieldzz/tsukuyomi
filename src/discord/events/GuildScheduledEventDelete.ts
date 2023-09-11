import {Events, GuildScheduledEvent} from "discord.js";
import {GrantPointDefinitionType, handlePoint} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

module.exports = new TsukuyomiEvent({
    name: Events.GuildScheduledEventDelete,
    run: async (client: TsukuyomiClient, event: GuildScheduledEvent) => {
        await handlePoint(event.creator!, event.guild!, GrantPointDefinitionType.CREATE_EVENT, true)
    }
})