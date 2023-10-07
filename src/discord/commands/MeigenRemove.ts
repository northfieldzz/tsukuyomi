import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class MeigenRemove implements TsukuyomiCommand {
  builder = new SlashCommandBuilder()
    .setName('meigen-remove')
    .setDescription('名言/迷言の削除')
    .addIntegerOption(option => option
      .setName('id')
      .setDescription('名言の内容でフィルターします．')
      .setRequired(true)
    )

  async run(client: TsukuyomiClient, interaction: CommandInteraction) {
    if (interaction.guild) {
      await interaction.deferReply()
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      const id = interaction.options.getInteger('id')!
      const meigen = await prisma.meigen.delete({
        where: {
          id: id
        }
      })
      const guild = await interaction.guild.fetch()
      const speaker = guild.members.cache.get(meigen.userId)
      return await interaction.followUp(`${interaction.user}が${speaker}の名言/迷言「${meigen.content}」を削除しました．`)
    } else {
      return await interaction.reply('Direct Messageでは使用できません')
    }
  }
}