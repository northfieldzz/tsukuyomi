import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class PigPresent implements TsukuyomiCommand {
  builder = new SlashCommandBuilder()
    .setName('pig-present')
    .setDescription('pigをプレゼントします')
    .addUserOption(option => option
      .setName('to')
      .setDescription('ポイントを付与する相手を指定します')
      .setRequired(true)
    )
    .addIntegerOption(option => option
      .setName('amount')
      .setDescription('ポイントをいくつ付与するか指定します')
      .setRequired(true)
    )

  async run(client: TsukuyomiClient, interaction: CommandInteraction) {
    if (interaction.guild) {
      await interaction.deferReply()
      const sender = interaction.user
      const receiver = interaction.options.getUser('to')!
      const guildId = interaction.guild.id
      const point = interaction.options.data[1].value as number

      const sample = await interaction.guild.members.fetch(receiver.id)
      await sample.setNickname('sample')
      try {
        await prisma.$transaction(async (prisma) => {
          const senderPoint = await prisma.point.upsert({
            where: {
              userId_guildId: {
                userId: sender.id,
                guildId: guildId
              }
            },
            create: {
              userId: sender.id,
              guildId: guildId,
              value: -point
            },
            update: {
              value: {
                increment: -point
              }
            }
          })
          if (senderPoint.value < 0) {
            throw new Error(`${sender.globalName}は${point}pigも持っていない貧乏人です`)
          }
          await prisma.point.upsert({
            where: {
              userId_guildId: {
                userId: receiver.id,
                guildId: guildId
              }
            },
            create: {
              userId: receiver.id,
              guildId: guildId,
              value: point
            },
            update: {
              value: {
                increment: point
              }
            }
          })
        })
        return await interaction.followUp(`${sender.globalName}は${point}pigを${receiver.globalName}に渡しました．`)
      } catch (error: unknown) {
        let message = 'エラーで渡せなかったヨン．ごめんね！'
        if (error instanceof Error) {
          message = error.message
        }
        return await interaction.followUp(message)
      }
    } else {
      return await interaction.reply('Direct Messageでは使用できません')
    }
  }
}