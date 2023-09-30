import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {CharacterGenerator, summary} from "../../lib/prisma/CharacterStatus";
import {prisma} from "../../lib/prisma";


export class StatusGenerate implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('status-generate')
        .setDescription('ステータスを作成します')
        .addUserOption(option => option
            .setName('target')
            .setDescription('対象者')
        )

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (interaction.guildId) {
            const guildId = interaction.guildId
            await interaction.deferReply()
            let target = interaction.options.getUser('target')
            if (!target) {
                target = interaction.user
            }
            await interaction.followUp(`${target.globalName}のステータスを作成します`)
            await prisma.$transaction(async (prisma) => {
                const oldStatus = await prisma.characterStatus.findUnique({
                    where: {
                        guildId_userId: {
                            guildId: guildId,
                            userId: target!.id
                        }
                    }
                })
                if (oldStatus) {
                    await interaction.followUp(`もうステータスはあるみたい．つくりなおすお`)
                }
                const status = CharacterGenerator.generate()
                const newStatus = await prisma.characterStatus.upsert({
                    where: {
                        guildId_userId: {
                            guildId: guildId,
                            userId: target!.id
                        }
                    },
                    create: {
                        guildId: guildId,
                        userId: target!.id,
                        ...status
                    },
                    update: {
                        ...status
                    }
                })
                await interaction.followUp(summary(target?.username, newStatus))
            })
        } else {
            return await interaction.reply('特定のサーバから送信してください．')
        }
    }
}

