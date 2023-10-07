import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";
import {Prisma} from "@prisma/client";

export class MeigenList implements TsukuyomiCommand {
  builder = new SlashCommandBuilder()
    .setName('meigen-list')
    .setDescription('名言/迷言の一覧')
    .addUserOption(option => option
      .setName('user-filter')
      .setDescription('ユーザーでフィルターします．')
    )
    .addStringOption(option => option
      .setName('meigen-filter')
      .setDescription('名言の内容でフィルターします．')
    )
    .addIntegerOption(option => option
      .setName('page')
      .setDescription('ページ番号を指定します．')
      .setMinValue(1)
      .setMaxValue(100)
    )

  async run(client: TsukuyomiClient, interaction: CommandInteraction) {
    if (interaction.guild) {
      await interaction.deferReply()

      const user = interaction.options.getUser('user-filter')!
      const userId = user?.id ?? undefined
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      const contentFilter = interaction.options.getString('meigen-filter') ?? undefined
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      const page = interaction.options.getInteger('page') ?? 100
      const where: Prisma.MeigenWhereInput = {
        userId: userId,
        content: {
          contains: contentFilter
        }
      }
      const meigens = await prisma.meigen.findMany({
          take: page,
          where: where
        }
      )
      const display: String[] = []
      const guild = await interaction.guild.fetch()
      for (const meigen of meigens) {
        const speaker = guild.members.cache.get(meigen.userId)
        const register = guild.members.cache.get(meigen.registerId)
        display.push(`${speaker}: "${meigen.content}" added by ${register} (id is "${meigen.id}")`)
      }
      return await interaction.followUp(`${display.join('\n')}`)
    } else {
      return await interaction.reply('Direct Messageでは使用できません')
    }
  }
}