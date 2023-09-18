import {TsukuyomiCommand} from "../structures/Command";
import {CommandInteraction, SlashCommandBuilder, User} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {handlePoint} from "../../lib/prisma/Point";

export class PigGrant implements TsukuyomiCommand {
    builder = new SlashCommandBuilder()
        .setName('pig-grant')
        .setDescription('pigを付与します')
        .addUserOption(option => option
            .setName('to')
            .setDescription('ポイントを付与する相手を指定します．')
            .setRequired(true)
        )
        .addIntegerOption(option => option
            .setName('amount')
            .setDescription('ポイントをいくつ付与するか指定します')
            .setChoices(
                {name: '1pig', value: 1},
                {name: '2pig', value: 2},
                {name: '3pig', value: 3},
                {name: '4pig', value: 4},
                {name: '5pig', value: 5}
            )
            .setRequired(true)
        )

    async run(client: TsukuyomiClient, interaction: CommandInteraction) {
        if (!interaction.guild) {
            return await interaction.reply('特定のサーバから送信してください．')
        } else {
            await interaction.deferReply()
            const user = interaction.options.data[0].user as User
            const grantPoint = interaction.options.data[1].value as number
            const point = await handlePoint(user, interaction.guild, grantPoint, false)
            return await interaction.followUp(`${user.globalName}に${grantPoint}を付与しました． ${user.globalName}の所持pigは${point}です．`)
        }
    }
}