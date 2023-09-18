import {TsukuyomiCommand} from "../structures/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class PigShow implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('pig-show')
        .setDescription('所有ポイントを表示')

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (!interaction.guildId) {
            return await interaction.reply('特定のサーバから送信してください．')
        }
        const point = await prisma.point.findUnique({
            where: {
                userId_guildId: {
                    userId: interaction.user.id,
                    guildId: interaction.guildId!
                }
            }
        })
        return await interaction.reply(`${interaction.user.globalName}は現在${point?.value || 0}pig保有してます`)
    }
}