import {Client, Events} from "discord.js";
import {GrantPointDefinitionType, handlePoint} from "../lib/prisma";

export function registerThread(client: Client) {
    client.on(Events.ThreadCreate, async (threadChannel) => {
        const owner = await threadChannel.fetchOwner()
        await handlePoint(owner!.user!, threadChannel.guild, GrantPointDefinitionType.CREATE_THREAD, false)
    })

    client.on(Events.ThreadDelete, async (threadChannel) => {
        const owner = await threadChannel.fetchOwner()
        await handlePoint(owner!.user!, threadChannel.guild, GrantPointDefinitionType.CREATE_THREAD, true)
    })
}

