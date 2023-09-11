import {ClientEvents} from 'discord.js';
import TsukuyomiClient from "./Clients";

export class TsukuyomiEvent {
    /**
     * Event name
     */
    name: keyof ClientEvents;

    /**
     * Runs the event
     */
    run: (client: TsukuyomiClient, ...eventArgs: any) => any;

    /**
     * Creates a new event
     * @param options Event options
     */
    constructor(options: NonNullable<TsukuyomiEvent>) {
        this.name = options.name
        this.run = options.run
    }
}