import {REST, Routes} from 'discord.js'
import {commands as commandClasses} from "../src/discord/config"

const token = process.env.DISCORD_TOKEN!
const guildId = process.env.DISCORD_GUILD_ID!
const clientId = process.env.DISCORD_APPLICATION_ID!

const commands = [];

for (const commandClass of commandClasses) {
    const command = new commandClass()
    commands.push(command.builder.toJSON());
}

// Construct and prepare an instance of the REST module
const rest = new REST().setToken(token);

// and deploy your commands!
(async () => {
    try {
        console.log(`Started refreshing ${commands.length} application (/) commands.`)
        await rest.put(Routes.applicationGuildCommands(clientId, guildId), {body: commands})
        console.log(`Successfully reloaded application (/) commands.`)
    } catch (error) {
        // And of course, make sure you catch and log any errors!
        console.error(error);
    }
})();