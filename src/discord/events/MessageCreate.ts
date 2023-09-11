import {Events, Message} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";


module.exports = new TsukuyomiEvent({
    name: Events.MessageCreate,
    run: async (client: TsukuyomiClient, message: Message) => {
        if (message.author === client.user) {
            return
        }
        message.channel
    }
})