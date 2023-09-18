import {TsukuyomiCommand} from "../structures/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class PigPresent implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('present-pig')
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
        if (!interaction.guild) {
            return await interaction.reply('特定のサーバから送信してください．')
        }
        await interaction.deferReply()
        const sender = interaction.user
        const receiver = interaction.options.getUser('to')!
        const point = interaction.options.data[1].value as number
        try {
            await prisma.$transaction(async (prisma) => {
                const senderPoint = await prisma.point.upsert({
                    where: {
                        userId_guildId: {
                            userId: sender.id,
                            guildId: interaction.guild!.id
                        }
                    },
                    create: {
                        userId: sender.id,
                        guildId: interaction.guild!.id,
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
                            guildId: interaction.guild!.id
                        }
                    },
                    create: {
                        userId: receiver.id,
                        guildId: interaction.guild!.id,
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

    }
}