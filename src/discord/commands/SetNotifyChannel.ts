import {TsukuyomiCommand} from "../structures/Command";
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
        await interaction.deferReply()
        if (!interaction.guildId) {
            return await interaction.followUp('特定のサーバから送信してください．')
        }
        const channel = interaction.options.data[0].channel
        if (!channel) {
            return await interaction.followUp('ちゃうねんちゃうねん，channelがないねん')
        }
        await prisma.notificationChannel.upsert({
            where: {
                guildId: interaction.guildId!
            },
            create: {
                guildId: interaction.guildId!,
                channelId: channel!.id
            },
            update: {
                channelId: channel!.id
            }
        })
        return await interaction.followUp(`お知らせチャンネルを${channel.name}に設定しました`)
    }
}