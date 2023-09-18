import {TsukuyomiCommand} from "../structures/Command";
import {CommandInteraction, SlashCommandBuilder, User} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";
import {summary} from "../../lib/prisma/CharacterStatus";

export class StatusShow implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('status-show')
        .setDescription('ステータスを表示します．')

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (!interaction.guild) {
            return await interaction.reply('特定のサーバから送信してください．')
        } else {
            await interaction.deferReply()
            await interaction.followUp('探してるよ～ん')
            const user = interaction.user
            const status = await prisma.characterStatus.findUnique({
                where: {
                    guildId_userId: {
                        guildId: interaction.guildId!,
                        userId: user.id
                    }
                }
            })
            if (!status) {
                return await interaction.followUp(`Characterが作成されていません`)
            } else {
                return await interaction.followUp(summary(user.globalName, status))
            }
        }
    }
}