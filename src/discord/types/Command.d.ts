import {CommandInteraction, SlashCommandBuilder, SlashCommandSubcommandsOnlyBuilder} from "discord.js";
import TsukuyomiClient from "../structures/Clients";

export interface TsukuyomiCommand {
    builder: Omit<SlashCommandBuilder, "addSubcommand" | "addSubcommandGroup"> | SlashCommandSubcommandsOnlyBuilder
    disabled?: boolean
    permissions?: string[]
    run: (client: TsukuyomiClient, interaction: CommandInteraction) => any
}