import {Events, GuildMember} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

module.exports = new TsukuyomiEvent({
    name: Events.GuildMemberRemove,
    run: async (client: TsukuyomiClient, member: GuildMember) => {

    }
})