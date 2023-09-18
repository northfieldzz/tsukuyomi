import {CommandInteraction, SlashCommandBuilder, SlashCommandSubcommandsOnlyBuilder} from "discord.js";
import TsukuyomiClient from "./Clients";

export interface TsukuyomiCommand {
    builder: Omit<SlashCommandBuilder, "addSubcommand" | "addSubcommandGroup"> | SlashCommandSubcommandsOnlyBuilder
    disabled?: boolean
    permissions?: string[]
    run: (client: TsukuyomiClient, interaction: CommandInteraction) => any
}