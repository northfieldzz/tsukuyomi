import {Client, ClientOptions, Collection, GuildMember, REST} from "discord.js"
import {TsukuyomiEvent} from "../types/Event"
import {TsukuyomiCommand} from "../types/Command"
import {events, commands} from "../config";


export default class TsukuyomiClient extends Client {
    commands: Collection<string, TsukuyomiCommand> = new Collection()

    constructor(options: ClientOptions) {
        super(options)
    }

    async start() {
        await this.register()
        const token = process.env.DISCORD_TOKEN
        await this.login(token)
        this.rest = new REST().setToken(token!)
    }

    async register() {
        for (const event of events) {
            const ev = new event() as TsukuyomiEvent
            this.on(ev.name, async (...args) => {
                try {
                    await ev.run(this, ...args);
                } catch (error) {
                    console.error(`An error occurred in '${ev.name}' event.\n${error}\n`);
                }
            })
            console.info(`Set event handler ${ev.name}`)
        }
        for (const command of commands) {
            const com = new command as TsukuyomiCommand
            this.commands.set(com.builder.name, com)
        }
    }
}

type Pig = number

async function setPigPropertyToNickname(member: GuildMember, pig: Pig) {
    await member.setNickname(`${member.user.username}@${pig}pig`)
}