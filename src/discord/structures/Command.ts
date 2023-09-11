import {CommandInteraction, SlashCommandBuilder} from "discord.js";
import TsukuyomiClient from "./Clients";

export class TsukuyomiCommand {
    builder: SlashCommandBuilder
    disabled?: boolean
    permissions?: string[]
    run: (client: TsukuyomiClient, interaction: CommandInteraction) => any

    constructor(options: NonNullable<TsukuyomiCommand>) {
        this.builder = options.builder
        this.disabled = options.disabled
        this.permissions = options.permissions
        this.run = options.run
    }
}