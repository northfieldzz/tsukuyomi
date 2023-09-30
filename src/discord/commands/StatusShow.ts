import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";
import {summary} from "../../lib/prisma/CharacterStatus";

export class StatusShow implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('status-show')
        .setDescription('ステータスを表示します．')

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (interaction.guild) {
            await interaction.deferReply()
            await interaction.followUp('探してるよ～ん')
            const user = interaction.user
            const status = await prisma.characterStatus.findUnique({
                where: {
                    guildId_userId: {
                        guildId: interaction.guild.id,
                        userId: user.id
                    }
                }
            })
            if (!status) {
                return await interaction.followUp(`Characterが作成されていません`)
            } else {
                return await interaction.followUp(summary(user.globalName, status))
            }
        } else {
            return await interaction.reply('特定のサーバから送信してください．')
        }
    }
}