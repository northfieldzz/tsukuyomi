import {AnyThreadChannel, Events} from "discord.js";
import {GrantPointDefinitionType, handlePoint} from "../../lib/prisma";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

module.exports = new TsukuyomiEvent({
    name: Events.ThreadDelete,
    run: async (client: TsukuyomiClient, threadChannel: AnyThreadChannel) => {
        const owner = await threadChannel.fetchOwner()
        await handlePoint(owner!.user!, threadChannel.guild, GrantPointDefinitionType.CREATE_THREAD, true)
    }
})