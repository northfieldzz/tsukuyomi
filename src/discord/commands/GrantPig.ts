import {TsukuyomiCommand} from "../structures/Command";
import {SlashCommandBuilder} from "discord.js";


module.exports = new TsukuyomiCommand({
    builder: new SlashCommandBuilder()
        .setName('grant-pig')
        .setDescription('pigを付与します'),
    run: async (client, interaction) => {
        await interaction.reply('Pong')
    }
})


