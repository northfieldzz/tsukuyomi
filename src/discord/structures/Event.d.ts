import {ClientEvents} from 'discord.js';
import TsukuyomiClient from "./Clients";

export interface TsukuyomiEvent {
    name: keyof ClientEvents
    run: (client: TsukuyomiClient, ...eventArgs: any) => any
}