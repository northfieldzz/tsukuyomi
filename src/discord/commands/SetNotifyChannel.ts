import {TsukuyomiCommand} from "../types/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import {ChannelType} from "discord-api-types/v10"
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";

export class SetNotifyChannel implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('set-notify-channel')
        .setDescription('お知らせチャンネルを指定します')
        .addChannelOption(option => option
            .setName('channel')
            .setDescription('チャンネル')
            .setRequired(true)
            .addChannelTypes(ChannelType.GuildText)
        )

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (interaction.guildId) {
            await interaction.deferReply()
            const channel = interaction.options.data[0].channel
            if (!channel) {
                return await interaction.followUp('ちゃうねんちゃうねん，channelがないねん')
            }
            await prisma.notificationChannel.upsert({
                where: {
                    guildId: interaction.guildId
                },
                create: {
                    guildId: interaction.guildId,
                    channelId: channel!.id
                },
                update: {
                    channelId: channel!.id
                }
            })
            return await interaction.followUp(`お知らせチャンネルを${channel.name}に設定しました`)
        } else {
            return await interaction.followUp('Direct Messageでは使用できません')
        }
    }
}