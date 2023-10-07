import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class MeigenAdd implements TsukuyomiCommand {
  builder = new SlashCommandBuilder()
    .setName('meigen-add')
    .setDescription('名言/迷言の登録')
    .addUserOption(option => option
      .setName('user')
      .setDescription('発言した人を指定します．')
      .setRequired(true)
    )
    .addStringOption(option => option
      .setName('meigen')
      .setDescription('名言の内容を入力してください．')
      .setRequired(true)
    )

  async run(client: TsukuyomiClient, interaction: CommandInteraction) {
    if (interaction.guildId) {
      await interaction.deferReply()
      const user = interaction.options.getUser('user')!
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      const meigen = interaction.options.getString('meigen')!
      await prisma.meigen.create({
        data: {
          userId: user.id,
          content: meigen,
          registerId: interaction.user.id
        }
      })
      return await interaction.followUp(`${interaction.user}が${user}の名言「${meigen}」を登録しました．`)
    } else {
      return await interaction.followUp('Direct Messageでは使用できません')
    }
  }
}