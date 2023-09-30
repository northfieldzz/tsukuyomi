import {ClientEvents, CommandInteraction, Events} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../types/Event";

export class InteractionCreate implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.InteractionCreate

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (!interaction.isCommand()) {
            return
        }
        const command = client.commands.get(interaction.commandName)
        if (!command) {
            return
        }
        try {
            await command!.run(client, interaction)
        } catch (error: unknown) {
            let message = 'エラーが発生しました'
            if (error instanceof Error) {
                message = error.message
            }
            await interaction.followUp({content: message, ephemeral: true})
        }
    }
}