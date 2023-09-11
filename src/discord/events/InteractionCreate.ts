import {CommandInteraction, Events} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

module.exports = new TsukuyomiEvent({
    name: Events.InteractionCreate,
    run: async (client: TsukuyomiClient, interaction: CommandInteraction) => {
        if (!interaction.isCommand()) {
            return
        }
        const command = client.commands.get(interaction.commandName)
        if (!command) {
            return
        }
        try {
            await command!.run(client, interaction)
        } catch (error) {
            return await interaction.followUp('sample')
        }
    }
})