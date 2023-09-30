import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class PigShow implements TsukuyomiCommand {
  builder = new SlashCommandBuilder()
    .setName('pig-show')
    .setDescription('所有ポイントを表示')

  async run(client: TsukuyomiClient, interaction: CommandInteraction) {
    if (interaction.guildId) {
      const point = await prisma.point.findUnique({
        where: {
          userId_guildId: {
            userId: interaction.user.id,
            guildId: interaction.guildId
          }
        }
      })
      return await interaction.reply(`${interaction.user.globalName}は現在${point?.value || 0}pig保有してます`)
    } else {
      return await interaction.reply('Direct Messageでは使用できません')
    }
  }
}