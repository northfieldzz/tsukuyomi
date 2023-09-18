import {ClientEvents, Events, GuildMember} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

export class GuildMemberRemove implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.GuildMemberRemove

    async run(client: TsukuyomiClient, member: GuildMember) {

    }
}