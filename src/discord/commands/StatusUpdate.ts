import {TsukuyomiCommand} from "../structures/Command";
import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {prisma} from "../../lib/prisma";
import {summary} from "../../lib/prisma/CharacterStatus";


export class StatusUpdate implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('status-update')
        .setDescription('ステータスを更新します')
        .addUserOption(option => option
            .setName('target')
            .setDescription('対象者')
        )
        .addIntegerOption(option => option.setName('str').setDescription('筋力'))
        .addIntegerOption(option => option.setName('con').setDescription('体力'))
        .addIntegerOption(option => option.setName('pow').setDescription('パワー，精神力'))
        .addIntegerOption(option => option.setName('dex').setDescription('俊敏性'))
        .addIntegerOption(option => option.setName('app').setDescription('外見'))
        .addIntegerOption(option => option.setName('siz').setDescription('体格'))
        .addIntegerOption(option => option.setName('int').setDescription('知性'))
        .addIntegerOption(option => option.setName('edu').setDescription('教育'))
        .addIntegerOption(option => option.setName('luk').setDescription('幸運'))
        .addIntegerOption(option => option.setName('san').setDescription('正気度'))
        .addIntegerOption(option => option.setName('ida').setDescription('アイデア'))
        .addIntegerOption(option => option.setName('know').setDescription('知識'))
        .addIntegerOption(option => option.setName('hp').setDescription('ヘルスポイント'))
        .addIntegerOption(option => option.setName('mov').setDescription('移動率'))
        .addIntegerOption(option => option.setName('mp').setDescription('マジックポイント'))
        .addIntegerOption(option => option.setName('skl').setDescription('スキルポイント'))
        .addIntegerOption(option => option.setName('db').setDescription('ダメージボーナス'))
        .addIntegerOption(option => option.setName('build').setDescription('建築スキル'))


    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (!interaction.guildId) {
            return await interaction.reply('特定のサーバから送信してください．')
        } else {
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
                            guildId: interaction.guildId!,
                            userId: target!.id
                        }
                    }
                })
                if (!oldStatus) {
                    return await interaction.followUp(`まだステータスはないみたい．"/status-generate"でステータスを作ってね`)
                } else {
                    await interaction.followUp(summary(`${target?.globalName} old`, oldStatus))
                }

                const status = {
                    // @ts-ignore
                    str: interaction.options.getInteger('str') ?? oldStatus.str,
                    // @ts-ignore
                    con: interaction.options.getInteger('con') ?? oldStatus.con,
                    // @ts-ignore
                    pow: interaction.options.getInteger('pow') ?? oldStatus.pow,
                    // @ts-ignore
                    dex: interaction.options.getInteger('dex') ?? oldStatus.dex,
                    // @ts-ignore
                    app: interaction.options.getInteger('app') ?? oldStatus.app,
                    // @ts-ignore
                    siz: interaction.options.getInteger('siz') ?? oldStatus.siz,
                    // @ts-ignore
                    int: interaction.options.getInteger('int') ?? oldStatus.int,
                    // @ts-ignore
                    edu: interaction.options.getInteger('edu') ?? oldStatus.edu,
                    // @ts-ignore
                    luk: interaction.options.getInteger('luk') ?? oldStatus.luk,
                    // @ts-ignore
                    san: interaction.options.getInteger('san') ?? oldStatus.san,
                    // @ts-ignore
                    ida: interaction.options.getInteger('ida') ?? oldStatus.ida,
                    // @ts-ignore
                    know: interaction.options.getInteger('know') ?? oldStatus.know,
                    // @ts-ignore
                    hp: interaction.options.getInteger('hp') ?? oldStatus.hp,
                    // @ts-ignore
                    mov: interaction.options.getInteger('mov') ?? oldStatus.mov,
                    // @ts-ignore
                    mp: interaction.options.getInteger('mp') ?? oldStatus.mp,
                    // @ts-ignore
                    skl: interaction.options.getInteger('skl') ?? oldStatus.skl,
                    // @ts-ignore
                    db: interaction.options.getInteger('db') ?? oldStatus.db,
                    // @ts-ignore
                    build: interaction.options.getInteger('build') ?? oldStatus.build,
                }
                const newStatus = await prisma.characterStatus.upsert({
                    where: {
                        guildId_userId: {
                            guildId: interaction.guildId!,
                            userId: target!.id
                        }
                    },
                    create: {
                        guildId: interaction.guildId!,
                        userId: target!.id,
                        ...status
                    },
                    update: {
                        ...status
                    }
                })
                await interaction.followUp('↓')
                return await interaction.followUp(summary(`${target?.globalName} new`, newStatus))
            })
        }
    }
}