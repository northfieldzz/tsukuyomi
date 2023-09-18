import {ClientEvents, Events, Message} from "discord.js";
import TsukuyomiClient from "../structures/Clients";
import {TsukuyomiEvent} from "../structures/Event";

export class MessageCreate implements TsukuyomiEvent {
    name: keyof ClientEvents = Events.MessageCreate

    async run(client: TsukuyomiClient, message: Message) {
        if (message.author === client.user) {
            return
        }
        message.channel
    }
}