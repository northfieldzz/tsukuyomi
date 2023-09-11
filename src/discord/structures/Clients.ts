import {Client, ClientOptions, Collection, REST} from "discord.js";
import path from "path"
import requireAll from "require-all";
import {TsukuyomiEvent} from "./Event";
import {TsukuyomiCommand} from "./Command";

export default class TsukuyomiClient extends Client {
    // declare rest?: REST
    commands: Collection<string, TsukuyomiCommand> = new Collection()

    constructor(options: ClientOptions) {
        super(options);

    }

    async start() {
        await this.resolveModules()
        const token = process.env.DISCORD_TOKEN
        await this.login(token)
        this.rest = new REST().setToken(token!);
    }

    async resolveModules() {
        const sharedSettings = {
            recursive: true,
            filter: /\w*.[tj]s/g
        };
        // Register events
        requireAll({
            ...sharedSettings,
            dirname: path.join(__dirname, '../events'),
            resolve: (event: TsukuyomiEvent) => {
                this.on(event.name, async (...args) => {
                    try {
                        await event.run(this, ...args);
                    } catch (error) {
                        console.error(`An error occurred in '${event.name}' event.\n${error}\n`);
                    }
                })
                console.info(`Set event handler ${event.name}`)
            }
        })

        requireAll({
            ...sharedSettings,
            dirname: path.join(__dirname, '../commands'),
            resolve: (command: TsukuyomiCommand) => {
                if (command.disabled) {
                    return
                }
                if (command.permissions) {
                    command.builder.setDefaultPermission(false)
                }
                this.commands.set(command.builder.name, command);
                console.info(`Set command handler ${command.builder.name}`)
            }
        })
    }
}